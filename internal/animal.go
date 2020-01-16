package evolver

type animal struct {
	Mother string
	Father string
	Sex string
	TickLastMoved int
	evolvee
}

func NewAnimal(mother string, father string, sex string, tickLastMoved int, evolvee evolvee) *animal {
	return &animal{Mother: mother, Father: father, Sex: sex, TickLastMoved: tickLastMoved, evolvee: evolvee}
}


