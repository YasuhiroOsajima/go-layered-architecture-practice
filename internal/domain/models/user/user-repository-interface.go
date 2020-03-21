package user

type UserRepositoryInterface interface {
	Save(*User) error
	Find(*UserName) (*User, error)
}
