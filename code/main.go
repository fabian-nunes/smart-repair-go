package main

import (
	"fmt"
	"os"
	auth "smart-repair-go/auth"
	dash "smart-repair-go/dash"
	misc "smart-repair-go/misc"
	orders "smart-repair-go/orders"
)

func main() {
	misc.CreateDir()
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
