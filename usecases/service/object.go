package service

import (
	"http_server_arch/domain"
	"http_server_arch/repository"
)

type Object struct {
	repo repository.Object
}

func NewObject(repo repository.Object) *Object {
	return &Object{
		repo: repo,
	}
}

func (rs *Object) Get(id string) (*domain.Task, error) {
	return rs.repo.Get(id)
}

func (rs *Object) Put(task *domain.Task) error {
	return rs.repo.Put(task)
}

func (rs *Object) Post(task *domain.Task) error {
	return rs.repo.Post(task)
}

func (rs *Object) Delete(id string) error {
	return rs.repo.Delete(id)
}
