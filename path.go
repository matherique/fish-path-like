package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ConvertPath return a new path
func ConvertPath(path, userdir string) string {
	var pref string

	pathFromHome := strings.Replace(path, userdir, "", 1)
	if pathFromHome != path {
		pref = "~"
	}
	// check if pathfromhome is igual "/"
	if pathFromHome == string(filepath.Separator) {
		pref = "/"
	}

	// divide by /
	ss := strings.Split(pathFromHome, string(filepath.Separator))
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
	return fmt.Sprintf("%s%s", pref, result)
}
