package day3

import (
	"adventofcode2023/utils"
	"fmt"
)

func Day3() {

	fmt.Println("Day 3: input test")
	lines := utils.ReadLines("day3/input.txt")
	numbers := getNumbers(lines, false)
	sum := 0
	for _, num := range numbers {
		fmt.Println(num.number, "From: (", num.initialdigit[0], ",", num.initialdigit[1], ") To: (",
			num.finaldigit[0], ",", num.finaldigit[1], ") - ", isPartNumber(lines, num))
		if isPartNumber(lines, num) {
			sum += num.number
		}
	}
	fmt.Println("Sum of part numbers: ", sum)

	// fmt.Println("\nDay 3: input")
	// lines = utils.ReadLines("day3/input.txt")
	// numbers = getNumbers(lines, false)
	// sum = 0

	// for _, num := range numbers {
	// 	fmt.Print(num.number, " From: (", num.initialdigit[0], ",", num.initialdigit[1], ") To: (",
	// 		num.finaldigit[0], ",", num.finaldigit[1], ")")
	// 	if isPartNumber(lines, num) {
	// 		fmt.Println(" - true")
	// 		sum += num.number
	// 	} else {
	// 		fmt.Println(" - false")
	// 	}
	// }

	fmt.Println("Sum of part numbers: ", sum)

	//initialize Gears
	gears := make([][]gear, len(lines))
	for i := range gears {
		gears[i] = make([]gear, len(lines[i]))
		for j := range gears[i] {
			gears[i][j].nums = make(map[int]bool)
		}
	}
	//print Gears
	for i := 0; i < len(gears); i++ {
		for j := 0; j < len(gears[i]); j++ {
			fmt.Print(gears[i][j].adjs, " ")
		}
		fmt.Println()
	}
	fmt.Println("\nDay 3: part 2")
	for _, num := range numbers {
		getGears(lines, num, gears)
	}
	//print Gears
	var gearRatios []int
	for i := 0; i < len(gears); i++ {
		for j := 0; j < len(gears[i]); j++ {
			fmt.Print(gears[i][j].adjs, " ")
			if gears[i][j].adjs == 2 {
				// go thourh the map and add the numbers to the slice
				gearRatio := 1
				for k := range gears[i][j].nums {
					gearRatio *= k
				}
				gearRatios = append(gearRatios, gearRatio)
			}
		}
		fmt.Println()
	}
	// print adjacents
	fmt.Println("gearRatios: ", gearRatios)

	//sum all gearRatios
	sum = 0
	for _, gearRatio := range gearRatios {
		sum += gearRatio
	}
	fmt.Println("Sum of gearRatios: ", sum)
}
