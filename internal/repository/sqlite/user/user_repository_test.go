package user

import (
	"testing"

	user_model "go-layered-architecture-practice/internal/domain/models/user"
)

func TestSqliteUserRepository(t *testing.T) {
	userName1, _ := user_model.NewUserName("xxx")
	userId1, _ := user_model.NewUserId("99")
	userMail1, _ := user_model.NewUserMailAddress("test@sample.hoge")
	userType1, _ := user_model.NewUserType(user_model.Normal)
	newUser1 := user_model.NewUser(userId1, userName1, userMail1, userType1)

	userRepository := NewUserRepository()

	err := userRepository.Save(newUser1)
	if err != nil {
		t.Errorf(err.Error())
	}

	resultUser, err := userRepository.Find(newUser1.Id())
	if err != nil {
		t.Errorf(err.Error())
	}

	if !resultUser.Equals(newUser1) {
		t.Errorf("Not matched.")
	}
}
