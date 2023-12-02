package day2

import (
	"strconv"
	strings "strings"
)

func GetSumPowers(filename string) int {
	lines := readLines(filename)
	sum := 0
	for i := 0; i < len(lines); i++ {
		sum += GamePower(lines[i])
	}
	return sum
}

func GamePower(s string) int {
	// s = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	var maxmap = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	aux := strings.Split(s, ": ")
	if len(aux) != 2 {
		return 0
	}
	s = aux[1]
	aux = strings.Split(s, "; ")
	s = strings.Join(aux, ", ")
	aux = strings.Split(s, ", ")

	for i := 0; i < len(aux); i++ {
		numcolor := strings.Split(aux[i], " ")
		num, _ := strconv.Atoi(string(numcolor[0]))
		color := numcolor[1]
		if num > maxmap[color] {
			maxmap[color] = num
		}
	}
	return maxmap["red"] * maxmap["green"] * maxmap["blue"]
}
