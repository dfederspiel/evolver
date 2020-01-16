package main

import (
	evolver "evolver/internal"
	"fmt"
)

func main(){
	e := evolver.NewEvolvee("amoeba", 123, 0,0)
	fmt.Println(e.Display())

	a := evolver.NewAnimal("Mother", "Father", "Male", 1, e)
	fmt.Println(a.Display())

	a2 := evolver.NewManga(1, 0,10)
	fmt.Println(a2.Display())
	a2.Age()
	fmt.Println(a2.Display())
	a2.Grow()
	fmt.Println(a2.Display())
	a2.Spread(0,0,2,4)

}
