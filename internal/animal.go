package evolver

type animal struct {
	mother string
	father string
	sex string
	tickLastMoved int
	evolvee
}

func NewAnimal(mother string, father string, sex string, tickLastMoved int, evolvee evolvee) *animal {
	return &animal{mother: mother, father: father, sex: sex, tickLastMoved: tickLastMoved, evolvee: evolvee}
}


