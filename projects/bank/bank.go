package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("\nEnter choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Println("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount > 0.0 {
				accountBalance += depositAmount
				fmt.Printf("Deposit successful! Your new balance is: %.02f\n", accountBalance)
			} else {
				fmt.Println("Error: Unable to deposit a negative amount")
			}
		case 3:
			var withdrawAmount float64
			fmt.Println("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount < 0 {
				fmt.Println("Error: Unable to withdraw a negative amount")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Error: Insufficient funds")
			} else {
				accountBalance -= withdrawAmount
				fmt.Printf("Withdraw successful! Your new balance is: %.02f\n", accountBalance)
			}
		case 4:
			fmt.Println("\nGoodbye!")
			return
		}
	}

	// fmt.Println("\nYour choice: ", choice)
}
