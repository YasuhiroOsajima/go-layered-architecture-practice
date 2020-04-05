package server

import (
	"fmt"

	"go-layered-architecture-practice/cmd/share"
	"go-layered-architecture-practice/pkg/user"
)

var app user.UserApplicationService

func init() {
	sqlite := share.InitializeUserRepository()
	userService := share.InitializeUserService()
	userFactory := share.InitializeUserFactory()
	app = user.NewUserApplicationService(sqlite, userService, userFactory)
}

func RegisterUser(c Context) {
	err := app.Register("aaa", "test@sample.hoge")
	if err != nil {
		fmt.Println(err)
	}
}
