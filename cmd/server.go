package main

import (
	"fmt"

	"go-layered-architecture-practice/internal/domain/services"
	user_repo "go-layered-architecture-practice/internal/repository/sqlite/user"
	user_app "go-layered-architecture-practice/pkg/user"
)

func main() {
	sqlite := user_repo.NewUserRepository()
	userService := services.NewUserService(sqlite)

	app := user_app.NewUserApplicationService(sqlite, userService)
	err := app.Register("aaa", "test@sample.hoge")
	if err != nil {
		fmt.Println(err)
		return
	}
}
