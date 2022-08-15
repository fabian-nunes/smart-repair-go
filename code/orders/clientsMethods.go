package orders

import (
	"bufio"
	"fmt"
	"os"
)

func printAllClients() {
	fmt.Println("Listing all orders...")
	fmt.Println("----------------------------------------------------")
	for _, client := range clients {
		fmt.Println("Name: ", client.name, "\nEmail: ", client.email, "\nPhone: ", client.phone, "\nRepairs: ", client.repairs)
		fmt.Println("----------------------------------------------------")
	}
}

func printSingleClient() {
	fmt.Println("finding a client...")
	fmt.Println("----------------------------------------------------")
	fmt.Println("Enter the email of the client you want to find: ")
	var email string
	_, err := fmt.Scan(&email)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, client := range clients {
			if client.email == email {
				fmt.Println("----------------------------------------------------")
				fmt.Println("Name: ", client.name, "\nEmail: ", client.email, "\nPhone: ", client.phone, "\nRepairs: ", client.repairs)
				fmt.Println("----------------------------------------------------")
			}
		}
	}
}

func addClient() {
	fmt.Println("----------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	//clear buffer
	reader.ReadString('\n')
	fmt.Println("To add a client, please enter the following information: ")
	name, _ := getInput("Name of the client: ", reader)
	email, _ := getInput("Email of the client: ", reader)

	for !validE(email) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Email is not valid. Please try again.")
		email, _ = getInput("Email of the client: ", reader)
	}

	for findClient(email, 0) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Client with this email already exists. Please enter a different email: ")
		email, _ = getInput("Email of the client: ", reader)
	}

	phone, _ := getInput("Phone of the client: ", reader)
	for !validPhone(phone) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Phone is not valid. Please try again.")
		phone, _ = getInput("Phone of the client: ", reader)
	}

	for findClient(phone, 1) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Client with this phone already exists. Please enter a different phone: ")
		phone, _ = getInput("Phone of the client: ", reader)
	}

	clients = append(clients, newClient(name, email, phone))
	fmt.Println("----------------------------------------------------")
	fmt.Println("Client added successfully!")
	fmt.Println("----------------------------------------------------")
}

func findClient(value string, search int) bool {
	for _, client := range clients {
		if search == 0 {
			if client.email == value {
				return true
			}
		} else {
			if client.phone == value {
				return true
			}
		}
	}
	return false
}

func findIndex(email string) int {
	for i, client := range clients {
		if client.email == email {
			return i
		}
	}
	return -1
}

func editClient() bool {

	if len(clients) == 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("There are no clients to edit. Please add a client first.")
		return false
	}

	fmt.Println("----------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	//clear buffer
	reader.ReadString('\n')
	fmt.Println("To edit a client, please enter the following information: ")
	email, _ := getInput("Email of the client: ", reader)
	if !findClient(email, 0) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Client with this email does not exist. Please try again.")
		return false
	}

	index := findIndex(email)

	var choice int

	for choice != 4 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("What do you want to edit?")
		fmt.Println("1. Name")
		fmt.Println("2. Email")
		fmt.Println("3. Phone")
		fmt.Println("4. Back")
		fmt.Println("----------------------------------------------------")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
		} else {
			switch choice {
			case 1:
				reader.ReadString('\n')
				fmt.Println("----------------------------------------------------")
				fmt.Println("Enter the new name of the client: ")
				name, _ := getInput("Name of the client: ", reader)
				clients[index].name = name

				fmt.Println("----------------------------------------------------")
				fmt.Println("Client name successfully updated!")
				fmt.Println("----------------------------------------------------")
			case 2:
				reader.ReadString('\n')
				fmt.Println("----------------------------------------------------")
				fmt.Println("Enter the new email of the client: ")
				emailN, _ := getInput("Email of the client: ", reader)

				for !validE(emailN) {
					fmt.Println("----------------------------------------------------")
					fmt.Println("Email is not valid. Please try again.")
					email, _ = getInput("Email of the client: ", reader)
				}

				for findClient(emailN, 0) {
					fmt.Println("----------------------------------------------------")
					fmt.Println("Client with this email already exists. Please enter a different email: ")
					emailN, _ = getInput("Email of the client: ", reader)
				}
				clients[index].email = emailN

				fmt.Println("----------------------------------------------------")
				fmt.Println("Client email successfully updated!")
				fmt.Println("----------------------------------------------------")

			case 3:
				reader.ReadString('\n')
				fmt.Println("----------------------------------------------------")
				fmt.Println("Enter the new phone of the client: ")
				phoneN, _ := getInput("Phone of the client: ", reader)

				for !validPhone(phoneN) {
					fmt.Println("----------------------------------------------------")
					fmt.Println("Phone is not valid. Please try again.")
					phoneN, _ = getInput("Phone of the client: ", reader)
				}

				for findClient(phoneN, 1) {
					fmt.Println("----------------------------------------------------")
					fmt.Println("Client with this phone already exists. Please enter a different phone: ")
					phoneN, _ = getInput("Phone of the client: ", reader)
				}
				clients[index].phone = phoneN

				fmt.Println("----------------------------------------------------")
				fmt.Println("Client phone successfully updated!")
				fmt.Println("----------------------------------------------------")

			case 4:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Returning to Client menu...")
				fmt.Println("----------------------------------------------------")

			default:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Invalid choice. Please try again.")
				fmt.Println("----------------------------------------------------")
			}
		}
	}

	return true
}

func deleteClient() bool {
	fmt.Println("----------------------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	//clear buffer
	reader.ReadString('\n')
	fmt.Println("To delete a client, please enter the following information: ")
	email, _ := getInput("Email of the client: ", reader)
	if !findClient(email, 0) {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Client with this email does not exist. Please try again.")
		return false
	}

	index := findIndex(email)
	clients = append(clients[:index], clients[index+1:]...)
	fmt.Println("----------------------------------------------------")
	fmt.Println("Client deleted successfully!")
	fmt.Println("----------------------------------------------------")

	removeRepairClient(email)
	return true
}

func changeRepairClient(email string, action int) {
	index := findIndex(email)
	if action == 0 {
		clients[index].repairs += 1
	} else {
		clients[index].repairs -= 1
	}
}
