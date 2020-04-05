package user

import (
	"errors"

	user_model "go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
)

type UserApplicationService struct {
	userRepository user_model.UserRepositoryInterface
	userService    services.UserService
	userFactory    services.UserFactory
}

func NewUserApplicationService(repo user_model.UserRepositoryInterface, service services.UserService, factory services.UserFactory) UserApplicationService {
	return UserApplicationService{repo, service, factory}
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

	newUser, err := u.userFactory.Create(userName, userMailAddress)
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
	return err
}

func (u UserApplicationService) Get(command UserGetCommand) (UserGetResult, error) {
	var userGetResult UserGetResult
	var userData UserData

	id, idErr := command.GetId()
	name, nameErr := command.GetName()

	if idErr != nil {
		userId, err := user_model.NewUserId(id)
		if err != nil {
			return userGetResult, err
		}

		user, err := u.userRepository.Find(userId)
		if err != nil {
			return userGetResult, err
		}

		if user != nil {
			return userGetResult, errors.New("target user is not found")
		}

		userData = NewUserData(user)

	} else if nameErr != nil {
		userName, err := user_model.NewUserName(name)
		if err != nil {
			return userGetResult, err
		}

		users, err := u.userRepository.FindAll(userName)
		if err != nil {
			return userGetResult, err
		}

		if len(users) == 0 {
			return userGetResult, errors.New("target user is not found")
		}

		if len(users) != 1 {
			return userGetResult, errors.New("target user name is duplicated")
		}

		user := users[0]
		userData = NewUserData(user)

	} else {
		return userGetResult, errors.New("both arguments were not specified")
	}

	userGetResult = NewUserGetResult(userData)
	return userGetResult, nil
}

func (u UserApplicationService) Delete(id string) error {
	userId, err := user_model.NewUserId(id)
	if err != nil {
		return err
	}

	user, err := u.userRepository.Find(userId)
	if err != nil {
		return err
	}

	err = u.userRepository.Delete(user)
	return err
}
