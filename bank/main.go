package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "Balance.txt"

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

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Failed to retrieve account balance:", err)
		// panic("Failed to retrieve account balance")
    }

	fmt.Println("Welcome to Go Bank!")
	for {
	fmt.Println("Please Select an Option:")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var option int
	fmt.Scan(&option)
	fmt.Println("Your Option is:", option)

	switch option {
		case 1:
			fmt.Println("Checking Balance...")
            fmt.Printf("Your current balance is: %.2f\n", accountBalance)
			continue
        case 2:
			var depositAmount float64
            fmt.Print("Enter the amount to deposit: ")
            fmt.Scan(&depositAmount)
		if depositAmount <= 0 {
            fmt.Println("Invalid deposit amount. Please enter a positive amount.")
            continue
        }
		    accountBalance += depositAmount
		    fmt.Println("Depositing Money...")
            fmt.Printf("Deposit successful. New balance: %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
			continue
		case 3:
		    var withdrawAmount float64
		    fmt.Print("Enter the amount to withdraw: ")
		    fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance || withdrawAmount <= 0 {
            fmt.Println("Invalid withdrawal amount. Please enter a positive amount.")
            continue
        }
		    accountBalance -= withdrawAmount
		    fmt.Println("Withdrawing Money...")
		    fmt.Printf("Withdrawal successful. New balance: %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
			continue
		case 4:
			fmt.Println("Exiting...")
			fmt.Println("Thanks for Choosing our Bank")
			return
        default:
			fmt.Println("Invalid option. Please try again.")
			continue
		}
	}
}