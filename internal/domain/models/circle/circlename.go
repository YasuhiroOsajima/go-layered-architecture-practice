package circle

import "errors"

type CircleName string

func NewClassName(name string) (CircleName, error) {
	if name == "" {
		return "", errors.New("circleName is not specified")
	}

	return (CircleName)(name), nil
}

func (n CircleName) AsString() string {
	return string(n)
}
