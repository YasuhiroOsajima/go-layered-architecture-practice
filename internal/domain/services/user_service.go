package services

import "go-layered-architecture-practice/internal/domain/models/user"

type UserService struct {
	repository user.UserRepositoryInterface
}

func NewUserService(repository user.UserRepositoryInterface) UserService {
	return UserService{repository}
}

func (s UserService) Exists(targetUser *user.User) (bool, error) {
	res, err := s.repository.Find(targetUser.Id())
	if err != nil {
		return false, err
	}

	if res != nil {
		return true, nil
	} else {
		return false, nil
	}
}
