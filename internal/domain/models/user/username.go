package user

import "errors"

type userName struct {
	value string
}

func NewUserName(name string) (*userName, error) {
	if name == "" {
		return nil, errors.New("userName is not specified")
	}

	if len(name) < 3 {
		return nil, errors.New("userName should be over 3 characters")
	}

	if len(name) > 20 {
		return nil, errors.New("userName should be less than 20 characters")
	}

	return &userName{name}, nil
}

func (u userName) Value() string {
	return u.value
}
