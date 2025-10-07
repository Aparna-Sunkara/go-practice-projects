package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	var groceryList []string

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n Grocery List App ")
		fmt.Println("1. Add Item")
		fmt.Println("2. View Items")
		fmt.Println("3. Remove Item")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter item name: ")
			item, _ := reader.ReadString('\n')
			item = strings.TrimSpace(item)
			if item == "" {
				fmt.Println("Item name cannot be empty.")
				continue
			}
			groceryList = append(groceryList, item)
			fmt.Printf("'%s' added to grocery list.\n", item)

		case 2:
			if len(groceryList) == 0 {
				fmt.Println(" Your grocery list is empty.")
			} else {
				fmt.Println(" Grocery List:")
				for i, item := range groceryList {
					fmt.Printf("%d. %s\n", i+1, item)
				}
			}

		case 3:
			if len(groceryList) == 0 {
				fmt.Println("No items to remove.")
				continue
			}
			fmt.Println("Enter item number to remove:")
			var index int
			fmt.Scanln(&index)
			if index < 1 || index > len(groceryList) {
				fmt.Println("Invalid item number.")
				continue
			}
			removed := groceryList[index-1]
			groceryList = append(groceryList[:index-1], groceryList[index:]...)
			fmt.Printf("'%s' removed from the list.\n", removed)

		case 4:
			fmt.Println(" Exiting Grocery List App. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number from 1 to 4.")
		}
	}
}
