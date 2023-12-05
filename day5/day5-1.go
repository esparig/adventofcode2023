package day5

import (
	"fmt"
	"strconv"
	"strings"
)

// create an enum so I can use a word for an index
const (
	seed_to_soil            = iota //  == 0
	soil_to_fertilizer      = iota //  == 1
	fertilizer_to_water     = iota //  == 2
	water_to_light          = iota //  == 3
	light_to_temperature    = iota //  == 4
	temperature_to_humidity = iota //  == 5
	humidity_to_location    = iota //  == 6
)

// create struct that contains src, dest and width (all three are integers)
type fx struct {
	src   int
	dest  int
	width int
}

func getSeeds(lines []string) []int {
	seeds := make([]int, 0)
	//for each s
	for _, s := range strings.Fields(strings.Split(lines[0], ":")[1]) {
		//convert s to int and store in a slice
		num, _ := strconv.Atoi(s)
		seeds = append(seeds, num)
	}
	return seeds
}

func getMappings(lines []string) [][]fx {
	// create an array of 7 arrays. Those 7 arrays will contain fx structs
	mappings := make([][]fx, 7)
	currentmapping := -1
	//go trough lines from line 3
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		// if line contains string " map:"
		if strings.Contains(line, " map:") {
			currentmapping++
		} else {
			// if line is empty ignore it
			for line == "" {
				break
			}
			// if line containts three numbers
			nums := strings.Fields(line)
			if len(nums) == 3 {
				// convert numbers to int
				dest, _ := strconv.Atoi(nums[0])
				src, _ := strconv.Atoi(nums[1])
				width, _ := strconv.Atoi(nums[2])
				// create a new fx struct
				f := fx{src, dest, width}
				// add fx struct to array of fx structs
				mappings[currentmapping] = append(mappings[currentmapping], f)
			}
		}
	}
	return mappings
}

func processInput(lines []string) ([][]fx, []int) {

	seeds := getSeeds(lines)

	mappings := getMappings(lines)

	return mappings, seeds
}

func getLocation(mappings [][]fx, seed int, verbose bool) int {

	// go though mappings[seed_to_soil] and find the value for the key seed
	soil := -1
	for _, m := range mappings[seed_to_soil] {
		if m.src <= seed && seed < m.src+m.width {
			soil = m.dest + seed - m.src
			break
		}
	}
	if soil == -1 {
		soil = seed
	}
	if verbose {
		fmt.Println("Soil: ", soil)
	}

	// go though mappings[soil_to_fertilizer] and find the value for the key soil
	fertilizer := -1
	for _, m := range mappings[soil_to_fertilizer] {
		if m.src <= soil && soil < m.src+m.width {
			fertilizer = m.dest + soil - m.src
			break
		}
	}
	if fertilizer == -1 {
		fertilizer = soil
	}
	if verbose {
		fmt.Println("Fertilizer: ", fertilizer)
	}

	// go though mappings[fertilizer_to_water] and find the value for the key fertilizer
	water := -1
	for _, m := range mappings[fertilizer_to_water] {
		if m.src <= fertilizer && fertilizer < m.src+m.width {
			water = m.dest + fertilizer - m.src
			break
		}
	}
	if water == -1 {
		water = fertilizer
	}
	if verbose {
		fmt.Println("Water: ", water)
	}

	// go though mappings[water_to_light] and find the value for the key water
	light := -1
	for _, m := range mappings[water_to_light] {
		if m.src <= water && water < m.src+m.width {
			light = m.dest + water - m.src
			break
		}
	}
	if light == -1 {
		light = water
	}
	if verbose {
		fmt.Println("Light: ", light)
	}

	// go though mappings[light_to_temperature] and find the value for the key light
	temperature := -1
	for _, m := range mappings[light_to_temperature] {
		if m.src <= light && light < m.src+m.width {
			temperature = m.dest + light - m.src
			break
		}
	}
	if temperature == -1 {
		temperature = light
	}
	if verbose {
		fmt.Println("Temperature: ", temperature)
	}

	// go though mappings[temperature_to_humidity] and find the value for the key temperature
	humidity := -1
	for _, m := range mappings[temperature_to_humidity] {
		if m.src <= temperature && temperature < m.src+m.width {
			humidity = m.dest + temperature - m.src
			break
		}
	}
	if humidity == -1 {
		humidity = temperature
	}
	if verbose {
		fmt.Println("Humidity: ", humidity)
	}

	// go though mappings[humidity_to_location] and find the value for the key humidity
	location := -1
	for _, m := range mappings[humidity_to_location] {
		if m.src <= humidity && humidity < m.src+m.width {
			location = m.dest + humidity - m.src
			break
		}
	}
	if location == -1 {
		location = humidity
	}
	if verbose {
		fmt.Println("Location: ", location)
	}
	return location
}
