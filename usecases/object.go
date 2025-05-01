package usecases

import "http_server_arch/domain"

type Object interface {
	Get(id string) (*domain.Task, error)
	Put(task *domain.Task) error
	Post(task *domain.Task) error
	Delete(id string) error
}
