package user

import (
	"fmt"
	"go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
)

type UserApplicationService struct {
	userRepo    user.UserRepositoryInterface
	userService services.UserService
}

func NewUserApplicationService(repo user.UserRepositoryInterface, service services.UserService) UserApplicationService {
	return UserApplicationService{repo, service}
}

func (u UserApplicationService) Register(name string) error {
	userName, err := user.NewUserName(name)
	if err != nil {
		return err
	}

	newUser, err := user.NewUserInit(userName)
	if err != nil {
		return err
	}
	fmt.Println(newUser)

	return nil
}
