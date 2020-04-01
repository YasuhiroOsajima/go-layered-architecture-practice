package circle

import "errors"

type CircleId string

func NewCircleId(id string) (CircleId, error) {
	if id == "" {
		return "", errors.New("circleId is not specified")
	}

	return (CircleId)(id), nil
}

func (i CircleId) AsString() string {
	return string(i)
}
