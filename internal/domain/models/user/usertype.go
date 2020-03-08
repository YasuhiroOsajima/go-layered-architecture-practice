package user

import "errors"

const (
	normal  = "Normal"
	premium = "Premium"
)

type UserType struct {
	value string
}

func NewUserType(userType string) (*UserType, error) {
	if userType != normal && userType != premium {
		return nil, errors.New("invalid UserType is specified")
	}

	return &UserType{userType}, nil
}

func (u UserType) Value() string {
	return u.value
}

func (u UserType) Normal() string {
	return normal
}

func (u UserType) Premium() string {
	return premium
}
