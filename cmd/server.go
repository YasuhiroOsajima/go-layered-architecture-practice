package main

import (
	"go-layered-architecture-practice/internal/domain/models/user"
)

func main() {
	userName1, _ := user.NewUserName("aaa")
	userId1, _ := user.NewUserId("1")
	userType1, _ := user.NewUserType(user.Normal)
	newUser1 := user.NewUser(userId1, userName1, userType1)

	userName2, _ := user.NewUserName("aaa")
	userId2, _ := user.NewUserId("1")
	userType2, _ := user.NewUserType(user.Normal)
	newUser2 := user.NewUser(userId2, userName2, userType2)

	newUser1.Equals(newUser2)
}
