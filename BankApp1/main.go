package main

import (
	"fmt"
)


type Account struct {
	Name    string
	Balance float64
}


var accounts = make(map[string]*Account)


func createAccount(name string) {
	if _, exists := accounts[name]; exists {
		fmt.Println(" Account already exists!")
		return
	}
	accounts[name] = &Account{Name: name, Balance: 0}
	fmt.Println("Account created successfully!")
}


func deposit(name string, amount float64) {
	acc, ok := accounts[name]
	if !ok {
		fmt.Println("Account not found!")
		return
	}
	if amount <= 0 {
		fmt.Println("Invalid amount.")
		return
	}
	acc.Balance += amount
	fmt.Printf("%.2f deposited. New Balance: %.2f\n", amount, acc.Balance)
}


func viewBalance(name string) {
	acc, ok := accounts[name]
	if !ok {
		fmt.Println("Account not found!")
		return
	}
	fmt.Printf("%s's Balance: %.2f\n", name, acc.Balance)
}


func transfer(from, to string, amount float64) {
	sender, ok1 := accounts[from]
	receiver, ok2 := accounts[to]

	if !ok1 || !ok2 {
		fmt.Println("One or both accounts not found!")
		return
	}
	if amount <= 0 {
		fmt.Println("Invalid amount.")
		return
	}
	if sender.Balance < amount {
		fmt.Println("Not enough balance!")
		return
	}

	sender.Balance -= amount
	receiver.Balance += amount
	fmt.Printf("%.2f transferred from %s to %s\n", amount, from, to)
}

func main() {
	for {
		fmt.Println("\n Simple Bank App ")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. View Balance")
		fmt.Println("4. Transfer")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			createAccount(name)

		case 2:
			var name string
			var amount float64
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter amount: ")
			fmt.Scanln(&amount)
			deposit(name, amount)

		case 3:
			var name string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			viewBalance(name)

		case 4:
			var from, to string
			var amount float64
			fmt.Print("Enter sender name: ")
			fmt.Scanln(&from)
			fmt.Print("Enter receiver name: ")
			fmt.Scanln(&to)
			fmt.Print("Enter amount: ")
			fmt.Scanln(&amount)
			transfer(from, to, amount)

		case 5:
			fmt.Println("Thank you for using Simple Bank App!")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
