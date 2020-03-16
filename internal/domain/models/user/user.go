package user

type User struct {
	id       userId
	name     userName
	userType userType
}

func NewUser(id userId, name userName, userTpe userType) *User {
	return &User{id, name, userTpe}
}

func (u User) Equals(user *User) bool {
	return u.id == user.id
}

func (u User) Id() userId {
	return u.id
}

func (u User) Name() userName {
	return u.name
}

func (u User) Type() userType {
	return u.userType
}

func (u User) IsPremium() bool {
	premiumUser := newPremiumUserType()
	return u.userType == premiumUser
}

func (u *User) ChangeName(name userName) {
	u.name = name
}

func (u *User) Upgrade() {
	u.userType = newPremiumUserType()
}

func (u *User) Downgrade() {
	u.userType = newNormalUserType()
}
