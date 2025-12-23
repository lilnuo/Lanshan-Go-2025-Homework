package storage

import (
	"fmt"
	"lesson6/model"
	"sync"
)

type Storage interface {
	Create(task *model.Task) error
	GetById(id string) (*model.Task, error)
	GetAll() ([]*model.Task, error)
	Update(id string, task *model.Task) error
	Delete(id string) error
}
type MemoryStorage struct {
	tasks   map[string]*model.Task
	mu      sync.RWMutex
	idCount int
}

func MemoryStorageNew() Storage {
	return &MemoryStorage{
		tasks: make(map[string]*model.Task),
	}
}
func (m *MemoryStorage) Create(task *model.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.idCount++
	taskID := fmt.Sprintf("%d", m.idCount)
	m.tasks[taskID] = task
	return nil
}
func (m *MemoryStorage) GetById(id string) (*model.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	task, ok := m.tasks[id]
	if !ok {
		return nil, fmt.Errorf("task not found")
	}
	return task, nil
}
func (m *MemoryStorage) GetAll() ([]*model.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	allTasks := make([]*model.Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks, nil
}
func (m *MemoryStorage) Update(id string, updateTask *model.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.tasks[id]; ok {
		fmt.Errorf("task not found")
	}
	updateTask.ID = id
	m.tasks[id] = updateTask
	return nil
}
func (m *MemoryStorage) Delete(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.tasks[id]; ok {
		fmt.Errorf("task not found")
	}
	delete(m.tasks, id)
	return nil
}
