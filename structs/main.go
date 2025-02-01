package main

import (
	"fmt"
	"structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser *user.User
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
        fmt.Println(err)
        return
    }

	admin := user.NewAdmin("example.com", "example")

	admin.User.OutputUserValues()

	appUser.OutputUserValues()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
