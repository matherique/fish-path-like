package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fpa, err := filepath.Abs(".")

	if err != nil {
		log.Fatalf("could not get absolute file path: %v", err)
	}

	fromhome := strings.Replace(fpa, os.Getenv("HOME"), "", 1)
	ss := strings.Split(fromhome, string(filepath.Separator))

	var fc []string

	for _, p := range ss {
		if p != "" {
			fc = append(fc, p[0:1])
		}

	}

	result := strings.Join(fc, string(filepath.Separator))
	if result != "" {
		result = "/" + result
	}

	fmt.Fprintf(os.Stdout, fmt.Sprintf("~%s", result))
}
