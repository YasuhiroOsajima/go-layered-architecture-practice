package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	user_app "go-layered-architecture-practice/pkg/user"
)

// Command data
type CircleGetCommand struct {
	id           string
	name         string
	ReturnStatus int
	ReturnJSON   interface{}
}

func NewCircleGetCommand() CircleGetCommand {
	return CircleGetCommand{"", "", 0, nil}
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

func (c *CircleGetCommand) Status(status int) {
	c.ReturnStatus = status
}
func (c *CircleGetCommand) JSON(status int, data interface{}) {
	c.ReturnStatus = status
	returnJSON, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	c.ReturnJSON = returnJSON
}

type UserGetCommand struct {
	id           string
	name         string
	ReturnStatus int
	ReturnJSON   interface{}
}

func NewUserGetCommand() UserGetCommand {
	return UserGetCommand{"", "", 0, nil}
}

func (u UserGetCommand) GetId() (string, error) {
	if u.id == "" {
		return u.id, errors.New("user id is not specified")
	}

	return u.id, nil
}

func (u UserGetCommand) GetName() (string, error) {
	if u.name == "" {
		return u.name, errors.New("name is not specified")
	}

	return u.name, nil
}

func (u UserGetCommand) SetId(id string) {
	u.id = id
}

func (u UserGetCommand) SetName(name string) {
	u.name = name
}

func (u *UserGetCommand) Status(status int) {
	u.ReturnStatus = status
}
func (u *UserGetCommand) JSON(status int, data interface{}) {
	u.ReturnStatus = status
	returnJSON, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	u.ReturnJSON = returnJSON
}

func main() {
	sqlite := InitializeUserRepository()
	userService := InitializeUserService()
	userFactory := InitializeUserFactory()
	app := user_app.NewUserApplicationService(sqlite, userService, userFactory)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Input user name")
		fmt.Print("> ")
		scanner.Scan()
		userName := scanner.Text()

		fmt.Println("Input mail address")
		fmt.Print("> ")
		scanner.Scan()
		mailAddress := scanner.Text()

		err := app.Register(userName, mailAddress)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("--------------------------------")
		fmt.Println("user created")
		fmt.Println("--------------------------------")
		fmt.Println("user name")
		fmt.Println("- " + userName)
		fmt.Println("--------------------------------")

		fmt.Println("continue? (y/n)")
		fmt.Print("> ")
		scanner.Scan()
		yOrN := scanner.Text()
		if yOrN == "n" {
			break
		}
	}
}
