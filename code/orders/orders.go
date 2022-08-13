package orders

import (
	"bufio"
	"fmt"
	"strings"
)

var clients []Client
var repairs []Repair

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	text, err := r.ReadString('\n')
	return strings.TrimSpace(text), err
}

func ClientsMain() {
	var choice int

	fmt.Println("----------------------------------------------------")
	fmt.Println("Welcome to the client management system.")
	fmt.Println("----------------------------------------------------")

	findFile()

	for choice != 8 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Please select an option below: \n1. List all orders \n2. Search for a client \n3. Add a Client \n4. Edit a client \n5. Delete a client " +
			"\n6. Save orders to file \n7. Help \n8. Exit")
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
				fmt.Println("You have selected to view the help menu.")
			case 8:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Returning to the dashboard")
				fmt.Println("----------------------------------------------------")
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}

	}
}

func RepairsMain() {
	var choice int

	fmt.Println("----------------------------------------------------")
	fmt.Println("Welcome to the repair management system.")
	fmt.Println("----------------------------------------------------")

	findFile()
	findFileR()

	for choice != 9 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Please select an option below: \n1. List all repairs \n2. Search for a repair \n3. Add a repair \n4. Edit a repair \n5. Change Status of Repair \n6. Delete a repair " +
			"\n7. Save repairs to file \n8. Help \n9. Exit")
		fmt.Println("Please enter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
		} else {
			switch choice {
			case 1:
				printRepairs()
			case 2:
				printRepairSingle()
			case 3:
				addRepair()
			case 4:
				editRepair()
			case 5:
				changeStatus()
			case 6:
				deleteRepair()
			case 7:
				saveFileR()
			case 8:
				fmt.Println("You have selected to view the help menu.")
			case 9:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Returning to the dashboard")
				fmt.Println("----------------------------------------------------")
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}

	}
}
