package day5

import (
	"adventofcode2023/utils"
	"fmt"
	"math"
)

func Day5() {
	lines := utils.ReadLines("day5/input.txt")

	// mappings, seeds := processInput(lines)
	// fmt.Println("Seed to soil:\n", mappings[seed_to_soil])
	// fmt.Println("Soil to fertilizer:\n", mappings[soil_to_fertilizer])
	// fmt.Println("Fertilizer to water:\n", mappings[fertilizer_to_water])
	// fmt.Println("Water to light:\n", mappings[water_to_light])
	// fmt.Println("Light to temperature:\n", mappings[light_to_temperature])
	// fmt.Println("Temperature to humidity:\n", mappings[temperature_to_humidity])
	// fmt.Println("Humidity to location:\n", mappings[humidity_to_location])

	// // initialize minloc to MAX INT
	// minloc := 2147483647
	// for _, seed := range seeds {
	// 	loc := getLocation(mappings, seed, false)
	// 	fmt.Println("Seed: ", seed, " Location: ", loc)
	// 	if loc < minloc {
	// 		minloc = loc
	// 	}
	// }
	// println("Nearest location: ", minloc)

	mappings, seeds := processInput2(lines)
	fmt.Println("Seeds: ", seeds)

	//map to know if m has been visited
	visited := make(map[int]bool)

	// mappings, _ := processInput2(lines)
	// fmt.Println(getLocation(mappings, 111819874, false))

	// initialize minloc to MAX INT
	minloc := math.MaxInt64
	for _, seed := range seeds {
		for i := 0; i < seed.numberofseeds; i++ {
			if !visited[seed.seed+i] {
				visited[seed.seed+i] = true
				loc := getLocation(mappings, seed.seed+i, false)
				//fmt.Println("Seed: ", seed.seed+i, " Location: ", loc)
				if loc < minloc {
					minloc = loc
				}
			}
		}
	}
	println("Nearest location: ", minloc)

}
