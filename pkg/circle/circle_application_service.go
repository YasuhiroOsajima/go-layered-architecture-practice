package circle

import (
	"errors"
	circle_model "go-layered-architecture-practice/internal/domain/models/circle"
	"go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
)

type CircleApplicationService struct {
	circleRepository circle_model.CircleRepositoryInterface
	circleService    services.CircleService
	circleFactory    services.CircleFactory
	userRepository   user.UserRepositoryInterface
}

func NewCircleApplicationService(circleRepo circle_model.CircleRepositoryInterface, service services.CircleService, factory services.CircleFactory, userRepo user.UserRepositoryInterface) CircleApplicationService {
	return CircleApplicationService{circleRepo, service, factory, userRepo}
}

func (c CircleApplicationService) Create(name, userId string) (result CircleGetResultInterface) {
	ownerId, err := user.NewUserId(userId)
	if err != nil {
		result.Status(500)
		return result
	}

	owner, err := c.userRepository.Find(ownerId)
	if err != nil {
		result.Status(500)
		return result
	}
	if owner == nil {
		result.JSON(404, errors.New("specified owner user is not exists"))
		return result
	}

	circleName, err := circle_model.NewCircleName(name)
	if err != nil {
		result.Status(500)
		return result
	}

	newCircle, err := c.circleFactory.Create(circleName, owner)
	if err != nil {
		result.Status(500)
		return result
	}

	exists, err := c.circleService.Exists(newCircle)
	if err != nil {
		result.Status(500)
		return result
	}

	if exists {
		result.JSON(400, errors.New("same name circle is already exists"))
		return result

	}

	err = c.circleRepository.Save(newCircle)
	if err != nil {
		result.Status(500)
		return result
	}

	result.Status(200)
	return result
}

func (c CircleApplicationService) Update(id, name string) (result CircleGetResultInterface) {
	circleId, err := circle_model.NewCircleId(id)
	if err != nil {
		result.Status(500)
		return result
	}

	circle, err := c.circleRepository.Find(circleId)
	if err != nil {
		result.Status(500)
		return result
	}

	circleName, err := circle_model.NewCircleName(name)
	if err != nil {
		result.Status(500)
		return result
	}

	circle.ChangeName(circleName)
	found, err := c.circleService.Exists(circle)
	if err != nil {
		result.Status(500)
		return result
	}
	if found {
		result.JSON(400, errors.New("same name circle is already exists"))
		return result
	}

	err = c.circleRepository.Save(circle)
	if err != nil {
		result.Status(500)
		return result
	}

	result.Status(200)
	return result
}

func (c CircleApplicationService) Join(circleId, userId string) (result CircleGetResultInterface) {
	targetUserId, err := user.NewUserId(userId)
	if err != nil {
		result.Status(500)
		return result
	}

	targetUser, err := c.userRepository.Find(targetUserId)
	if err != nil {
		result.Status(500)
		return result
	}
	if targetUser == nil {
		result.JSON(404, errors.New("specified user is not exists"))
		return result
	}

	targetCircleId, err := circle_model.NewCircleId(circleId)
	if err != nil {
		result.Status(500)
		return result
	}

	targetCircle, err := c.circleRepository.Find(targetCircleId)
	if err != nil {
		result.Status(500)
		return result
	}
	if targetCircle == nil {
		result.JSON(404, errors.New("specified circle is not exists"))
		return result
	}

	if targetCircle.IsFull() {
		result.JSON(409, errors.New("target circle has full of members"))
		return result
	}

	err = targetCircle.Join(targetUser)
	if err != nil {
		result.Status(500)
		return result
	}

	err = c.circleRepository.Save(targetCircle)
	if err != nil {
		result.Status(500)
		return result
	}

	result.Status(200)
	return result
}

func (c CircleApplicationService) Get(command CircleGetCommandInterface) (result CircleGetResultInterface) {
	var circleData CircleData

	id, idErr := command.GetId()
	name, nameErr := command.GetName()

	if idErr != nil {
		circleId, err := circle_model.NewCircleId(id)
		if err != nil {
			result.Status(500)
			return result
		}

		circle, err := c.circleRepository.Find(circleId)
		if err != nil {
			result.Status(500)
			return result
		}

		if circle == nil {
			result.JSON(404, errors.New("target circle is not found"))
			return result
		}

		circleData = NewCircleData(circle)

	} else if nameErr != nil {
		circleName, err := circle_model.NewCircleName(name)
		if err != nil {
			result.JSON(500, err)
			return result
		}

		circles, err := c.circleRepository.FindAll(circleName)
		if err != nil {
			result.JSON(500, err)
			return result
		}
		if len(circles) == 0 {
			result.JSON(404, errors.New("target circle is not found"))
			return result
		}

		if len(circles) != 1 {
			result.JSON(400, errors.New("target circle is not found"))
			return result
		}

		circle := circles[0]
		circleData = NewCircleData(circle)
		result.JSON(200, circleData)

	} else {
		result.JSON(400, errors.New("both arguments were not specified"))
		return result
	}

	return result
}

func (c CircleApplicationService) GetRecommended(circleName circle_model.CircleName) (result CircleGetResultInterface) {
	circles, err := c.circleRepository.FindAll(circleName)
	if err != nil {
		result.Status(500)
		return result
	}

	var recommended []*circle_model.Circle
	for _, c := range circles {
		if circle_model.IsRecommended(c) {
			recommended = append(recommended, c)
		}
	}

	result.JSON(200, recommended)
	return result
}
