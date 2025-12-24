package service

import (
	"errors"
	"time"

	"lesson6.1/model"

	"lesson6.1/storage"
)

type TodoService struct {
	storage storage.Storage
}

func NewTodoService(s storage.Storage) *TodoService {
	return &TodoService{storage: s}
}
func (s *TodoService) Create(title string) (*model.Task, error) {
	if title == "" {
		return nil, errors.New("title is required")
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
func (s *TodoService) Get(id string) (*model.Task, error) {
	return s.storage.GetById(id)
}
func (s *TodoService) GetAll() ([]*model.Task, error) {
	return s.storage.GetAll()
}

func (s *TodoService) Update(id string, title string, completed bool) (*model.Task, error) {
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
func (s *TodoService) Delete(id string) error {
	return s.storage.Delete(id)
}
