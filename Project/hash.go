package main

import "golang.org/x/crypto/bcrypt"

func hashPassword(pasword string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(pasword), 15)
	return string(bytes), err
}

func compareHashAndPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
