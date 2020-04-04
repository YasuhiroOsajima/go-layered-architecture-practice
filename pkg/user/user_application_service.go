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

func (u UserApplicationService) Register(name, mailAddress string) (result UserGetResultInterface) {
	userName, err := user_model.NewUserName(name)
	if err != nil {
		result.Status(500)
		return result
	}

	userMailAddress, err := user_model.NewUserMailAddress(mailAddress)
	if err != nil {
		result.Status(500)
		return result
	}

	newUser, err := u.userFactory.Create(userName, userMailAddress)
	if err != nil {
		result.Status(500)
		return result
	}

	exists, err := u.userService.Exists(newUser)
	if err != nil {
		result.Status(500)
		return result
	}

	if exists {
		result.JSON(400, errors.New("same name user is already exists"))
		return result
	}

	err = u.userRepository.Save(newUser)
	if err != nil {
		result.Status(500)
		return result
	}

	result.Status(200)
	return result
}

func (u UserApplicationService) Update(command UserUpdateCommand) (result UserGetResultInterface) {
	userId, err := user_model.NewUserId(command.GetId())
	if err != nil {
		result.Status(500)
		return result
	}

	user, err := u.userRepository.Find(userId)
	if err != nil {
		result.Status(500)
		return result
	}

	nameArg, err := command.GetName()
	if err != nil {
		userName, err := user_model.NewUserName(nameArg)
		if err != nil {
			result.Status(500)
			return result
		}

		user.ChangeName(userName)

		found, err := u.userService.Exists(user)
		if err != nil {
			result.Status(500)
			return result
		}
		if found {
			result.JSON(400, errors.New("same name user is already exists"))
			return result
		}
	}

	mailAddress, err := command.GetMailAddress()
	if err != nil {
		userMailAddress, err := user_model.NewUserMailAddress(mailAddress)
		if err != nil {
			result.Status(500)
			return result
		}

		user.ChangeMailAddress(userMailAddress)
	}

	err = u.userRepository.Save(user)
	if err != nil {
		result.Status(500)
		return result
	}

	result.Status(200)
	return result
}

func (u UserApplicationService) Get(command UserGetCommandInterface) {
	var userData UserData

	id, idErr := command.GetId()
	name, nameErr := command.GetName()

	if idErr != nil {
		userId, err := user_model.NewUserId(id)
		if err != nil {
			command.JSON(400, err)
			return
		}

		user, err := u.userRepository.Find(userId)
		if err != nil {
			command.Status(500)
			return
		}

		if user != nil {
			command.JSON(400, errors.New("target user is not found"))
			return
		}

		userData = NewUserData(user)

	} else if nameErr != nil {
		userName, err := user_model.NewUserName(name)
		if err != nil {
			command.JSON(400, err)
			return
		}

		users, err := u.userRepository.FindAll(userName)
		if err != nil {
			command.Status(500)
			return
		}

		if len(users) == 0 {
			command.JSON(400, errors.New("target user is not found"))
			return
		}

		if len(users) != 1 {
			command.JSON(400, errors.New("target user name is duplicated"))
			return
		}

		user := users[0]
		userData = NewUserData(user)

	} else {
		command.JSON(400, errors.New("both arguments were not specified"))
		return
	}

	command.JSON(200, userData)
	return
}

func (u UserApplicationService) Delete(id string) (result UserGetResultInterface) {
	userId, err := user_model.NewUserId(id)
	if err != nil {
		result.JSON(400, err)
		return result
	}

	user, err := u.userRepository.Find(userId)
	if err != nil {
		result.Status(500)
		return result
	}

	err = u.userRepository.Delete(user)
	if err != nil {
		result.Status(500)
		return result
	}

	result.Status(200)
	return result
}
