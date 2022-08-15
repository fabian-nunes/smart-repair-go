package auth

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	text, err := r.ReadString('\n')
	return strings.TrimSpace(text), err
}

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func AuthenticationMain() bool {
	var choice int
	var logedIn bool

	fmt.Println("Hello, welcome to Smart Repair!")

	for choice != 4 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Please select an option below: \n1. Login \n2. Register" +
			"\n3. Help" + "\n4. Exit")
		fmt.Println("Please enter your choice: ")
		fmt.Println("----------------------------------------------------")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Invalid choice. Please try again.")
			fmt.Println("----------------------------------------------------")
		} else {
			switch choice {
			case 1:
				logedIn = login()
				if logedIn {
					return true
				}
			case 2:
				logedIn = register()
				if logedIn {
					return true
				}
			case 3:
				fmt.Println("----------------------------------------------------")
				fmt.Println("To enter the system, you must login first, if you do not have a account register first." +
					"To choose different options, use the number next to the option you want to choose.")
				fmt.Println("----------------------------------------------------")
			case 4:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Thank you for using Smart Repair!")
				fmt.Println("----------------------------------------------------")
			default:
				fmt.Println("----------------------------------------------------")
				fmt.Println("Invalid choice. Please try again.")
				fmt.Println("----------------------------------------------------")
			}
		}

	}
	return false
}
