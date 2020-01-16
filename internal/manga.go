package evolver

import "fmt"

type manga struct {
	Life int
	Capacity float32
	evolvee
}

func (m manga) age() bool {
	m.Life -= 1
	m.Maturity += 1

	if m.Life <= 1 {
		return true
	}
	return false
}

func (m manga) grow() {
	m.Capacity = m.Capacity * (1 + MANGA_GROW)
	if m.Capacity > MANGA_CAPACITY {
		m.Capacity = MANGA_CAPACITY
	} else if m.Capacity < 0 {
		m.Capacity = 0
	}
}

func (m manga) display() {
	fmt.Println("TODO: add diagnostics")
}

func (m manga) spread(x, y, maturity, order int) bool {
	//var look_counter = 0
	//var x_counter = -1
	//var y_counter = -1
	//var manga_created = false

	for xCounter := -1; xCounter < 2; xCounter++ {
		for yCounter := -1; yCounter < 2; yCounter++ {
			//var lookX = x + xCounter
			//var lookY = y + yCounter
			//var cell_index = findCellIndex(lookX, lookY)
		}
	}
	return true
}
