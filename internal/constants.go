package evolver

import "math/rand"

const MAX_X = 5
const MAX_Y = 5
const TICKS_PER_DAY = 24
const TICKS_PER_MONTH = TICKS_PER_DAY * 30
const TICKS_PER_ANNUM = TICKS_PER_MONTH * 12

const MANGA_THRESHOLD = 7
const MANGA_LIFE = TICKS_PER_DAY * 60
const MANGA_CAPACITY = 8.0
const MANGA_GROW = .25
const MANGA_MATURITY = TICKS_PER_DAY * 10

const CELL_FERTILE = TICKS_PER_DAY * 30

var manga_list_counter = 0
var manga_live_counter = 0
var manga_dead_counter = 0

var Grid[MAX_X * MAX_Y]cell

var Manga[]manga
//var Predo[]predo
//var eggPredo[]eggPredo

func findCellIndex(x, y int) int {
	for i, element := range Grid {
		if element.Y == y && element.X == x {
			return i
		}
	}
	return 0
}

func lifeExpectancy(max int) int {
	return max * ((rand.Intn(max - 1) + 1) / 100)
}