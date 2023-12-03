package day3

// matrix of numbers to store the number of adjacents numbrs to a * symbol
// but instead of ints its of a struct called gear
type gear struct {
	//adjacents
	adjs int
	// map of numbers (key: number, value: true)
	nums map[int]bool
}

// a number is a part number if it's adjacent to a symbol
func getGears(lines []string, n number, gears [][]gear) {
	row := n.initialdigit[0]
	inicol := n.initialdigit[1]
	endcol := n.finaldigit[1]

	// check if the number is adjacent to a symbol *
	if inicol > 0 {
		if rune(lines[row][inicol-1]) == '*' { // left
			gears[row][inicol-1].adjs++
			gears[row][inicol-1].nums[n.number] = true
		}
		if row > 0 {
			if rune(lines[row-1][inicol-1]) == '*' { // up-left
				gears[row-1][inicol-1].adjs++
				gears[row-1][inicol-1].nums[n.number] = true
			}
		}
		if row < len(lines)-1 {
			if rune(lines[row+1][inicol-1]) == '*' { // down-left
				gears[row+1][inicol-1].adjs++
				gears[row+1][inicol-1].nums[n.number] = true
			}
		}
	}
	for i := inicol; i <= endcol; i++ {
		if row > 0 {
			if rune(lines[row-1][i]) == '*' { // up
				gears[row-1][i].adjs++
				gears[row-1][i].nums[n.number] = true
			}
		}
		if row < len(lines)-1 {
			if rune(lines[row+1][i]) == '*' { // down
				gears[row+1][i].adjs++
				gears[row+1][i].nums[n.number] = true
			}
		}
	}
	if endcol < len(lines[row])-1 {
		if rune(lines[row][endcol+1]) == '*' { // right
			gears[row][endcol+1].adjs++
			gears[row][endcol+1].nums[n.number] = true
		}
		if row > 0 {
			if rune(lines[row-1][endcol+1]) == '*' { // up-right
				gears[row-1][endcol+1].adjs++
				gears[row-1][endcol+1].nums[n.number] = true
			}
		}
		if row < len(lines)-1 {
			if rune(lines[row+1][endcol+1]) == '*' { // down-right
				gears[row+1][endcol+1].adjs++
				gears[row+1][endcol+1].nums[n.number] = true
			}
		}
	}
}
