package circle

import "go-layered-architecture-practice/pkg/user"

// Result object is to avoid dependencies to application service objects by client layer.

type CircleGetResult struct {
	Id      string
	Name    string
	Owner   user.UserGetResult
	Members []user.UserGetResult
}

func NewCircleGetResult(circle CircleData) CircleGetResult {
	owner := user.NewUserGetResult(circle.Owner)

	var members []user.UserGetResult
	for _, m := range circle.Members {
		member := user.NewUserGetResult(m)
		members = append(members, member)
	}

	return CircleGetResult{circle.Id, circle.Name, owner, members}
}
