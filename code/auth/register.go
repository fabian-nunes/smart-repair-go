package auth

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
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

func register() {
	reader := bufio.NewReader(os.Stdin)

	//clear buffer
	reader.ReadString('\n')

	fmt.Println("To register, please enter a username and a password.")
	name, _ := getInput("Username: ", reader)
	//fmt.Println("Username: " + name)
	password, _ := getInput("Password: ", reader)
	passwordH := hashPassword(password)
	//fmt.Println("Password: " + passwordH)

	user := newUser(name, passwordH)
	savedU := user.Username + ":" + user.Password

	data := []byte(savedU)

	err := os.WriteFile("files/user.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("You have successfully registered!")
}
