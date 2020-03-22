package main

import (
	"fmt"

	user_model "go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
	user_repo "go-layered-architecture-practice/internal/repository/sqlite/user"
)

func main() {
	userName1, _ := user_model.NewUserName("aaa")
	userId1, _ := user_model.NewUserId("1")
	userType1, _ := user_model.NewUserType(user_model.Normal)
	newUser1 := user_model.NewUser(userId1, userName1, userType1)

	userName2, _ := user_model.NewUserName("aaa")
	userId2, _ := user_model.NewUserId("1")
	userType2, _ := user_model.NewUserType(user_model.Normal)
	newUser2 := user_model.NewUser(userId2, userName2, userType2)

	newUser1.Equals(newUser2)

	fmt.Println(userName1 == userName2)

	sqlite, err := user_repo.NewUserRepository()
	if err != nil {
		fmt.Println(err)
		return
	}
	userService := services.NewUserService(sqlite)
	result, err := userService.Exists(newUser1)
	fmt.Println(result)
	fmt.Println(err)
}
