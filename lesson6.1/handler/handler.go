package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	model2 "lesson6.1/model"
	"lesson6.1/service"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{service: s}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *TodoHandler) Create(cxt context.Context, c *app.RequestContext) {
	var req struct {
		Title string `json:"title" tag:"required"`
	}
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, &Response{
			Code: 400, Message: err.Error(),
		})
		return
	}
	task, err := h.service.Create(req.Title)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &Response{
			Code: 500, Message: err.Error(),
		})
	}
	c.JSON(consts.StatusOK, &Response{
		Code: 200, Message: "success", Data: task,
	})
}
func (h *TodoHandler) GetAll(ctx context.Context, c *app.RequestContext) {
	tasks, err := h.service.GetAll()
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &Response{
			Code: 500, Message: err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, &Response{
		Code: 200, Message: "success", Data: tasks,
	})
}
func (h *TodoHandler) Get(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	task, err := h.service.Get(id)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &Response{
			Code: 500, Message: err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, &Response{
		Code: 200, Message: "success", Data: task,
	})
}
func (h *TodoHandler) Update(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	var req model2.Task
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, &Response{
			Code: 400, Message: err.Error(),
		})
		return
	}
	task, err := h.service.Update(id, req.Title, req.Completed)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &Response{
			Code: 500, Message: err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, &Response{
		Code: 200, Message: "success", Data: task,
	})
}
func (h *TodoHandler) Delete(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &Response{
			Code: 500, Message: err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, &Response{
		Code: 200, Message: "Task delete"})
}
