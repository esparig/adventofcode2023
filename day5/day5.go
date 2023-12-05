package day5

import (
	"adventofcode2023/utils"
	"fmt"
)

func Day5() {
	lines := utils.ReadLines("day5/input.txt")
	// maps, seeds := processInput_(lines)
	// fmt.Println("Seed to soil:\n", maps[seed_to_soil])
	// fmt.Println("Soil to fertilizer:\n", maps[soil_to_fertilizer])
	// fmt.Println("Fertilizer to water:\n", maps[fertilizer_to_water])
	// fmt.Println("Water to light:\n", maps[water_to_light])
	// fmt.Println("Light to temperature:\n", maps[light_to_temperature])
	// fmt.Println("Temperature to humidity:\n", maps[temperature_to_humidity])
	// fmt.Println("Humidity to location:\n", maps[humidity_to_location])
	// // initialize minloc to MAX INT
	// minloc := 2147483647
	// for _, seed := range seeds {
	// 	loc := getLocation_(maps, seed, false)
	// 	fmt.Println("Seed: ", seed, " Location: ", loc)
	// 	if loc < minloc {
	// 		minloc = loc
	// 	}
	// }
	// println("Nearest location: ", minloc)

	mappings, seeds := processInput(lines)
	fmt.Println("Seed to soil:\n", mappings[seed_to_soil])
	fmt.Println("Soil to fertilizer:\n", mappings[soil_to_fertilizer])
	fmt.Println("Fertilizer to water:\n", mappings[fertilizer_to_water])
	fmt.Println("Water to light:\n", mappings[water_to_light])
	fmt.Println("Light to temperature:\n", mappings[light_to_temperature])
	fmt.Println("Temperature to humidity:\n", mappings[temperature_to_humidity])
	fmt.Println("Humidity to location:\n", mappings[humidity_to_location])

	// initialize minloc to MAX INT
	minloc := 2147483647
	for _, seed := range seeds {
		loc := getLocation(mappings, seed, false)
		fmt.Println("Seed: ", seed, " Location: ", loc)
		if loc < minloc {
			minloc = loc
		}
	}
	println("Nearest location: ", minloc)

}
