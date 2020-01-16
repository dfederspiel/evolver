package evolver

import "fmt"

type manga struct {
	Life int
	Capacity float32
	evolvee
}

func NewManga(id, x, y int) manga {
	return manga{
		Life:     lifeExpectancy(MANGA_LIFE),
		Capacity: MANGA_CAPACITY,
		evolvee: NewEvolvee("manga", id, x, y),
	}
}

func (m *manga) Age() bool {
	m.Life -= 1
	m.Maturity += 1

	if m.Life <= 1 {
		return true
	}
	return false
}

func (m *manga) Grow() {
	m.Capacity = m.Capacity * (1 + MANGA_GROW)
	if m.Capacity > MANGA_CAPACITY {
		m.Capacity = MANGA_CAPACITY
	} else if m.Capacity < 0 {
		m.Capacity = 0
	}
}

func (m *manga) Display() string {
	return fmt.Sprintf("Class=%s ID=%d XY=%s Maturity=%d, Life=%d, Capacity=%f", m.Classification, m.Id, m.XY(), m.Maturity, m.Life, m.Capacity)
}

func (m *manga) Spread(x, y, maturity, order int) bool {
	var lookCounter = 0
	var xCounter = -1
	var yCounter = -1
	var mangaCreated = false

	for xCounter = -1; xCounter < 2; xCounter++ {
		for yCounter = -1; yCounter < 2; yCounter++ {
			var lookX = x + xCounter
			var lookY = y + yCounter
			var cellIndex = findCellIndex(lookX, lookY)
			var cell = Grid[cellIndex]
			if cell.Id == 0 && maturity > MANGA_MATURITY && cell.Fertile {
				manga_list_counter += 1
				manga_live_counter += 1

				m := NewManga(manga_list_counter, lookX, lookY)
				Manga = append(Manga, m)
				cell.SetManga(m)
				mangaCreated = true
			} else {
				lookCounter += 1
				yCounter += 1


			}
		}
		xCounter += 1
		yCounter -= 1
	}
	return mangaCreated
}
