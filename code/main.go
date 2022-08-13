package main

import (
	"fmt"
	"log"
	"os"
	auth "smart-repair-go/auth"
	dash "smart-repair-go/dash"
	orders "smart-repair-go/orders"
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
				orders.ClientsMain()
			case 2:
				orders.RepairsMain()
			case 4:
				fmt.Println("Thank you for using Smart Repair!")
				logedIn = false
			}
		}
	}
	os.Exit(0)
}
