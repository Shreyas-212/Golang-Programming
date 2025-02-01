package user

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

type Admin struct {
	email string
	password string
	User User
}

func (u User)OutputUserValues() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
	    password: password,
		User: User{
			firstName: "Admin",
            lastName:  "Admin",
            birthDate: "",
            createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
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
