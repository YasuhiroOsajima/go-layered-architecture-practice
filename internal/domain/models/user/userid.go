package user

import "errors"

type userId struct {
	value string
}

func NewUserId(id string) (*userId, error) {
	if id == "" {
		return nil, errors.New("userId is not specified")
	}

	return &userId{id}, nil
}

func (u userId) Value() string {
	return u.value
}
