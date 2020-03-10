package circle

import "errors"

type circleName string

func NewClassName(name string) (circleName, error) {
	if name == "" {
		return "", errors.New("circleName is not specified")
	}

	return (circleName)(name), nil
}
