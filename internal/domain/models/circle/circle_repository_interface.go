package circle

type CircleRepositoryInterface interface {
	Save(*Circle) error
	Find(CircleId) (*Circle, error)
	FindAll(CircleName) ([]*Circle, error)
}
