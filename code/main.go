package main

import (
	"log"
	"os"
	auth "smart-repair-go/auth"
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
	auth.AuthenticationMain()
}
