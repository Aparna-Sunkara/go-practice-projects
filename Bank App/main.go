package main

import (
	"fmt"
	"strconv"
)

type Account struct {
	Id      int
	Name    string
	balance float64
}
type Bank struct {
	accounts map[int]Account
	nextID   int
}

func NewBank() Bank {
	return Bank{
		accounts: make(map[int]Account),
		nextID:   1,
	}
}
func (b Bank) CreateAccount(name string) Account {
	acc := Account{
		Id:      b.nextID,
		Name:    name,
		balance: 0,
	}
	b.accounts[acc.Id] = acc
	b.nextID++
	return acc
}
func (b Bank) ViewBalance(id int) (float64, error) {
	acc, ok := b.accounts[id]
	if !ok {
		return 0, fmt.Errorf("account not found")
	}
	return acc.balance, nil
}
func (b Bank) Deposit(id int, amount float64) error {
	acc, ok := b.accounts[id]
	if !ok {
		return fmt.Errorf("account not found")
	}
	if amount <= 0 {
		return fmt.Errorf("invalid amount")
	}
	acc.balance += amount
	return nil
}
func (b Bank) Transfer(fromID, toID int, amount float64) error {
	from, ok1 := b.accounts[fromID]
	to, ok2 := b.accounts[toID]
	if !ok1 || !ok2 {
		return fmt.Errorf("account not found")
	}
	if amount <= 0 {
		return fmt.Errorf("invalid amount")
	}
	if from.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	from.balance -= amount
	to.balance += amount
	return nil
}
func main() {
	bank := NewBank()

	for {
		fmt.Println("\n Bank App")
		fmt.Println("1. Create Account")
		fmt.Println("2. View Balance")
		fmt.Println("3. Deposit Amount")
		fmt.Println("4. Transfer Balance")
		fmt.Println("5. Exit")

		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			fmt.Println("Enter account holder name: ")
			var name string
			fmt.Scanln(&name)
			acc := bank.CreateAccount(name)
			fmt.Printf("Account created! ID: %d, Name: %s, Balance: %.2f\n",
				acc.Id, acc.Name, acc.balance)
		case "2":
			fmt.Println("Enter account ID")
			var id string
			fmt.Scanln(&id)
			idx, _ := strconv.Atoi(id)
			balance, err := bank.ViewBalance(idx)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Balance for Account %d: %.2f\n", id, balance)
			}
		case "3":
			fmt.Println("Enter account ID")
			var id string
			fmt.Scanln(&id)
			idx, _ := strconv.Atoi(id)
			fmt.Println("Enter amount to deposit")
			var amt int
			fmt.Scan(&amt)
			amtStr := ("Enter amount to deposit: ")
			amtx, _ := strconv.ParseFloat(amtStr, 64)
			if err := bank.Deposit(idx, amtx); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}
		case "4":
			fmt.Println("enter sender account ID:")
			var id string
			fmt.Scanln(&id)
			fromID, _ := strconv.Atoi(id)
			fmt.Println("Enter receiver account ID:")
			var rid string
			fmt.Scan(&rid)
			toID, _ := strconv.Atoi(rid)
			fmt.Println("Enter amount to transfer:")
			var amt int
			fmt.Scanln(&amt)
			if err := bank.Transfer(fromID, toID, float64(amt)); err != nil {
				fmt.Println("Error:", err)
			}
		case "5":
			fmt.Println("Exiting Bank App. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
