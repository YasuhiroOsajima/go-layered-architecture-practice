package circle

// Result object is to avoid dependencies to client layer.
// This interface may express Web framework's return object.
// Result object separates application service data object and clients.

type CircleGetResultInterface interface {
	Status(int)
	JSON(int, interface{})
}
