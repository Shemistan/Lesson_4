package entities

type Wife struct {
	Human
}

func (w *Wife) Clean(amount int) {
	w.Satiety -= 10
	if amount > 100 {
		amount = 100
	}
	w.houseResources.DirtOutcome(amount)
}

func (w *Wife) BuyFood(amount int) {
	w.Satiety -= 10
	w.houseResources.FoodIncome(amount)
}

func (w *Wife) BuyCatFood(amount int) {
	w.Satiety -= 10
	w.houseResources.CatFoodIncome(amount)
}

func (w *Wife) BuyFurCoat() {
	w.Satiety -= 10
	w.Happiness += 60
	w.houseResources.MoneyOutcome(650)
	w.houseResources.Totals.FurCoatBought++
}

func (w *Wife) WifeLogic() {
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
				w.BuyFurCoat()
			}
			w.PetCat()
		} else {
			w.Clean(w.houseResources.Dirtiness)
			w.toDo = "Clean"
		}
	}
}
