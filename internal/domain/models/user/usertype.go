package user

import "errors"

const (
	Normal  = "Normal"
	Premium = "Premium"
)

type UserType string

func NewUserType(uType string) (UserType, error) {
	if uType != Normal && uType != Premium {
		return "", errors.New("invalid userType is specified")
	}

	return (UserType)(uType), nil
}

func newNormalUserType() UserType {
	normalUser, _ := NewUserType(Normal)
	return normalUser
}

func newPremiumUserType() UserType {
	premiumUser, _ := NewUserType(Premium)
	return premiumUser
}
