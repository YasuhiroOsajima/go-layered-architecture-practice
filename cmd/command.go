package main

import (
	"bufio"
	"fmt"
	"os"

	user_app "go-layered-architecture-practice/pkg/user"
)

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
