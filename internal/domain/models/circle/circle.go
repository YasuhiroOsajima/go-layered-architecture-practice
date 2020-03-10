package circle

import "go-layered-architecture-practice/internal/domain/models/user"

type circle struct {
	id      circleId
	name    circleName
	owner   user.User
	members []user.User
}

func NewCircle(id circleId, name circleName, owner user.User, members []user.User) *circle {
	return &circle{id, name, owner, members}
}

func (c circle) IsFull() bool {
	return c.CountMembers() >= 30
}

func (c circle) CountMembers() int {
	return len(c.members) + 1
}
