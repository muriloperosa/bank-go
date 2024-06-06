package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const file string = "balance.txt"

func main() {

	accountBalance, err := fileops.GetFloatFromFile(file, 0)

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("-----")
		// panic(err) // stop execution and show errors
	}
	for {

		presentOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var deposit float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += deposit
			fileops.WriteFloatFile(file, accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			var withdraw float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			} else if accountBalance-withdraw < 0 {
				fmt.Println("Invalid amount. You don't have enough balance.")
				continue
			}

			accountBalance -= withdraw
			fileops.WriteFloatFile(file, accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
