package main

import (
	"fmt"
	"os"
)

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func ensureExists(path string) {
	if exists(path) == false {
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
