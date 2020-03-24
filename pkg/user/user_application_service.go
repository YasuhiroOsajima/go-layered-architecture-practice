package user

import (
	"errors"
	"go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
)

type UserApplicationService struct {
	userRepository user.UserRepositoryInterface
	userService    services.UserService
}

func NewUserApplicationService(repo user.UserRepositoryInterface, service services.UserService) UserApplicationService {
	return UserApplicationService{repo, service}
}

func (u UserApplicationService) Register(name string) error {
	userName, err := user.NewUserName(name)
	if err != nil {
		return err
	}

	newUser, err := user.NewUserInit(userName, u.userRepository)
	if err != nil {
		return err
	}

	exists, err := u.userService.Exists(newUser)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("same name user is already exists")
	}

	err = u.userRepository.Save(newUser)
	return err
}
