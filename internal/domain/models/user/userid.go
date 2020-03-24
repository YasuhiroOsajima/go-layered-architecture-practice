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

func NewUserIdRandom() (UserId, error) {
	var userId UserId
	randomId, err := uuid.NewRandom()
	if err != nil {
		return userId, err
	}
	return (UserId)(randomId.String()), nil
}
