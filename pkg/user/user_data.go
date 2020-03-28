package user

import user_model "go-layered-architecture-practice/internal/domain/models/user"

type UserData struct {
	Id          string
	Name        string
	MailAddress string
}

func NewUserData(user *user_model.User) UserData {
	return UserData{user.Id().AsString(), user.Name().AsString(), user.MailAddress().AsString()}
}
