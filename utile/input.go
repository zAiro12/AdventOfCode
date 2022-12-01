package utile

import (
	"fmt"
	"os"
)

func Input(nameFile string) *os.File {
	file, err := os.Open(nameFile)

	if err != nil {
		fmt.Println("File non trovato")
	}

	return file
}
