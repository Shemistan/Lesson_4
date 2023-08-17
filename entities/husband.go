package entities

type Husband struct {
	Human
}

func (h *Husband) PlayComp() {
	h.Happiness += HAPPINESS_FROM_PLAYING_PC
	h.Satiety -= SATIETY_REDUCTION
	h.houseResources.Totals.CompPlayed++
}

func (h *Husband) DoWork() {
	h.Satiety -= SATIETY_REDUCTION
	h.houseResources.MoneyIncome(150)
}

func (h *Husband) HusbandSimulation() {
	h.Check()

	if !h.IsAlive {
		return
	}

	if h.Satiety < 70 && h.houseResources.Food > 0 {
		h.Eat(30)
		h.toDo = "Eat"
	} else if h.Happiness < 95 {
		h.toDo = "HappinessUp"
		if h.Happiness < 80 {
			h.PlayComp()
		}
		h.PetCat()
	} else {
		h.DoWork()
		h.toDo = "Work"
	}
}
