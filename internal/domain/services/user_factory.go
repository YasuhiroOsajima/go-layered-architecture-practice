package services

import (
	"github.com/google/uuid"

	"go-layered-architecture-practice/internal/domain/models/user"
)

type UserFactory struct {
	repository user.UserRepositoryInterface
}

func NewUserFactory(repository user.UserRepositoryInterface) UserFactory {
	return UserFactory{repository}
}

func (f UserFactory) Create(name user.UserName, mailAddress user.UserMailAddress) (*user.User, error) {
	var userId *user.UserId

	for {
		randomId, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		randomUserId := (user.UserId)(randomId.String())
		found, err := f.repository.Find(randomUserId)
		if err != nil {
			return nil, err
		}

		if found == nil {
			userId = &randomUserId
			break
		}
	}

	userType := user.NewUserNormal()
	newUser := user.NewUser(*userId, name, mailAddress, userType)

	return newUser, nil
}
