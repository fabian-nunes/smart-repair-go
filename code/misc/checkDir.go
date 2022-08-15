package misc

import (
	"log"
	"os"
)

func CreateDir() {
	_, err := os.Stat("files")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("files", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
}
