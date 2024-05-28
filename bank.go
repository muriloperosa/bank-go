package main

import "fmt"

func main() {

	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		var deposit float64
		fmt.Print("Your deposit: ")
		fmt.Scan(&deposit)

		accountBalance += deposit
		fmt.Println("Balance updated! New amount:", accountBalance)

	} else if choice == 3 {
		var withdraw float64
		fmt.Print("Withdraw amount: ")
		fmt.Scan(&withdraw)

		accountBalance -= withdraw
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}
