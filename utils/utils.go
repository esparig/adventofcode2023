package utils

import (
	"io/ioutil"
	"strings"
)

// function to read lines from a file and return a slice of strings
func ReadLines(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
