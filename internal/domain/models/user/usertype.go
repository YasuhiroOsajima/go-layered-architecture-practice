package user

import "errors"

const (
	Normal  = "Normal"
	Premium = "Premium"
)

type userType string

func NewUserType(uType string) (userType, error) {
	if uType != Normal && uType != Premium {
		return "", errors.New("invalid userType is specified")
	}

	return (userType)(uType), nil
}

func newNormalUserType() userType {
	normalUser, _ := NewUserType(Normal)
	return normalUser
}

func newPremiumUserType() userType {
	premiumUser, _ := NewUserType(Premium)
	return premiumUser
}
