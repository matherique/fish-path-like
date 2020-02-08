package main

import "testing"

func TestConvertPath(t *testing.T) {
	tt := []struct {
		path     string
		expected string
	}{
		{path: "/home/user/apath/bpath", expected: "~/a/bpath"},
		{path: "/home/user/apath/bpath/cpath/dpath/epath/fpath/gpath/hpath/ipath/jpath", expected: "~/a/b/c/d/e/f/g/h/i/jpath"},
		{path: "/home/user", expected: "~"},
		{path: "/home", expected: "/home"},
		{path: "/", expected: "/"},
	}

	for _, tc := range tt {
		newPath := ConvertPath(tc.path, "/home/user")
		if newPath != tc.expected {
			t.Errorf("conveted path %s should be %s; got %s", tc.path, tc.expected, newPath)
		}
	}
}
