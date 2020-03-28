package user

import "errors"

type UserUpdateCommand struct {
	id          string
	name        string
	mailAddress string
}

func NewUserUpdateCommand(id string) UserUpdateCommand {
	return UserUpdateCommand{id, nil, nil}
}

func (c UserUpdateCommand) GetId() string {
	return c.id
}

func (c UserUpdateCommand) GetName() (string, error) {
	if c.name == "" {
		return c.name, errors.New("name is not specified")
	}

	return c.name, nil
}

func (c UserUpdateCommand) GetMailAddress() (string, error) {
	if c.mailAddress == "" {
		return c.mailAddress, errors.New("mail address is not specified")
	}
	return c.mailAddress, nil
}

func (c UserUpdateCommand) SetName(name string) {
	c.name = name
}

func (c UserUpdateCommand) SetMailAddress(mailAddress string) {
	c.mailAddress = mailAddress
}
