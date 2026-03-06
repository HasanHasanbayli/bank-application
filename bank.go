package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to the bank!")

	for {
		printMenu()

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			checkBalance(accountBalance)
		case 2:
			accountBalance = deposit(accountBalance)
		case 3:
			accountBalance = withdraw(accountBalance)
		case 4:
			fmt.Println("Thank you for using the bank. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("\nWhat would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
}

func checkBalance(balance float64) {
	fmt.Printf("Your account balance is: $%.2f\n", balance)
}

func deposit(balance float64) float64 {
	var amount float64
	fmt.Print("Enter the amount to deposit: ")
	if _, err := fmt.Scan(&amount); err != nil {
		fmt.Println("Invalid input.")
		return balance
	}

	if amount <= 0 {
		fmt.Println("Invalid deposit amount!")
		return balance
	}

	balance += amount
	fmt.Printf("You have deposited $%.2f. Your new balance is: $%.2f\n", amount, balance)
	return balance
}

func withdraw(balance float64) float64 {
	var amount float64
	fmt.Print("Enter the amount to withdraw: ")
	if _, err := fmt.Scan(&amount); err != nil {
		fmt.Println("Invalid input.")
		return balance
	}

	if amount <= 0 {
		fmt.Println("Invalid withdrawal amount!")
		return balance
	}

	if amount > balance {
		fmt.Println("Insufficient funds!")
		return balance
	}

	balance -= amount
	fmt.Printf("You have withdrawn $%.2f. Your new balance is: $%.2f\n", amount, balance)
	return balance
}
