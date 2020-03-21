package user

import "errors"

type UserId string

func NewUserId(id string) (UserId, error) {
	if id == "" {
		return "", errors.New("userId is not specified")
	}

	return (UserId)(id), nil
}
