package main

import (
	"fmt"

	user_app "go-layered-architecture-practice/pkg/user"
)

func main() {
	sqlite := InitializeUserRepository()
	userService := InitializeUserService()
	userFactory := InitializeUserFactory()

	app := user_app.NewUserApplicationService(sqlite, userService, userFactory)
	err := app.Register("aaa", "test@sample.hoge")
	if err != nil {
		fmt.Println(err)
		return
	}
}
