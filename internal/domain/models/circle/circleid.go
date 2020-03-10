package circle

import "errors"

type circleId string

func NewCircleId(id string) (circleId, error) {
	if id == "" {
		return "", errors.New("circleId is not specified")
	}

	return (circleId)(id), nil
}
