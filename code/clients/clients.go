package clients

import (
	"fmt"
)

var clients []Client

func ClientsMain() {
	var choice int

	fmt.Println("----------------------------------------------------")
	fmt.Println("Welcome to the client management system.")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Search for a client file...")

	findFile()

	for choice != 7 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Please select an option below: \n1. List all clients \n2. Search for a client \n3. Add a Client \n4. Edit a client \n5. Delete a client " +
			"\n6. Save clients to file \n7. Exit")
		fmt.Println("Please enter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
		} else {
			switch choice {
			case 1:
				printAllClients()
			case 2:
				printSingleClient()
			case 3:
				addClient()
			case 4:
				editClient()
			case 5:
				deleteClient()
			case 6:
				saveFile()
			case 7:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Returning to the dashboard")
				fmt.Println("----------------------------------------------------")
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}

	}
}
