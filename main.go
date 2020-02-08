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

	result := ConvertPath(dpa, os.Getenv("HOME"))

	fmt.Fprintf(os.Stdout, result)
}
