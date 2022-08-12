package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func login() bool {
	fmt.Println("----------------------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	//clear buffer
	reader.ReadString('\n')

	fmt.Println("To login, please enter your username and password.")

	name, _ := getInput("Username: ", reader)

	password, _ := getInput("Password: ", reader)
	password = hashPassword(password)

	user := newUser(name, password)
	//fmt.Println(user)

	f, err := os.ReadFile("files/user.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("User file does not exist. Please register first.")
			return false
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	savedU := string(f)
	savedUData := strings.Split(savedU, ":")

	if user.Username == savedUData[0] && user.Password == savedUData[1] {
		fmt.Println("You have successfully logged in!")
		return true
	} else {
		fmt.Println("Invalid username or password. Please try again.")
		return false
	}
}
