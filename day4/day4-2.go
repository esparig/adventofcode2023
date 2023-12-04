package day4

func getCopyCards(card card) []int {

	var copycards []int
	wins := 1
	j := 0
	for i := 0; i < len(card.winnings); i++ {
		for j < len(card.numbers) {
			if card.winnings[i] == card.numbers[j] {
				copycards = append(copycards, card.id+wins)
				wins++
			}
			if card.numbers[j] > card.winnings[i] {
				break
			}
			j++
		}
	}
	return copycards
}

// recursive function that counts the number of cards
func fun(cards []card, card card) int {
	copycards := getCopyCards(card)
	if len(copycards) == 0 { // base case
		return 1
	}
	count := 1
	for _, copycard := range copycards {
		count += fun(cards, cards[copycard-1])
	}
	return count
}

// function to call recursive function
func funAux(cards []card) int {
	sum := 0
	for _, card := range cards {
		sum += fun(cards, cards[card.id-1])
	}
	return sum
}
