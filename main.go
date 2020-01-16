package main

import (
	evolvee "evolver/internal"
	"fmt"
)

func main(){
	e := evolvee.New("manga", 123, 0,0)
	fmt.Println(e.Classification)

}
