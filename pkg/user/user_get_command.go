package user

import "errors"

type UserGetCommand struct {
	id   string
	name string
}

func NewUserGetCommand() UserGetCommand {
	return UserGetCommand{"", ""}
}

func (c UserGetCommand) GetId() (string, error) {
	if c.id == "" {
		return c.id, errors.New("user id is not specified")
	}

	return c.id, nil
}

func (c UserGetCommand) GetName() (string, error) {
	if c.name == "" {
		return c.name, errors.New("name is not specified")
	}

	return c.name, nil
}

func (c UserGetCommand) SetId(id string) {
	c.id = id
}

func (c UserGetCommand) SetName(name string) {
	c.name = name
}
