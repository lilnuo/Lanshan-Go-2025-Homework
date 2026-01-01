package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret_key")
var userDB = map[string]User{
	"admin": {Username: "admin", Password: "123456", Role: "admin"},
	"user":  {Username: "user", Password: "123456", Role: "user"},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user User) (string, error) {
	claims := Claims{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "需要认证"})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "格式错误",
			})
			c.Abort()
			return
		}
		tokenString := parts[1]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效或已过期",
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*Claims); ok {
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)
		}
		c.Next()
	}
}
func RequireRole(requireRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "not found role",
			})
			c.Abort()
			return
		}
		if userRole != requireRole {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "全线不足（需要+requireRole+",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误",
		})
		return
	}
	storedUser, exists := userDB[u.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "账号密码错误",
		})
		return
	}
	token, _ := GenerateToken(storedUser)
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "登录成功",
		"role":    storedUser.Role,
	})
}
func GetProfile(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"user": username,
		"msg":  "欢迎查看个人资料",
	})
}
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete success",
	})

}
func main() {
	r := gin.Default()
	r.POST("/login", Login)
	authorized := r.Group("/api")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/profile", GetProfile)
	}
	admin := r.Group("/api/admin")
	admin.Use(AuthMiddleware())
	admin.Use(RequireRole("admin"))
	{
		admin.GET("/users", DeleteUser)
	}
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
