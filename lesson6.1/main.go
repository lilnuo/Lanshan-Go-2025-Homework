package main

import (
	"log"

	"github.com/cloudwego/hertz/pkg/app/server"
	"lesson6.1/handler"
	"lesson6.1/service"
	"lesson6.1/storage"
)

func main() {
	s := storage.MemoryStorageNew()
	todoService := service.NewTodoService(s)
	todoHandler := handler.NewTodoHandler(todoService)
	h := server.Default()
	api := h.Group("/api/v1")
	{
		api.POST("/task", todoHandler.Create)
		api.GET("/task", todoHandler.GetAll)
		api.GET("/task/:id", todoHandler.Get)
		api.PUT("/task/:id", todoHandler.Update)
		api.DELETE("/task/:id", todoHandler.Delete)
	}
	log.Println("Starting server at :8888")
	h.Spin()

}
