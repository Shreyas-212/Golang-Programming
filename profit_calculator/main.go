package main

import (
	"errors"
	"fmt"
	"os"
)

const calculatoinFile = "calculaton.txt"

func writeCalculationToFile(grossProfit, netProfit, profit, ratio float64) {
	calculationText := fmt.Sprintf("grossProfit:%.2f\nnetProfit:%.2f\nprofit:%.2f\nratio:%.2f\n", grossProfit, netProfit, profit, ratio)
	os.WriteFile(calculatoinFile, []byte(calculationText), 0644)
}
func main() {
	
	revenue, err := userInput("Revenue: ")

	if err != nil {
        fmt.Println(err)
        return
    }

	expenses, err := userInput("Expenses: ")
	
	if err != nil {
        fmt.Println(err)
        return
    }
	taxes, err := userInput("Taxes: ")

	if err != nil {
        fmt.Println(err)
        return
    }

	grossProfit, netProfit, profit, ratio := calculations(revenue, expenses, taxes)

	writeCalculationToFile(grossProfit, netProfit, profit, ratio)

	fmt.Printf("Gross Profit: %.2f\n", grossProfit)
	fmt.Printf("Profit: %.2f\n", profit)
    fmt.Printf("Net Profit: %.2f\n", netProfit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)
}

func calculations(revenue, expenses, taxes float64) (float64, float64, float64, float64) {
	grossProfit := revenue - expenses 
	netProfit := revenue - expenses - taxes
	profit := grossProfit * (1 - taxes / 100)
    ratio := grossProfit / profit
	return grossProfit, netProfit, profit, ratio
}

func userInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Invalid user input")
	}
	return userInput, nil
}