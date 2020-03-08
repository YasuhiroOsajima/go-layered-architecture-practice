package user

import "errors"

type userName string

func NewUserName(name string) (userName, error) {
	if name == "" {
		return "", errors.New("userName is not specified")
	}

	if len(name) < 3 {
		return "", errors.New("userName should be over 3 characters")
	}

	if len(name) > 20 {
		return "", errors.New("userName should be less than 20 characters")
	}

	return (userName)(name), nil
}
