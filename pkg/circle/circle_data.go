package circle

import (
	circle_model "go-layered-architecture-practice/internal/domain/models/circle"
	"go-layered-architecture-practice/pkg/user"
)

type CircleData struct {
	Id      string
	Name    string
	Owner   user.UserData
	Members []user.UserData
}

func NewCircleData(circle *circle_model.Circle) CircleData {
	owner := user.NewUserData(circle.Owner())
	var members []user.UserData
	for _, u := range circle.Members() {
		members = append(members, user.NewUserData(u))
	}

	return CircleData{circle.Id().AsString(), circle.Name().AsString(), owner, members}
}
