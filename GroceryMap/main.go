package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	groceryList := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nGrocery List Menu:")
		fmt.Println("1. Add/Update Item")
		fmt.Println("2. Remove Item")
		fmt.Println("3. View List")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter item name: ")
			scanner.Scan()
			item := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter quantity: ")
			scanner.Scan()
			qtyStr := strings.TrimSpace(scanner.Text())
			qty, err := strconv.Atoi(qtyStr)
			if err != nil || qty <= 0 {
				fmt.Println("Invalid quantity. Please enter a positive number.")
				continue
			}

			groceryList[item] = qty
			fmt.Printf("%s added/updated with quantity %d\n", item, qty)

		case "2":
			fmt.Print("Enter item name to remove: ")
			scanner.Scan()
			item := strings.TrimSpace(scanner.Text())

			if _, exists := groceryList[item]; exists {
				delete(groceryList, item)
				fmt.Printf("%s removed from the list.\n", item)
			} else {
				fmt.Println("Item not found in the list.")
			}

		case "3":
			if len(groceryList) == 0 {
				fmt.Println("Your grocery list is empty.")
			} else {
				fmt.Println("Your Grocery List:")
				for item, qty := range groceryList {
					fmt.Printf("- %s: %d\n", item, qty)
				}
			}

		case "4":
			fmt.Println("Exiting Grocery List App. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
