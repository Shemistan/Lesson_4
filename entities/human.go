package entities

import "fmt"

type Human struct {
	Entity
	Happiness uint32
}

func (h *Human) Init(name string, houseResources *Resources) {
	h.Entity.Init(name, houseResources)
	h.Happiness = START_HEPPINESS
}

func (h *Human) Eat(amount uint32) {
	if h.houseResources.Food < amount {
		amount = h.houseResources.Food
	}
	h.Satiety += amount
	h.houseResources.FoodDecrease(amount)
}

func (h *Human) PetCat() {
	h.Happiness += HAPPINESS_FROM_PETTING_A_CAT
}

func (h *Human) Check() {
	h.Entity.Check()

	if h.houseResources.Dirtiness > 90 {
		h.Happiness -= HAPPINESS_REDUCTION_CUZ_DIRTY
	}

	if h.Happiness < 10 {
		h.IsAlive = false
	}
}

func (h *Human) GetData() {
	h.Entity.PrintData()
	fmt.Println("	Happiness: ", h.Happiness)
}
