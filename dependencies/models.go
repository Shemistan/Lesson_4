package Lesson_4

import "fmt"

type House struct {
	Money int16
	Food  int16
	Dirt  int16
}

const (
	DailySatietyReduction                 int16 = 10
	HappinessIncreaseAfterPlayingComputer int16 = 20
	HappinessIncreaseAfterBuyingCoat      int16 = 60
	SalaryDaily                           int16 = 150
	PriceOfCoat                           int16 = 350
	DirtAmount                            int16 = 100
)

func Death() {
	fmt.Print("Person died")
}

type Person struct {
	Name      string
	Satiety   int16
	Happiness int16
}

type Husband struct {
	Person
	House *House
}

func (h *Husband) HusbandEat(food int16) {
	if food > h.House.Food {
		fmt.Println("Not enough food husband")
	} else {
		h.Satiety += food
		h.House.Food -= food
	}
}

func (h *Husband) PlayComputer() {
	h.Satiety -= DailySatietyReduction
	h.Happiness += HappinessIncreaseAfterPlayingComputer
}

func (h *Husband) GoWork() {
	h.Satiety -= DailySatietyReduction
	h.House.Money += SalaryDaily
}

type Wife struct {
	Person
	House *House
}

func (w *Wife) WifeEat(food int16) {
	if food > w.House.Food {
		fmt.Println("Not enough food wife")
	} else {
		w.Satiety += food
		w.House.Food -= food
	}
}

func (w *Wife) BuyGroceries(foodAmount int16) {
	w.Satiety -= DailySatietyReduction
	if foodAmount > w.House.Money {
		fmt.Println("Not enough money")
	} else {
		w.House.Money -= foodAmount
		w.House.Food += foodAmount
	}
}

func (w *Wife) BuyCoat() {
	w.Satiety -= DailySatietyReduction
	w.House.Money -= PriceOfCoat
	w.Happiness += HappinessIncreaseAfterBuyingCoat
}

func (w *Wife) CleanHouse() {
	w.Satiety -= DailySatietyReduction
	w.House.Dirt -= DirtAmount
}
