package auth

import (
	"bufio"
	"fmt"
	"os"
)

func register() bool {
	fmt.Println("----------------------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	//clear buffer
	reader.ReadString('\n')

	fmt.Println("To register, please enter a username and a password.")
	name, _ := getInput("Username: ", reader)
	//fmt.Println("Username: " + name)
	password, _ := getInput("Password: ", reader)
	password = hashPassword(password)
	//fmt.Println("Password: " + passwordH)

	user := newUser(name, password)
	savedU := user.Username + ":" + user.Password

	data := []byte(savedU)

	err := os.WriteFile("files/user.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("You have successfully registered!")
	fmt.Println("----------------------------------------------------")
	return true
}
