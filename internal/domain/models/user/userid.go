package user

import "errors"

type UserId struct {
	value string
}

func NewUserId(id string) (*UserId, error) {
	if id == "" {
		return nil, errors.New("UserId is not specified")
	}

	return &UserId{id}, nil
}

func (u UserId) Value() string {
	return u.value
}
