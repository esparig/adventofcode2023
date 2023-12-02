package main

import (
	"io/ioutil"
	"strconv"
	strings "strings"
)

var m = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func IsAPossibleGame(s string) bool {
	// s = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	aux := strings.Split(s, ": ")
	if len(aux) != 2 {
		return false
	}
	s = aux[1]
	aux = strings.Split(s, "; ")
	s = strings.Join(aux, ", ")
	aux = strings.Split(s, ", ")

	for i := 0; i < len(aux); i++ {
		numcolor := strings.Split(aux[i], " ")
		num, _ := strconv.Atoi(string(numcolor[0]))
		color := numcolor[1]
		if num > m[color] {
			return false
		}
	}

	return true
}

// function to read lines from a file
func readLines(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func GetSumIDsPossibleGames(filename string) int {
	lines := readLines(filename)
	sum := 0
	for i := 0; i < len(lines); i++ {
		if IsAPossibleGame(lines[i]) {
			sum += i + 1
		}
	}
	return sum
}
