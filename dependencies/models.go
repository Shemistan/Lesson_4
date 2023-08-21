package Lesson_4

import "fmt"

type House struct {
	Money int16
	Food  int16
	Dirt  int16
}

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
	h.Satiety -= 10
	h.Happiness += 20
}

func (h *Husband) GoWork() {
	h.Satiety -= 10
	h.House.Money += 150
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
	w.Satiety -= 10
	if foodAmount > w.House.Money {
		fmt.Println("Not enough money")
	} else {
		w.House.Money -= foodAmount
		w.House.Food += foodAmount
	}
}

func (w *Wife) BuyCoat() {
	w.Satiety -= 10
	w.House.Money -= 350
	w.Happiness += 60
}

func (w *Wife) CleanHouse() {
	w.Satiety -= 10
	w.House.Dirt -= 100
}
