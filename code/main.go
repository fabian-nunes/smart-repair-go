package main

import (
	"fmt"
	"log"
	"os"
	auth "smart-repair-go/auth"
	dash "smart-repair-go/dash"
)

func createDir() {
	_, err := os.Stat("files")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("files", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
}

func main() {
	createDir()
	var active bool = true

	for active {
		logedIn := auth.AuthenticationMain()

		if !logedIn {
			active = false
		}

		for logedIn {
			option := dash.DashMain()
			switch option {
			case 1:
				fmt.Println("You have selected to view the clients menu.")
			case 2:
				fmt.Println("You have selected to view the repairs menu.")
			case 4:
				fmt.Println("Thank you for using Smart Repair!")
				logedIn = false
			}
		}
	}
	os.Exit(0)
}
