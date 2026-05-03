package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}

func main() {

	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter expected rate of return: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	var fv, rfv = calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)
	fmt.Printf("Future value: %.2f\n", fv)
	fmt.Printf("Future real value: %.2f\n", rfv)

}
