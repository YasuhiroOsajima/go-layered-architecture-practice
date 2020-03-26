package user

import (
	"errors"
	user_model "go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
)

type UserApplicationService struct {
	userRepository user_model.UserRepositoryInterface
	userService    services.UserService
}

func NewUserApplicationService(repo user_model.UserRepositoryInterface, service services.UserService) UserApplicationService {
	return UserApplicationService{repo, service}
}

func (u UserApplicationService) Register(name string) error {
	userName, err := user_model.NewUserName(name)
	if err != nil {
		return err
	}

	newUser, err := user_model.NewUserInit(userName, u.userRepository)
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

func (u UserApplicationService) Get(Id string) (*user_model.User, error) {
	userId, err := user_model.NewUserId(Id)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
