package main

import (
	"errors"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt"
)

const secretKey = "secretkey"
func generateToken(email string, studentId int64) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "id": studentId,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })
    
    return token.SignedString([]byte(secretKey))  
}
func verifyToken(token string) (int64, error) {
    parsedToken, err := jwt.Parse(token, func (token *jwt.Token) (interface{}, error) {
        _, ok := token.Method.(*jwt.SigningMethodHMAC)
        if !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return []byte(secretKey), nil  // Secret should be consistent
    })

    if err != nil {
        return 0, errors.New("could not parse token")
    }
    if !parsedToken.Valid {
        return 0, errors.New("invalid token")
    }

    claims, ok := parsedToken.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("invalid claims")
    }

    studentId, ok := claims["id"].(float64)
    if !ok {
        return 0, errors.New("invalid student id")
    }
    return int64(studentId), nil
}
