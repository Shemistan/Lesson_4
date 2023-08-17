package entities

import "fmt"

type Resources struct {
	Money     uint32
	CatFood   uint32
	Food      uint32
	Dirtiness uint32
	Totals    *Statistics
}

func (r *Resources) PrintResources() {
	fmt.Println("[RESOURCES]")
	fmt.Println("	Money: ", r.Money)
	fmt.Println("	CatFood: ", r.CatFood)
	fmt.Println("	Food: ", r.Food)
	fmt.Println("	Dirtiness: ", r.Dirtiness)
}

func (r *Resources) Init(statistics *Statistics) {
	r.Money = START_MONEY
	r.CatFood = START_CAT_FOOD
	r.Food = START_FOOD
	r.Dirtiness = START_DIRTINESS
	r.Totals = statistics
}

func (r *Resources) MoneyIncome(amount uint32) {
	r.Money += amount
	r.Totals.MoneyEarned += amount
}

func (r *Resources) MoneyOutcome(amount uint32) {
	r.Money -= amount
	r.Totals.MoneySpent += amount
}

func (r *Resources) FoodIncrease(amount uint32) {
	r.Food += amount
	r.MoneyOutcome(amount)
}

func (r *Resources) FoodDecrease(amount uint32) {
	r.Food -= amount
	r.Totals.FoodEaten += amount
}

func (r *Resources) CatFoodIncrease(amount uint32) {
	r.CatFood += amount
	r.MoneyOutcome(amount)
}

func (r *Resources) CatFoodDecrease(amount uint32) {
	r.CatFood -= amount
	r.Totals.CatFoodEaten += amount
}

func (r *Resources) DirtIncrease(amount uint32) {
	r.Dirtiness += amount
}

func (r *Resources) DirtDecrease(amount uint32) {
	r.Dirtiness -= amount
}
