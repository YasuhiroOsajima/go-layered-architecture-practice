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

func (u UserApplicationService) Register(name, mailAddress string) error {
	userName, err := user_model.NewUserName(name)
	if err != nil {
		return err
	}

	userMailAddress, err := user_model.NewUserMailAddress(mailAddress)
	if err != nil {
		return err
	}

	newUser, err := user_model.NewUserInit(userName, userMailAddress, u.userRepository)
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

func (u UserApplicationService) Update(command UserUpdateCommand) error {
	userId, err := user_model.NewUserId(command.GetId())
	if err != nil {
		return err
	}

	user, err := u.userRepository.Find(userId)
	if err != nil {
		return err
	}

	nameArg, err := command.GetName()
	if err != nil {
		userName, err := user_model.NewUserName(nameArg)
		if err != nil {
			return err
		}

		user.ChangeName(userName)

		found, err := u.userService.Exists(user)
		if err != nil {
			return err
		}
		if found {
			return errors.New("same name user is already exists")
		}
	}

	mailAddress, err := command.GetMailAddress()
	if err != nil {
		userMailAddress, err := user_model.NewUserMailAddress(mailAddress)
		if err != nil {
			return err
		}

		user.ChangeMailAddress(userMailAddress)
	}

	err = u.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (u UserApplicationService) Get(id string) (UserData, error) {
	var userData UserData
	userId, err := user_model.NewUserId(id)
	if err != nil {
		return userData, err
	}

	user, err := u.userRepository.Find(userId)
	if err != nil {
		return userData, err
	}

	userData = NewUserData(user)
	return userData, nil
}
