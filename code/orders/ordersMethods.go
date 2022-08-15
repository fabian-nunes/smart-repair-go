package orders

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func printRepairs() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("List of repairs:")
	fmt.Println("----------------------------------------------------")
	for _, repair := range repairs {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Code:", repair.code)
		fmt.Println("Name:", repair.name)
		fmt.Println("Date:", repair.date)
		fmt.Println("Cost:", repair.cost, "€")
		fmt.Println("Description:", repair.description)
		fmt.Println("Client:", repair.client)
		switch repair.status {
		case 0:
			fmt.Println("Status:", "Pending")
		case 1:
			fmt.Println("Status:", "In progress")
		case 2:
			fmt.Println("Status:", "Completed/Waiting Payment")
		case 3:
			fmt.Println("Status:", "Completed/Paid")
		}
		fmt.Println("----------------------------------------------------")
	}
}

func printRepairSingle() {
	fmt.Println("Finding repair...")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Enter the code of the repair:")
	var code int
	fmt.Scan(&code)
	for _, repair := range repairs {
		if repair.code == code {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Code:", repair.code)
			fmt.Println("Name:", repair.name)
			fmt.Println("Date:", repair.date)
			fmt.Println("Cost:", repair.cost, "€")
			fmt.Println("Description:", repair.description)
			fmt.Println("Client:", repair.client)
			switch repair.status {
			case 0:
				fmt.Println("Status:", "Pending")
			case 1:
				fmt.Println("Status:", "In progress")
			case 2:
				fmt.Println("Status:", "Completed/Waiting Payment")
			case 3:
				fmt.Println("Status:", "Completed/Paid")
			}
			fmt.Println("----------------------------------------------------")
		}
	}
}

func addRepair() bool {
	if len(clients) == 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("There are no clients. Please add a client first.")
		fmt.Println("----------------------------------------------------")
		return false
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("Adding repair...")
	fmt.Println("----------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	//clear buffer
	reader.ReadString('\n')

	code := rand.Intn(100)

	if findRepair(code) {
		code = rand.Intn(100)
	}

	name, _ := getInput("Name of the repair: ", reader)

	fmt.Println("Enter the cost of the repair:")
	var cost float64
	_, err := fmt.Scan(&cost)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Invalid input.")
		return false
	}

	reader.ReadString('\n')
	description, _ := getInput("Description of the repair: ", reader)

	client, _ := getInput("Client of the repair: ", reader)

	if !findClient(client, 0) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Client does not exist. Please enter a different client: ")
		return false
	}

	repairs = append(repairs, newRepair(code, name, cost, description, client))

	changeRepairClient(client, 0)

	fmt.Println("----------------------------------------------------")
	fmt.Println("Repair added successfully!")
	return true
}

func findIndexRepair(code int) int {
	for i, repair := range repairs {
		if repair.code == code {
			return i
		}
	}
	return -1
}

func editRepair() bool {

	if len(repairs) == 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("There are no repairs. Please add a repair first.")
		fmt.Println("----------------------------------------------------")
		return false
	}

	if len(clients) == 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("There are no clients. Please add a client first.")
		fmt.Println("----------------------------------------------------")
		return false
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("Editing repair...")
	fmt.Println("----------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	//clear buffer
	reader.ReadString('\n')
	fmt.Println("Enter the code of the repair:")
	var code int
	fmt.Scan(&code)

	if !findRepair(code) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Repair does not exist. Please enter a different code: ")
		return false
	}

	index := findIndexRepair(code)

	if repairs[index].status != 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Repair cannot be edited. Please enter a different code: ")
		return false
	}
	var choice int

	for choice != 4 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("What do you want to edit?")
		fmt.Println("1. Name")
		fmt.Println("2. Cost")
		fmt.Println("3. Description")
		fmt.Println("4. Exit")
		fmt.Println("----------------------------------------------------")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a valid choice: ")
		} else {
			reader.ReadString('\n')
			switch choice {
			case 1:
				name, _ := getInput("Name of the repair: ", reader)
				repairs[index].name = name
			case 2:
				fmt.Println("Enter the cost of the repair:")
				var cost float64
				_, err := fmt.Scan(&cost)
				if err != nil {
					fmt.Println("Invalid cost.")
					return false
				}
				repairs[index].cost = cost
			case 3:
				description, _ := getInput("Description of the repair: ", reader)
				repairs[index].description = description
			case 4:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Returning to the repair menu...")
				fmt.Println("----------------------------------------------------")
				return true
			default:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Invalid choice. Please enter a valid choice: ")
			}
		}
	}
	return true
}

func changeStatus() bool {
	if len(repairs) == 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("There are no repairs. Please add a repair first.")
		fmt.Println("----------------------------------------------------")
		return false
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("Changing status of repair...")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Enter the code of the repair:")
	var code int
	fmt.Scan(&code)

	if !findRepair(code) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Repair does not exist. Please enter a different code: ")
		return false
	}

	index := findIndexRepair(code)
	if repairs[index].status == 3 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Repair is already completed. Please enter a different code: ")
		return false
	}
	repairs[index].status += 1
	fmt.Println("----------------------------------------------------")
	fmt.Println("Status changed successfully!")
	fmt.Println("----------------------------------------------------")
	return true
}

func deleteRepair() {
	if len(repairs) == 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("There are no repairs. Please add a repair first.")
		fmt.Println("----------------------------------------------------")
		return
	}
	fmt.Println("----------------------------------------------------")
	fmt.Println("Deleting repair...")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Enter the code of the repair:")
	var code int
	fmt.Scan(&code)

	if !findRepair(code) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Repair does not exist. Please enter a different code: ")
		return
	}

	index := findIndexRepair(code)
	if repairs[index].status != 0 && repairs[index].status != 3 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Repair cannot be deleted. Please enter a different code: ")
		return
	}

	changeRepairClient(repairs[index].client, 1)

	repairs = append(repairs[:index], repairs[index+1:]...)
	fmt.Println("----------------------------------------------------")
	fmt.Println("Repair deleted successfully!")
	fmt.Println("----------------------------------------------------")

}

func removeRepairClient(email string) {
	for _, repair := range repairs {
		if repair.client == email {
			repairs[findIndexRepair(repair.code)].client = "User Deleted"
		}
	}
}
