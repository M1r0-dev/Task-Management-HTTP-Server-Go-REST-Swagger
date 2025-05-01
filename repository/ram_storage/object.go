package ram_storage

import (
	"http_server_arch/domain"
	"http_server_arch/repository"
	"sync"
)

type Object struct {
	data map[string]*domain.Task
	mu   sync.RWMutex
}

func NewObject() *Object {
	return &Object{
		data: make(map[string]*domain.Task),
	}
}

func (rs *Object) Get(id string) (*domain.Task, error) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()
	task, exists := rs.data[id]
	if !exists {
		return nil, repository.ErrNotFound
	}
	return task, nil
}

func (rs *Object) Post(task *domain.Task) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	task, exists := rs.data[task.ID]
	if !exists {
		return repository.ErrNotFound
	}
	rs.data[task.ID] = task
	return nil

}

func (rs *Object) Put(task *domain.Task) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	if _, exists := rs.data[task.ID]; exists {
		return repository.ErrAlreadyExists
	}
	rs.data[task.ID] = task
	return nil
}

func (rs *Object) Delete(id string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	if _, exists := rs.data[id]; !exists {
		return repository.ErrNotFound
	}
	delete(rs.data, id)
	return nil
}
