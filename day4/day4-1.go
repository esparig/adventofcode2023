package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	id       int
	winnings []int
	numbers  []int
}

func getCards(lines []string) []card {
	// initialize cards array
	cards := make([]card, 0)
	// append a new card for each line
	for i := 0; i < len(lines); i++ {
		id := i + 1
		winnings := make([]int, 0)
		aux := strings.Split(lines[i], ": ")[1]
		line := strings.Split(aux, " | ")
		winningsstr := strings.Fields(line[0])
		// convert each string to int
		for _, str := range winningsstr {
			num, _ := strconv.Atoi(str)
			winnings = append(winnings, num)
		}
		slices.Sort(winnings)
		numbers := make([]int, 0)
		numbersstr := strings.Fields(line[1])
		// convert each string to int
		for _, str := range numbersstr {
			num, _ := strconv.Atoi(str)
			numbers = append(numbers, num)
		}
		slices.Sort(numbers)
		fmt.Println(id, winnings, numbers)
		cards = append(cards, card{id, winnings, numbers})
	}

	return cards
}

func getScore(card card) int {
	score := 0
	j := 0
	for i := 0; i < len(card.winnings); i++ {
		for card.winnings[i] > card.numbers[j] {
			j++
			if j == len(card.numbers) {
				return score
			}
		}
		for card.winnings[i] == card.numbers[j] {
			if score == 0 {
				score = 1
			} else {
				score = score << 1
			}
			j++
			if j == len(card.numbers) {
				return score
			}
		}
	}
	return score
}
