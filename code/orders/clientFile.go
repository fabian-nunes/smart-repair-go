package orders

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findFile() {
	fmt.Println("Search for a client file...")
	f, err := os.ReadFile("files/clients.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Client file does not exist. Register orders and save them to the file.")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if getSizeFile() == 0 {
			fmt.Println("Client file is empty. Register orders and save them to the file.")
		} else {
			fmt.Println("----------------------------------------------------")
			fmt.Println("Client file found.")
			fmt.Println("----------------------------------------------------")
			fmt.Println("Loading clients...")
			fmt.Println("----------------------------------------------------")
			clients = loadClients(f)
		}
	}
}

func getSizeFile() int64 {
	file, err := os.Open("files/clients.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return fileInfo.Size()
}

func loadClients(f []byte) []Client {
	var clients []Client
	for _, line := range strings.Split(string(f), "\n") {
		if line != "" {
			data := strings.Split(line, ";")
			repairs, _ := strconv.Atoi(data[3])
			clients = append(clients, Client{data[0], data[1], data[2], repairs})
		}
	}
	return clients
}

func saveFile() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("Saving Client file...")
	fmt.Println("----------------------------------------------------")
	file, err := os.Create("files/clients.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, client := range clients {
		file.WriteString(client.name + ";" + client.email + ";" + client.phone + ";" + strconv.Itoa(client.repairs) + "\n")
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("File saved successfully!")
	fmt.Println("----------------------------------------------------")
}
