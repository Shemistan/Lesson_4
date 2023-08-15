package entities

import "fmt"

type Resources struct {
	Money     int
	CatFood   int
	Food      int
	Dirtiness int
	Totals    *Statistics
}

func (r *Resources) GetResources() {
	fmt.Println("[RESOURCES]")
	fmt.Println("	Money: ", r.Money)
	fmt.Println("	CatFood: ", r.CatFood)
	fmt.Println("	Food: ", r.Food)
	fmt.Println("	Dirtiness: ", r.Dirtiness)
}

func (r *Resources) Init(statistics *Statistics) {
	r.Money = 100
	r.CatFood = 30
	r.Food = 50
	r.Dirtiness = 0
	r.Totals = statistics
}

func (r *Resources) MoneyIncome(amount int) {
	r.Money += amount
	r.Totals.MoneyEarned += amount
}

func (r *Resources) MoneyOutcome(amount int) {
	r.Money -= amount
	r.Totals.MoneySpent += amount
}

func (r *Resources) FoodIncome(amount int) {
	r.Food += amount
	r.MoneyOutcome(amount)
}

func (r *Resources) FoodOutcome(amount int) {
	r.Food -= amount
	r.Totals.FoodEaten += amount
}

func (r *Resources) CatFoodIncome(amount int) {
	r.CatFood += amount
	r.MoneyOutcome(amount)
}

func (r *Resources) CatFoodOutcome(amount int) {
	r.CatFood -= amount
	r.Totals.CatFoodEaten += amount
}

func (r *Resources) DirtIncome(amount int) {
	r.Dirtiness += amount
}

func (r *Resources) DirtOutcome(amount int) {
	r.Dirtiness -= amount
}
