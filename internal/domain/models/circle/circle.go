package circle

import (
	"errors"

	"go-layered-architecture-practice/internal/domain/models/user"
)

type circle struct {
	id      circleId
	name    circleName
	owner   user.User
	members []user.User
}

func NewCircle(id circleId, name circleName, owner user.User, members []user.User) *circle {
	return &circle{id, name, owner, members}
}

func (c circle) Id() circleId {
	return c.id
}

func (c circle) Name() circleName {
	return c.name
}

func (c circle) Owner() user.User {
	return c.owner
}

func (c circle) Members() []user.User {
	return c.members
}

func (c circle) IsFull() bool {
	return c.CountMembers() >= 30
}

func (c circle) CountMembers() int {
	return len(c.members) + 1
}

func (c *circle) Join(user user.User) error {
	if c.IsFull() {
		return errors.New("circle has full count of members")
	}
	c.members = append(c.members, user)
	return nil
}
