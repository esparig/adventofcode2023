package day5

import (
	"strconv"
	"strings"
)

// create struct for seeds, two integers: seed and number of seeds
type seed struct {
	seed          int
	numberofseeds int
}

func getSeeds2(lines []string) []seed {
	seeds := make([]seed, 0)
	//for each s
	seedsinfo := strings.Fields(strings.Split(lines[0], ":")[1])
	for i := 0; i < len(seedsinfo); i += 2 {
		currentseed, _ := strconv.Atoi(seedsinfo[i])
		numberofseeds, _ := strconv.Atoi(seedsinfo[i+1])
		s := seed{currentseed, numberofseeds}
		seeds = append(seeds, s)
	}
	return seeds
}

func processInput2(lines []string) ([][]fx, []seed) {

	seeds := getSeeds2(lines)

	mappings := getMappings(lines)

	return mappings, seeds
}
