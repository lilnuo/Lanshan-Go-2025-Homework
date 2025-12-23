package main

import (
	"lesson6/handler"
	"lesson6/service"
	"lesson6/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	s := storage.MemoryStorageNew()
	todoService := service.NewTodoService(s)
	todoHandler := handler.NewTodoHandler(todoService)
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.POST("/tasks", todoHandler.CreateTask)
		api.GET("/tasks/:id", todoHandler.GetTask)
		api.GET("/tasks", todoHandler.GetAllTask)
		api.PUT("/tasks/:id", todoHandler.UpdateTask)
		api.DELETE("/tasks/:id", todoHandler.DeleteTask)
	}
	log.Println("gin successfully serve running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("gin run error:", err)
	}
}
