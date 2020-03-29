//+build wireinject

package main

import (
	"github.com/google/wire"

	user_model "go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
	sqlite_user_repo "go-layered-architecture-practice/internal/repository/sqlite/user"
)

func InitializeSQLiteUserRepository() user_model.UserRepositoryInterface {
	wire.Build(
		wire.InterfaceValue(new(user_model.UserRepositoryInterface), sqlite_user_repo.NewUserRepository()),
	)

	return nil
}

func InitializeSQLiteUserService() (userService services.UserService) {
	wire.Build(
		wire.InterfaceValue(new(user_model.UserRepositoryInterface), sqlite_user_repo.NewUserRepository()),
		services.NewUserService,
	)

	return userService
}
