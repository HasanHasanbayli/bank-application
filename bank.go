package main

import "fmt"

func main() {

	fmt.Println("Welcome to the bank!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choise int

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choise)

	fmt.Println("Your choise", choise)
}
