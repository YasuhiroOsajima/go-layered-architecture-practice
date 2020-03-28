package services

import "go-layered-architecture-practice/internal/domain/models/user"

type UserService struct {
	repository user.UserRepositoryInterface
}

func NewUserService(repository user.UserRepositoryInterface) UserService {
	return UserService{repository}
}

func (s UserService) Exists(targetUser *user.User) (bool, error) {
	sameNames, err := s.repository.FindAll(targetUser.Name())
	if err != nil {
		return false, err
	}

	if len(sameNames) != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
