package day2

import "fmt"

func Day2() {
	if IsAPossibleGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green") {
		fmt.Println("Game 1 is possible")
	} else {
		fmt.Println("Game 1 is not possible")
	}
	fmt.Println("GetSumIDsPossibleGames for input-test1.txt:")
	fmt.Println(GetSumIDsPossibleGames("day2/input-test1.txt"))
	fmt.Println("GetSumIDsPossibleGames for input.txt:")
	fmt.Println(GetSumIDsPossibleGames("day2/input.txt"))

	fmt.Println("GetSumPowers for input-test1.txt:")
	fmt.Println(GetSumPowers("day2/input-test1.txt"))
	fmt.Println("GetSumPowers for input.txt:")
	fmt.Println(GetSumPowers("day2/input.txt"))
}
