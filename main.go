package main

import (
	evolver "evolver/internal"
	"fmt"
)

func main(){
	e := evolver.NewEvolvee("manga", 123, 0,0)
	fmt.Println(e.Display())

	a := evolver.NewAnimal("Mother", "Father", "Male", 1, *e)
	fmt.Println(a.Display())

}
