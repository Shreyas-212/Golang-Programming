package main 

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years, expectedReturnRate float64

	fmt.Print("Investment Amount: ") 
	fmt.Scan(&investmentAmount)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	fmt.Print("Expected Return Rate (%): ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	var futureRealValue = futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Value (adjusted for inflation):", futureRealValue)
}
