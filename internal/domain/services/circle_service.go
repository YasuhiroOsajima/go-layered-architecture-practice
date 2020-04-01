package services

import "go-layered-architecture-practice/internal/domain/models/circle"

type CircleService struct {
	repository circle.CircleRepositoryInterface
}

func NewCircleService(repository circle.CircleRepositoryInterface) CircleService {
	return CircleService{repository}
}

func (s CircleService) Exists(targetCircle *circle.Circle) (bool, error) {
	sameNames, err := s.repository.FindAll(targetCircle.Name())
	if err != nil {
		return false, err
	}

	if len(sameNames) != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
