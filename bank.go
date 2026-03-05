package main

import "fmt"

func main() {

	accountBalance := 1000.0

	fmt.Println("Welcome to the bank!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Your account balance is: $%.2f\n", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Enter the amount to deposit: ")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Printf("You have deposited $%.2f. Your new balance is: $%.2f\n", depositAmount, accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Enter the amount to withdraw: ")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds!")
		} else {
			accountBalance -= withdrawAmount
			fmt.Printf("You have withdrawn $%.2f. Your new balance is: $%.2f\n", withdrawAmount, accountBalance)
		}
	} else {
		fmt.Println("Thank you for using the bank. Goodbye!")
	}
}
