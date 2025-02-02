package main

import (
	"fmt"
	"time"
	"errors"
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

func newUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("FirstName, LastName and BirthDate are required")
    }

	return &User{
		firstName: firstName,
        lastName:  lastName,
        birthDate: birthDate,
        createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var appUser *User
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
        fmt.Println(err)
        return
    }

	appUser.outputUserValues()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
