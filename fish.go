package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// get directory absolute path
	dpa, err := filepath.Abs(".")

	if err != nil {
		log.Fatalf("could not get absolute file path: %v", err)
	}

	home, env := os.UserHomeDir()

	if env != nil {
		log.Fatalf("could not get user home dir: %v", err)
	}

	result := ConvertPath(dpa, home)

	fmt.Fprintf(os.Stdout, result)
}
