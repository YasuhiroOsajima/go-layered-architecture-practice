package services

import (
	"go-layered-architecture-practice/internal/domain/models/circle"
	"go-layered-architecture-practice/internal/domain/models/user"
)

type CircleFactory struct {
	repository circle.CircleRepositoryInterface
}

func NewCircleFactory(repository circle.CircleRepositoryInterface) CircleFactory {
	return CircleFactory{repository}
}

func (f CircleFactory) Create(name circle.CircleName, owner *user.User) (*circle.Circle, error) {
	return nil, nil
}
