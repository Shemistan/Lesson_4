package entities

type Husband struct {
	Human
}

func (h *Husband) PlayComp() {
	h.Happiness += 20
	h.Satiety -= 10
	h.houseResources.Totals.CompPlayed++
}

func (h *Husband) Work() {
	h.Satiety -= 10
	h.houseResources.MoneyIncome(150)
}

func (h *Husband) HusbandLogic() {
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
		h.Work()
		h.toDo = "Work"
	}
}
