package day4

import (
	"adventofcode2023/utils"
	"fmt"
)

func Day4() {
	lines := utils.ReadLines("day4/input.txt")
	fmt.Println(lines)
	cards := getCards(lines)
	sum := 0
	for _, card := range cards {
		score := getScore(card)
		fmt.Println("Score: ", card.id)
		sum += score
	}
	fmt.Println("Sum: ", sum)
}
