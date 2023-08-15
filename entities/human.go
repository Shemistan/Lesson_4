package entities

import "fmt"

type Human struct {
	Entity
	Happiness int
}

func (h *Human) Init(name string, houseResources *Resources) {
	h.Entity.Init(name, houseResources)
	h.Happiness = 100
}

func (h *Human) Eat(amount int) {
	if h.houseResources.Food < amount {
		amount = h.houseResources.Food
	}
	h.Satiety += amount
	h.houseResources.FoodOutcome(amount)
}

func (h *Human) PetCat() {
	h.Happiness += 5
}

func (h *Human) Check() {
	h.Entity.Check()

	if h.houseResources.Dirtiness > 90 {
		h.Happiness -= 10
	}

	if h.Happiness < 10 {
		h.IsAlive = false
	}
}

func (h *Human) GetData() {
	h.Entity.GetData()
	fmt.Println("	Happiness: ", h.Happiness)
}
