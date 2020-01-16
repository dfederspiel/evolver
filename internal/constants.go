package evolver

const MAX_X = 5
const MAX_Y = 5
const MANGA_CAPACITY = 8.0
const MANGA_GROW = .25

var grid[MAX_X * MAX_Y]cell

func findCellIndex(x, y int) int {
	for i, element := range grid {
		if element.Y == y && element.X == x {
			return i
		}
	}
	return 0
}