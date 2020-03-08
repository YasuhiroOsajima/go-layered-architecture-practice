package user

type user struct {
	id       userId
	name     userName
	userType userType
}

func NewUser(id userId, name userName, userTpe userType) *user {
	return &user{id, name, userTpe}
}

func (u user) IsPremium() bool {
	premiumUser, _ := NewUserType(premium)
	return u.userType == premiumUser
}

func (u *user) ChangeName(name userName) {
	u.name = name
}
