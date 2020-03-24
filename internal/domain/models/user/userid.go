package user

import (
	"errors"

	"github.com/google/uuid"
)

type UserId string

func NewUserId(id string) (UserId, error) {
	if id == "" {
		return "", errors.New("userId is not specified")
	}

	return (UserId)(id), nil
}

func NewUserIdRandom(repo UserRepositoryInterface) (UserId, error) {
	var userId UserId

	for {
		randomId, err := uuid.NewRandom()
		if err != nil {
			return userId, err
		}

		userId := (UserId)(randomId.String())
		found, err := repo.Find(userId)
		if err != nil {
			return userId, err
		}

		if found == nil {
			break
		}
	}

	return userId, nil
}
