package main

import (
	"fmt"

	"go-layered-architecture-practice/internal/domain/services"
	user_repo "go-layered-architecture-practice/internal/repository/sqlite/user"
	user_app "go-layered-architecture-practice/pkg/user"
)

func main() {
	inMemory, err := user_repo.NewUserRepository()
	if err != nil {
		fmt.Println(err)
		return
	}
	userService := services.NewUserService(inMemory)

	app := user_app.NewUserApplicationService(inMemory, userService)
	err = app.Register("aaa", "test@sample.hoge")
	if err != nil {
		fmt.Println(err)
		return
	}
}
