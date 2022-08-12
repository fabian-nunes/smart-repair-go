package auth

import (
	"fmt"
	"os"
)

func AuthenticationMain() {
	var choice int

	fmt.Println("Hello, welcome to Smart Repair!")

	for choice != 4 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Please select an option below: \n1. Login \n2. Register" +
			"\n3. Help" + "\n4. Exit")
		fmt.Println("Please enter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
		} else {
			switch choice {
			case 1:
				fmt.Println("You have selected to login.")
			case 2:
				register()
			case 3:
				fmt.Println("You have selected to view the help menu.")
			case 4:
				fmt.Println("Thank you for using Smart Repair!")
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}

	}
}
