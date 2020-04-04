package user

type UserGetCommandInterface interface {
	GetId() (string, error)
	GetName() (string, error)
	Status(int)
	JSON(int, interface{})
}
