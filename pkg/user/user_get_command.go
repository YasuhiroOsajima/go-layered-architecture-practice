package user

type UserGetCommandInterface interface {
	GetId() (string, error)
	GetName() (string, error)
}
