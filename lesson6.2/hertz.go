package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtsecret = []byte("secret")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var userDB = map[string]User{
	"administrator": {Username: "administrator", Password: "1232222", Role: "admin"},
	"user1":         {Username: "user1", Password: "123456", Role: "user"},
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtsecret)
}
func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(consts.StatusUnauthorized, map[string]string{"error": "未登录"})
			c.Abort()
			return
		}
		tokenString := authHeader
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtsecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalid",
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*Claims); ok {
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)
		}
		c.Next(ctx)
	}
}
func RequireRole(requiredRole string) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(consts.StatusUnauthorized, map[string]string{"error": "no role found"})
			c.Abort()
			return
		}
		if role != requiredRole {
			c.JSON(consts.StatusForbidden, map[string]string{"error": "全线不足"})
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
func Login(ctx context.Context, c *app.RequestContext) {
	var u User
	if err := c.BindAndValidate(&u); err != nil {
		c.JSON(consts.StatusBadRequest, map[string]string{"error": "参数错误"})
		return
	}
	storedUser, exists := userDB[u.Username]
	if !exists || storedUser.Password != u.Password {
		c.JSON(consts.StatusUnauthorized, map[string]string{"error": "账号或密码错误"})
		return
	}
	token, _ := GenerateToken(storedUser)
	c.JSON(consts.StatusOK, map[string]string{
		"msg":   "login success",
		"token": token,
		"role":  storedUser.Role,
	})
}
func GetProfile(ctx context.Context, c *app.RequestContext) {
	username := c.Param("username")
	c.JSON(consts.StatusOK, map[string]string{
		"msg":      "get profile success",
		"username": username,
	})
}
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]string{
		"msg": "delete user success",
	})
}
func main() {
	h := server.Default(server.WithHostPorts(":8080"))
	h.POST("/login", Login)
	authGroup := h.Group("/api")
	authGroup.Use(AuthMiddleware())
	{
		authGroup.GET("/profile", GetProfile)
		adminGroup := h.Group("/admin")
		adminGroup.Use(RequireRole("admin"))
		{
			adminGroup.DELETE("/users", DeleteUser)
		}
	}
	h.Spin()
}
