package handler

import (
	"lesson6/model"
	"lesson6/service"
	"net/http"
)
import "github.com/gin-gonic/gin"

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{service: s}
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func (h *TodoHandler) CreateTask(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: err.Error()})
		return
	}
	task, err := h.service.CreateTask(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: err.Error()})
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "success", Data: task})
}
func (h *TodoHandler) GetAllTask(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "success", Data: tasks})
}
func (h *TodoHandler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.service.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{Code: 404, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "success", Data: task})
}
func (h *TodoHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var req model.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: err.Error()})
		return
	}
	task, err := h.service.UpdateTask(id, req.Title, req.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "success", Data: task})
}
func (h *TodoHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "success"})
}
