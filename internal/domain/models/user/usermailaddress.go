package user

import "errors"

type UserMailAddress string

func NewUserMailAddress(mailAddress string) (UserMailAddress, error) {
	if mailAddress == "" {
		return "", errors.New("userMailAddress is not specified")
	}

	return (UserMailAddress)(mailAddress), nil
}

func (a UserMailAddress) AsString() string {
	return string(a)
}
