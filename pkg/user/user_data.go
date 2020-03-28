package user

import user_model "go-layered-architecture-practice/internal/domain/models/user"

type UserData struct {
	id   string
	name string
}

func NewUserData(user *user_model.User) UserData {
	return UserData{user.Id().AsString(), user.Name().AsString()}
}

func (u UserData) Id() string {
	return u.id
}

func (u UserData) Name() string {
	return u.name
}
