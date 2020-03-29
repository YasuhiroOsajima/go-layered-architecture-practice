//+build wireinject

package go_layered_architecture_practice

import (
	"github.com/google/wire"

	user_model "go-layered-architecture-practice/internal/domain/models/user"
	user_repo "go-layered-architecture-practice/internal/repository/sqlite/user"
)

func InitializeUserRepository() (userRepository user_model.UserRepositoryInterface) {
	wire.Build(
		user_repo.NewUserRepository,
	)

	return userRepository
}

//func InitializeUserService() (userService services.UserService, err error) {
//	wire.Build(
//		user.NewUserRepository,
//		services.NewUserService,
//	)
//
//	return userService, err
//}
