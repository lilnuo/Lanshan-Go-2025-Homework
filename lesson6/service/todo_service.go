package service

import (
	"errors"
	"lesson6/model"
	"lesson6/storage"
	"time"
)

type TodoService struct {
	storage storage.Storage
}

func NewTodoService(s storage.Storage) *TodoService {
	return &TodoService{storage: s}
}
func (s *TodoService) CreateTask(title string) (*model.Task, error) {
	if title == "" {
		return nil, errors.New("title is empty")
	}
	task := &model.Task{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	if err := s.storage.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}
func (s *TodoService) GetTask(id string) (*model.Task, error) {
	return s.storage.GetById(id)
}
func (s *TodoService) GetAllTasks() ([]*model.Task, error) {
	return s.storage.GetAll()
}
func (s *TodoService) UpdateTask(id string, title string, completed bool) (*model.Task, error) {
	updateTask := &model.Task{
		Title:     title,
		Completed: completed,
	}
	err := s.storage.Update(id, updateTask)
	if err != nil {
		return nil, err
	}
	return s.storage.GetById(id)
}
func (s *TodoService) DeleteTask(id string) error {
	return s.storage.Delete(id)
}
