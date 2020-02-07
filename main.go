package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// get file absolute path
	fpa, err := filepath.Abs(".")

	if err != nil {
		log.Fatalf("could not get absolute file path: %v", err)
	}
	// remove home user dir path from fap
	fromhome := strings.Replace(fpa, os.Getenv("HOME"), "", 1)
	// divide by /
	ss := strings.Split(fromhome, string(filepath.Separator))

	var fc []string
	// get only first char and add to slice

	last := len(ss) - 1

	for _, p := range ss[0:last] {
		if p != "" {
			fc = append(fc, p[0:1])
		}
	}

	fc = append(fc, ss[last])

	// join the result by /
	result := strings.Join(fc, string(filepath.Separator))

	if result != "" {
		result = "/" + result
	}
	// put int stdout the result
	fmt.Fprintf(os.Stdout, fmt.Sprintf("~%s", result))
}
