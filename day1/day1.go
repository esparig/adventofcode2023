package day1

import "fmt"

func Day1() {
	//getCalibrationNumberFromFile()
	//getCalibrationNumberFromFile2()
	//testgetDigitIndexFromString()
	//TestGetCalibrationNumber()
	fmt.Println(GetCalibrationNumberFromFile("input.txt"))
}

var s1 string = "abc123def"
var s2 string = "two1"
var s3 string = "two1nine"
var s4 string = "1nine5"

func TestGetCalibrationNumber() {
	fmt.Println(s1)
	fmt.Println(GetCalibrationNumber(s1))
	fmt.Println(s2)
	fmt.Println(GetCalibrationNumber(s2))
	fmt.Println(s3)
	fmt.Println(GetCalibrationNumber(s3))
	fmt.Println(s4)
	fmt.Println(GetCalibrationNumber(s4))
}

func testgetDigitIndexFromString() {
	// Test getDigitIndexFromString

	fmt.Println(s1)
	fmt.Println(getDigitIndexFromString(s1, numbers))
	fmt.Println(s2)
	fmt.Println(getDigitIndexFromString(s2, numbers))
	fmt.Println(s3)
	fmt.Println(getDigitIndexFromString(s3, numbers))
	fmt.Println(s4)
	fmt.Println(getDigitIndexFromString(s4, numbers))

	var digit int
	var index int

	fmt.Println(s1)
	digit, index = getDigitIndexFromString(reverseString(s1), srebmun)
	if digit == -1 {
		fmt.Println("No digit found")
	} else {
		newindex := len(s1) - len(srebmun[digit-1]) - index
		fmt.Println(digit, newindex)
	}

	fmt.Println(s2)
	digit, index = getDigitIndexFromString(reverseString(s2), srebmun)
	if digit == -1 {
		fmt.Println("No digit found")
	} else {
		newindex := len(s2) - len(srebmun[digit-1]) - index
		fmt.Println(digit, newindex)
	}

	fmt.Println(s3)
	digit, index = getDigitIndexFromString(reverseString(s3), srebmun)
	if digit == -1 {
		fmt.Println("No digit found")
	} else {
		newindex := len(s3) - len(srebmun[digit-1]) - index
		fmt.Println(digit, newindex)
	}

	fmt.Println(s4)
	digit, index = getDigitIndexFromString(reverseString(s4), srebmun)
	if digit == -1 {
		fmt.Println("No digit found")
	} else {
		newindex := len(s4) - len(srebmun[digit-1]) - index
		fmt.Println(digit, newindex)
	}
}
