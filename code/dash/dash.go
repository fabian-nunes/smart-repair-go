package dash

import "fmt"

func DashMain() int {

	var choice int

	fmt.Println("----------------------------------------------------")
	fmt.Println("Welcome to Smart Repair Dashboard!")

	for choice != 4 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Please select an option below: \n1. Clients \n2. Repairs" +
			"\n3. Help" + "\n4. Logout")
		fmt.Println("Please enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
		} else {
			if choice == 1 || choice == 2 || choice == 4 {
				return choice
			} else if choice == 3 {
				fmt.Println("----------------------------------------------------")
				fmt.Println("This is the Smart Repair Dashboard. It is a" + " \n" + "web application that allows you to manage your" + " \n" + "clients and their repairs. You can add, edit, and" + " \n" + "delete clients and repairs. You can also view the" + " \n" + "status of your repairs. You can also logout at any" + " \n" + "time.")
				fmt.Println("----------------------------------------------------")
			} else {
				fmt.Println("Invalid choice. Please try again.")
			}
		}
	}
	return 0
}
