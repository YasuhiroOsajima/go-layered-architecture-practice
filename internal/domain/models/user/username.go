package user

import "errors"

type UserName struct {
	value string
}

func NewUserName(name string) (*UserName, error) {
	if name == "" {
		return nil, errors.New("UserName is not specified")
	}

	if len(name) < 3 {
		return nil, errors.New("UserName should be over 3 characters")
	}

	if len(name) > 20 {
		return nil, errors.New("UserName should be less than 20 characters")
	}

	return &UserName{name}, nil
}

func (u UserName) Value() string {
	return u.value
}
