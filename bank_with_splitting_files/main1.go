package main 

import (
    "errors"
    "fmt"
    "os"
	"strconv"
)

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 00, errors.New("Failed to read balance from file")
	}
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance, nil
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}