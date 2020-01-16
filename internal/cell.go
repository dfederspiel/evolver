package evolver

type cell struct {
	evolvee
	Fertile bool
	Manga manga
	Predo predo
	//Danto danto
	EggPredo eggPredo
	FertilityCounter int
}

func NewCell(e evolvee) cell {
	return cell{
		evolvee: e,
		Fertile: true,
	}
}

func (c cell) SetManga(m manga) {

}

func (c cell) SetPredo(p predo){

}

func (c cell) SetEggPredo(e eggPredo){

}

func (c cell) Age(){
	c.FertilityCounter += 1
	if c.FertilityCounter > CELL_FERTILE {
		c.Fertile = true
	}
}

func (c cell) RemoveManga(){

}

func (c cell) RemovePredo(){

}

func (c cell) RemoveEggPredo(){

}