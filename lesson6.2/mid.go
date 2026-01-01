package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func LoggerMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		fmt.Printf("start timer%v", start)
		c.Next(ctx)
		end := time.Since(start)
		fmt.Printf("end timer%v", end)
	}
}
func RecoveryMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("find err =%v", err)
				c.JSON(consts.StatusInternalServerError, map[string]string{
					"msg": "服务器错误",
				})
				c.Abort()
			}
		}()
		c.Next(ctx)
	}
}
func CORSMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type")
		if string(c.Method()) == "OPTIONS" {
			c.Status(consts.StatusNoContent)
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
func PingHandler(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, map[string]string{
		"msg": "okk",
	})
}
func CrashHandler(ctx context.Context, c *app.RequestContext) {
	panic("boom boom")
}
func main() {
	h := server.New(
		server.WithHostPorts(":8080"))
	h.Use(RecoveryMiddleware())
	h.Use(CORSMiddleware())
	h.Use(LoggerMiddleware())
	h.GET("/ping", PingHandler)
	h.GET("/crash", CrashHandler)
	fmt.Println("success server ")
	h.Spin()
}
