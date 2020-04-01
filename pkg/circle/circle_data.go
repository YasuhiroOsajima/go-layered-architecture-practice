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

func NewCircleData(circle *circle_model.Circle, owner user.UserData, members []user.UserData) CircleData {
	return CircleData{circle.Id().AsString(), circle.Name().AsString(), owner, members}
}
