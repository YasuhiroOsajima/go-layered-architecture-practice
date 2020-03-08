package user

// User is entity of user domain.
type User struct {
	id       UserId
	name     UserName
	userType UserType
}

// NewUser is constructor of User.
func NewUser(id UserId, name UserName, userTpe UserType) *User {
	return &User{id, name, userTpe}
}
