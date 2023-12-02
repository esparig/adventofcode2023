package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

// function to extract calibration number from string
func extractCalibrationNumber(s string) int {
	// go thwough the string from the beginning until a digit is found
	// convert that digit to an int
	var first int
	var second int
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			first, _ = strconv.Atoi(string(s[i]))
			break
		}
	}
	// go through the string from the end until a digit is found
	// convert that digit to an int
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			second, _ = strconv.Atoi(string(s[i]))
			break
		}
	}
	// return the number formed by those two digits
	return first*10 + second
}

// function to read lines from a file
func readLines(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

// for each line in the input file, extract the calibration number
func getCalibrationNumberFromFile() {
	// read input file
	lines := readLines("input-test.txt")
	// initialize calibration number
	calibration := 0
	// for each line in the input file, extract the calibration number
	for _, line := range lines {
		fmt.Println(line)
		calNum := extractCalibrationNumber(line)
		fmt.Println(calNum)
		calibration += calNum
	}
	// print the calibration number
	fmt.Println("Calibration number", calibration)
}
