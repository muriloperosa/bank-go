package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const file string = "balance.txt"

func main() {

	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("-----")
		// panic(err) // stop execution and show errors
	}
	for {

		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(file, []byte(balanceTxt), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return 0, errors.New("Failed to read the balance file.")
	}

	balanceTxt := string(data)
	balance, err := strconv.ParseFloat(balanceTxt, 64)

	if err != nil {
		return 0, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}
