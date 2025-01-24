package diff

import (

    "fmt"
    "os"
	"strconv"
)
const accountBalanceFile = "Balance.txt"
func GetBalanceFromFile() (variable float64, error string) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 0.0,"Failed to read balance from file"
	}
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance, error
}
func WriteBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}