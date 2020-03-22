package user

type UserRepositoryInterface interface {
	Save(*User) error
	Find(UserId) (*User, error)
	FindAll(UserName) ([]*User, error)
	Delete(*User) error
}
