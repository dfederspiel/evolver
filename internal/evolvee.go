package evolver

import "fmt"

type evolvee struct {
	 Classification string
	 Id int
	 X int
	 Y int
	 Maturity int
}

func (e evolvee) XY() string {
	return fmt.Sprintf("%d,%d", e.X, e.Y)
}

func (e evolvee) Display() string {
	return fmt.Sprintf("Class=%s ID=%d XY=%s Maturity=%d", e.Classification, e.Id, e.XY(), e.Maturity)
}

func NewEvolvee(classification string, id, x, y int) *evolvee {
	e := new(evolvee)
	e.Classification = classification
	e.Id = id
	e.X = x
	e.Y = y
	e.Maturity = 1
	return e
}