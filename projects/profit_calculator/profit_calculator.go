// Package declaration
package main

// Import packages
import (
	"fmt"
)

func calculateEBT(revenue, expenses float64) float64 {
	return revenue - expenses
}

func calculateProfit(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	return calculateEBT(revenue, expenses), calculateProfit(ebt, taxRate), calculateRatio(ebt, profit)
}

func getUserInput(prompt string) float64 {
	var userInput float64
	fmt.Print(prompt)
	fmt.Scan(&userInput)
	return userInput
}

func displayFinanceReport(ebt, profit, ratio float64) {

}
func main() {
	// Variable declarations
	// var revenue, expenses, taxRate float64

	// revenue = getUserInput("Enter revenue: ")
	// expenses = getUserInput("Enter expenses: ")
	// taxRate = float64(getUserInput("Enter tax rate: "))

	// Calculate earnings before tax (EBT)
	// ebt := calculateEBT(revenue, expenses)
	// Calculate earnings after tax (profit)
	// profit := calculateProfit(ebt, taxRate)
	// Calculate ratio (EBT / profit)
	// ratio := calculateRatio(ebt, profit)

	displayFinanceReport(getUserInput("Enter revenue: "), getUserInput("Enter expenses: "), getUserInput("Enter tax rate: "))

	// Output calculations
	// fmt.Printf("Earnings before tax (EBT): $%.2f\n", ebt)
	// fmt.Printf("Earnings after tax (profit): $%.2f\n", profit)
	// fmt.Printf("Ratio: %.4f\n", ratio)
}
