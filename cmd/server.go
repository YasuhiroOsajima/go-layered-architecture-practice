package main

import "go-layered-architecture-practice/internal/domain/models/user"

func main() {
	userName, _ := user.NewUserName("")
	userId, _ := user.NewUserId("")
	userType, _ := user.NewUserType(user.Normal)
	newUser := user.NewUser(userId, userName, userType)
	println(newUser)
}
