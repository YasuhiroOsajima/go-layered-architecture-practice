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

func (c CircleApplicationService) Create(name, userId string) error {
	ownerId, err := user.NewUserId(userId)
	if err != nil {
		return err
	}

	owner, err := c.userRepository.Find(ownerId)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("specified owner user is not exists")
	}

	circleName, err := circle_model.NewCircleName(name)
	if err != nil {
		return err
	}

	newCircle, err := c.circleFactory.Create(circleName, owner)
	if err != nil {
		return err
	}

	exists, err := c.circleService.Exists(newCircle)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("same name circle is already exists")

	}

	err = c.circleRepository.Save(newCircle)
	if err != nil {
		return err
	}

	return nil
}

func (c CircleApplicationService) Update(id, name string) error {
	circleId, err := circle_model.NewCircleId(id)
	if err != nil {
		return err
	}

	circle, err := c.circleRepository.Find(circleId)
	if err != nil {
		return err
	}

	circleName, err := circle_model.NewCircleName(name)
	if err != nil {
		return err
	}

	circle.ChangeName(circleName)
	found, err := c.circleService.Exists(circle)
	if err != nil {
		return err
	}
	if found {
		return errors.New("same name circle is already exists")
	}

	err = c.circleRepository.Save(circle)
	if err != nil {
		return err
	}

	return nil
}

func (c CircleApplicationService) Join(circleId, userId string) error {
	targetUserId, err := user.NewUserId(userId)
	if err != nil {
		return err
	}

	targetUser, err := c.userRepository.Find(targetUserId)
	if err != nil {
		return err
	}
	if targetUser == nil {
		return errors.New("specified user is not exists")
	}

	targetCircleId, err := circle_model.NewCircleId(circleId)
	if err != nil {
		return err
	}

	targetCircle, err := c.circleRepository.Find(targetCircleId)
	if err != nil {
		return err
	}
	if targetCircle == nil {
		return errors.New("specified circle is not exists")
	}

	if targetCircle.IsFull() {
		return errors.New("target circle has full of members")
	}

	err = targetCircle.Join(targetUser)
	if err != nil {
		return err
	}

	err = c.circleRepository.Save(targetCircle)
	if err != nil {
		return err
	}

	return nil
}

func (c CircleApplicationService) Get(command CircleGetCommand) (CircleData, error) {
	var circleData CircleData

	id, idErr := command.GetId()
	name, nameErr := command.GetName()

	if idErr != nil {
		circleId, err := circle_model.NewCircleId(id)
		if err != nil {
			return circleData, err
		}

		circle, err := c.circleRepository.Find(circleId)
		if err != nil {
			return circleData, err
		}

		if circle == nil {
			return circleData, errors.New("target circle is not found")
		}

		circleData = NewCircleData(circle)

	} else if nameErr != nil {
		circleName, err := circle_model.NewCircleName(name)
		if err != nil {
			return circleData, err
		}

		circles, err := c.circleRepository.FindAll(circleName)
		if err != nil {
			return circleData, err
		}
		if len(circles) == 0 {
			return circleData, errors.New("target circle is not found")
		}

		if len(circles) != 1 {
			return circleData, errors.New("target circle name is duplicated")
		}

		circle := circles[0]
		circleData = NewCircleData(circle)

	} else {
		return circleData, errors.New("both arguments were not specified")
	}

	return circleData, nil
}
