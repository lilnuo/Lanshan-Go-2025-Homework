package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		method := c.Request.Method
		status := c.Writer.Status()
		latency := time.Since(start)
		path := c.Request.URL.Path
		log.Printf("%-6s |%13v| %3d |%s", method, latency, status, path)
	}
}
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered, err:%v\n", err)
				c.JSON(500, gin.H{
					"code":    500,
					"message": "Internal Server Error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) //No content
			return
		}
		c.Next()
	}
}
func main() {
	r := gin.New()
	r.Use(Recovery())
	r.Use(Logger())
	r.Use(CORS())
	r.GET("/api/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "张三",
			"age":  18,
		})
	})
	r.GET("/api/error", func(c *gin.Context) {
		var m map[string]string
		m["key"] = "value"
	})
	fmt.Println("Server Run On http://127.0.0.1:8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

// 自定义 Recovery
/*func MyRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(">>> [捕获成功] Recovery 中间件抓住了 Panic！错误信息:", err)
				// 强制返回 500 JSON
				c.JSON(500, gin.H{
					"status": "error",
					"msg":    "我是 Recovery 返回的错误信息",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func main() {
	// 必须用 New，不能用 Default
	r := gin.New()

	// 只注册 Recovery，去掉 Logger 等其他干扰
	r.Use(MyRecovery())

	// 测试接口
	r.GET("/test", func(c *gin.Context) {
		// 故意制造一个除以 0 的 panic
		var a = 0
		var b = 1 / a
		// 这行永远执行不到
		c.String(200, fmt.Sprintf("结果: %d", b))
	})

	r.Run(":8080")
}*/
