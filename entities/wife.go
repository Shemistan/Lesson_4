package entities

type Wife struct {
	Human
}

func (w *Wife) Clean(amount uint32) {
	w.Satiety -= SATIETY_REDUCTION
	if amount > 100 {
		amount = 100
	}
	w.houseResources.DirtDecrease(amount)
}

func (w *Wife) BuyFood(amount uint32) {
	w.Satiety -= SATIETY_REDUCTION
	w.houseResources.FoodIncrease(amount)
}

func (w *Wife) BuyCatFood(amount uint32) {
	w.Satiety -= SATIETY_REDUCTION
	w.houseResources.CatFoodIncrease(amount)
}

func (w *Wife) BuyCoat() {
	w.Satiety -= SATIETY_REDUCTION
	w.Happiness += 60
	w.houseResources.MoneyOutcome(COAT_COST)
	w.houseResources.Totals.CoatBought++
}

func (w *Wife) WifeSimulation() {
	w.Check()

	if w.IsAlive {
		if w.houseResources.Food < 60 {
			w.BuyFood(100 - w.houseResources.Food)
			w.toDo = "BuyFood"
		} else if w.Satiety < 70 && w.houseResources.CatFood > 0 {
			w.Eat(30)
			w.toDo = "Eat"
		} else if w.houseResources.CatFood < 20 {
			w.BuyCatFood(100 - w.houseResources.CatFood)
			w.toDo = "BuyCatFood"
		} else if w.Happiness < 95 {
			w.toDo = "HappinessUp"
			if w.Happiness < 40 {
				w.BuyCoat()
			}
			w.PetCat()
		} else {
			w.Clean(w.houseResources.Dirtiness)
			w.toDo = "Clean"
		}
	}
}
