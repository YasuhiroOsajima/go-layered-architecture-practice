package circle

import (
	"errors"

	"go-layered-architecture-practice/internal/domain/models/user"
)

type Circle struct {
	id      circleId
	name    circleName
	owner   user.User
	members []user.User
}

func NewCircle(id circleId, name circleName, owner user.User, members []user.User) *Circle {
	return &Circle{id, name, owner, members}
}

func (u Circle) Equals(user *Circle) bool {
	return u.id == user.id
}

func (c Circle) Id() circleId {
	return c.id
}

func (c Circle) Name() circleName {
	return c.name
}

func (c Circle) Owner() user.User {
	return c.owner
}

func (c Circle) Members() []user.User {
	return c.members
}

func (c Circle) IsFull() bool {
	return c.CountMembers() >= 30
}

func (c Circle) CountMembers() int {
	return len(c.members) + 1
}

func (c *Circle) Join(user user.User) error {
	if c.IsFull() {
		return errors.New("circle has full count of members")
	}
	c.members = append(c.members, user)
	return nil
}
