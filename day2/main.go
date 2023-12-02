package main

import "fmt"

func main() {
	if IsAPossibleGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green") {
		fmt.Println("Game 1 is possible")
	} else {
		fmt.Println("Game 1 is not possible")
	}
	fmt.Println("GetSumIDsPossibleGames for input-test1.txt:")
	fmt.Println(GetSumIDsPossibleGames("input-test1.txt"))
	fmt.Println("GetSumIDsPossibleGames for input.txt:")
	fmt.Println(GetSumIDsPossibleGames("input.txt"))
}
