package main

import (
	"fmt"
	"os"
)

func ensureExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func writeStringToFile(str, path string) {

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("unable to create file at [%s]", path)
		os.Exit(1)
	}
	defer f.Close()

	f.WriteString(str)
}
