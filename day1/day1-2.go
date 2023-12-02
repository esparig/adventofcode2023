package day1

import (
	"strconv"
	"strings"
	"unicode"
)

// array of strings "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"
var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var srebmun = []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

func getDigitIndexFromString(s string, substrings []string) (int, int) {
	index := len(s)
	var digit int = -1
	for n, subs := range substrings {
		current_index := strings.Index(s, subs)
		if current_index != -1 && current_index < index {
			index = current_index
			digit = n + 1
		}
	}
	return digit, index
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getFirstStringDigit(s string) (int, int) {
	var first int
	var firstindex int
	first, firstindex = getDigitIndexFromString(s, numbers)
	// fmt.Println("(first, first_index)", first, firstindex)
	return first, firstindex
}

func getLastStringDigit(s string) (int, int) {
	var second int
	var secondindex int
	var secondindexaux int
	second, secondindexaux = getDigitIndexFromString(reverseString(s), srebmun)
	if second != -1 {
		secondindex = len(s) - len(srebmun[second-1]) - secondindexaux
	} else {
		secondindex = -1
	}
	// fmt.Println("(second, second_index)", second, secondindex)
	return second, secondindex
}

func getFirstNumberDigit(s string) (int, int) {
	var first int
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			first, _ = strconv.Atoi(string(s[i]))
			return first, i
		}
	}
	return -1, -1
}

func getLastNumberDigit(s string) (int, int) {
	var second int
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			second, _ = strconv.Atoi(string(s[i]))
			return second, i
		}
	}
	return -1, -1
}

func GetCalibrationNumber(s string) int {
	var first int
	var firstaux int
	var firstindex int
	var firstindexaux int
	var second int
	var secondaux int
	var secondindex int
	var secondindexaux int
	first, firstindex = getFirstNumberDigit(s)
	firstaux, firstindexaux = getFirstStringDigit(s)
	if first == -1 {
		first = firstaux
	} else {
		if firstindexaux != -1 && firstindexaux < firstindex {
			first = firstaux
		}
	}
	second, secondindex = getLastNumberDigit(s)
	secondaux, secondindexaux = getLastStringDigit(s)
	if second == -1 {
		second = secondaux
	} else {
		if secondindexaux != -1 && secondindexaux > secondindex {
			second = secondaux
		}
	}
	if first == -1 || second == -1 {
		return 0
	}
	return first*10 + second
}

func GetCalibrationNumberFromFile(filename string) int {
	// read input file
	lines := readLines(filename)
	// initialize calibration number
	calibration := 0
	// for each line in the input file, extract the calibration number
	for _, line := range lines {
		calNum := GetCalibrationNumber(line)
		calibration += calNum
	}
	// print the calibration number
	return calibration
}
