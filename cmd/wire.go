//+build wireinject

package main

import (
	"github.com/google/wire"

	user_model "go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
	sqlite_user_repo "go-layered-architecture-practice/internal/repository/sqlite/user"
)

func InitializeUserRepository() user_model.UserRepositoryInterface {
	wire.Build(
		wire.InterfaceValue(new(user_model.UserRepositoryInterface), sqlite_user_repo.NewUserRepository()),
	)

	return nil
}

func InitializeUserService() (userService services.UserService) {
	wire.Build(
		wire.InterfaceValue(new(user_model.UserRepositoryInterface), sqlite_user_repo.NewUserRepository()),
		services.NewUserService,
	)

	return userService
}

func InitializeUserFactory() (userFactory services.UserFactory) {
	wire.Build(
		wire.InterfaceValue(new(user_model.UserRepositoryInterface), sqlite_user_repo.NewUserRepository()),
		services.NewUserFactory,
	)

	return userFactory
}
