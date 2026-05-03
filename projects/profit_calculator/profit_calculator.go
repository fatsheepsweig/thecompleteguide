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

func main() {
	// Variable declarations
	var revenue, expenses, taxRate float64

	// Prompt user for revenue
	fmt.Print("Enter revenue: ")
	// Read user input from stdin
	fmt.Scan(&revenue)
	// Prompt user for expenses
	fmt.Print("Enter expenses: ")
	// Read user input from stdin
	fmt.Scan(&expenses)
	// Prompt user for tax rate
	fmt.Print("Enter tax rate: ")
	// Read user input from stdin
	fmt.Scan(&taxRate)

	// Calculate earnings before tax (EBT)
	ebt := calculateEBT(revenue, expenses)
	// Calculate earnings after tax (profit)
	profit := calculateProfit(ebt, taxRate)
	// Calculate ratio (EBT / profit)
	ratio := calculateRatio(ebt, profit)

	// Output calculations
	fmt.Printf("Earnings before tax (EBT): $%.2f\n", ebt)
	fmt.Printf("Earnings after tax (profit): $%.2f\n", profit)
	fmt.Printf("Ratio: %.4f\n", ratio)
}
