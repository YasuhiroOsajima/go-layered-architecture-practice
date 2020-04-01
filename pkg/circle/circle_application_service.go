package circle

import (
	"errors"
	"go-layered-architecture-practice/internal/domain/models/circle"
	"go-layered-architecture-practice/internal/domain/models/user"
	"go-layered-architecture-practice/internal/domain/services"
)

type CircleApplicationService struct {
	circleRepository circle.CircleRepositoryInterface
	circleService    services.CircleService
	circleFactory    services.CircleFactory
	userRepository   user.UserRepositoryInterface
}

func NewCircleApplicationService(circleRepo circle.CircleRepositoryInterface, service services.CircleService, factory services.CircleFactory, userRepo user.UserRepositoryInterface) CircleApplicationService {
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

	circleName, err := circle.NewCircleName(name)
	if err != nil {
		return err
	}

	circle, err := c.circleFactory.Create(circleName, owner)
	if err != nil {
		return err
	}

	exists, err := c.circleService.Exists(circle)
	if err != nil {
		return err
	}

	if exists {
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

	targetCircleId, err := circle.NewCircleId(circleId)
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

	if len(targetCircle.Members()) >= 29 {
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
