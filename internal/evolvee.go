package evolvee

import "fmt"

type evolvee struct {
	 Classification string
	 Id int
	 X int
	 Y int
	 Maturity int
}

func xy(e evolvee) string {
	return fmt.Sprintf("%d,%d", e.X, e.Y)
}

func New(classification string, id, x, y int) *evolvee {
	e := new(evolvee)
	e.Classification = classification
	e.Id = id
	e.X = x
	e.Y = y
	e.Maturity = 1
	return e
}