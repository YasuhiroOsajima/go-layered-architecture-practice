package user

type UserGetResultInterface interface {
	Status(int)
	JSON(int, interface{})
}
