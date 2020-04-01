package circle

import (
	"go-layered-architecture-practice/internal/domain/models/circle"
)

type CircleRepository struct {
	db []*circle.Circle
}

func NewCircleRepository() *CircleRepository {
	var db []*circle.Circle
	return &CircleRepository{db}
}

func (r *CircleRepository) Save(targetCircle *circle.Circle) error {
	targetIndex := -1
	for i, c := range r.db {
		if targetCircle.Id() == c.Id() {
			targetIndex = i
			break
		}
	}

	if targetIndex >= 0 {
		r.db[targetIndex] = targetCircle
		return nil
	} else {
		r.db = append(r.db, targetCircle)
		return nil
	}
}

func (r CircleRepository) Find(targetCircleId circle.CircleId) (*circle.Circle, error) {
	var targetCircle *circle.Circle
	for _, c := range r.db {
		if c.Id() == targetCircleId {
			targetCircle = circle.NewCircle(c.Id(), c.Name(), c.Owner(), c.Members())
		}
	}
	return targetCircle, nil
}

func (r CircleRepository) FindAll(targetCircleName circle.CircleName) ([]*circle.Circle, error) {
	var circles []*circle.Circle
	for _, c := range r.db {
		if c.Name() == targetCircleName {
			copycircle := circle.NewCircle(c.Id(), c.Name(), c.Owner(), c.Members())
			circles = append(circles, copycircle)
		}
	}
	return circles, nil
}
