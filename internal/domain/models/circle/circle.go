package circle

import (
	"errors"

	"go-layered-architecture-practice/internal/domain/models/user"
)

type Circle struct {
	id      CircleId
	name    CircleName
	owner   *user.User
	members []*user.User
}

func NewCircle(id CircleId, name CircleName, owner *user.User, members []*user.User) *Circle {
	return &Circle{id, name, owner, members}
}

func (c Circle) Equals(user *Circle) bool {
	return c.id == user.id
}

func (c Circle) Id() CircleId {
	return c.id
}

func (c Circle) Name() CircleName {
	return c.name
}

func (c Circle) ChangeName(name CircleName) {
	c.name = name
}

func (c Circle) Owner() *user.User {
	return c.owner
}

func (c Circle) OwnerId() user.UserId {
	return c.owner.Id()
}

func (c Circle) Members() []*user.User {
	return c.members
}

func (c Circle) MemberIds() []user.UserId {
	var idList []user.UserId
	for _, u := range c.members {
		idList = append(idList, u.Id())
	}

	return idList
}

func (c Circle) IsFull() bool {
	return c.CountMembers() >= 30
}

func (c Circle) CountMembers() int {
	return len(c.members) + 1
}

func (c *Circle) Join(user *user.User) error {
	if c.IsFull() {
		return errors.New("circle has full count of members")
	}
	c.members = append(c.members, user)
	return nil
}
