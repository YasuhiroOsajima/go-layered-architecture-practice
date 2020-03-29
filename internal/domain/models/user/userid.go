package user

import "errors"

type UserId string

func NewUserId(id string) (UserId, error) {
	if id == "" {
		return "", errors.New("userId is not specified")
	}

	return (UserId)(id), nil
}

func (i UserId) AsString() string {
	return string(i)
}
