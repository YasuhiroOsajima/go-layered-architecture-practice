package user

import "errors"

const (
	normal  = "Normal"
	premium = "Premium"
)

type userType string

func NewUserType(uType string) (userType, error) {
	if uType != normal && uType != premium {
		return "", errors.New("invalid userType is specified")
	}

	return (userType)(uType), nil
}

func (u userType) Normal() string {
	return normal
}

func (u userType) Premium() string {
	return premium
}
