package user

import "errors"

type userId string

func NewUserId(id string) (userId, error) {
	if id == "" {
		return "", errors.New("userId is not specified")
	}

	return (userId)(id), nil
}
