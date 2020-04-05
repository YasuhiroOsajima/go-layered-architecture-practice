package circle

// Result object is to avoid dependencies to client layer.
// This interface may express Web framework's return object.
// Result object separates application service data object and clients.

import "errors"

type CircleGetCommand struct {
	id   string
	name string
}

func NewCircleGetCommand() CircleGetCommand {
	return CircleGetCommand{"", ""}
}

func (c CircleGetCommand) GetId() (string, error) {
	if c.id == "" {
		return c.id, errors.New("circle id is not specified")
	}

	return c.id, nil
}

func (c CircleGetCommand) GetName() (string, error) {
	if c.name == "" {
		return c.name, errors.New("circle is not specified")
	}

	return c.name, nil
}

func (c CircleGetCommand) SetId(id string) {
	c.id = id
}

func (c CircleGetCommand) SetName(name string) {
	c.name = name
}
