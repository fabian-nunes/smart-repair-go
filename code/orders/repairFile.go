package orders

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findFileR() {
	f, err := os.ReadFile("files/orders.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Order file does not exist. Register orders and save them to the file.")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Order file found.")
		fmt.Println("----------------------------------------------------")
		fmt.Println("Loading orders...")
		fmt.Println("----------------------------------------------------")
		repairs = loadRepairs(f)
	}
}

func loadRepairs(f []byte) []Repair {
	var repairs []Repair
	for _, line := range strings.Split(string(f), "\n") {
		if line != "" {
			data := strings.Split(line, ";")

			code, _ := strconv.Atoi(data[0])
			cost, _ := strconv.ParseFloat(data[3], 64)
			status, _ := strconv.Atoi(data[6])
			repairs = append(repairs, Repair{code, data[1], data[2], cost, data[4], data[5], status})

		}
	}
	return repairs
}

func saveFileR() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("Saving file...")
	fmt.Println("----------------------------------------------------")
	file, err := os.Create("files/orders.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, repair := range repairs {
		file.WriteString(strconv.Itoa(repair.code) + ";" + repair.name + ";" + repair.date + ";" + strconv.FormatFloat(repair.cost, 'f', 2, 64) + ";" + repair.description + ";" + repair.client + ";" + strconv.Itoa(repair.status) + "\n")
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("File saved successfully!")
	fmt.Println("----------------------------------------------------")
}
