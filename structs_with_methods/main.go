package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User)outputUserValues() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User 

	appUser = User{
		firstName: userFirstName,
        lastName:  userLastName,
        birthDate: userBirthDate,
        createdAt: time.Now(),
	}

	appUser.outputUserValues()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
