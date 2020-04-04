package circle

type CircleGetCommandInterface interface {
	GetId() (string, error)
	GetName() (string, error)
	Status(int)
	JSON(int, interface{})
}
