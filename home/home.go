package home

import (
	"fmt"
	"github.com/Shemistan/Lesson_4/human"
)

type DirtPoints struct {
	current     int32
	incremental int32
}

type Home struct {
	money      int32
	fridge     int32
	dirtPoints DirtPoints
	wife       human.Wife
	husband    human.Husband
}

// General Family Actions

func (home *Home) NextDay() {
	home.husband.CalculateHappinessLevel(home.dirtPoints.current)
	home.wife.CalculateHappinessLevel(home.dirtPoints.current)
	home.dirtPoints.current += home.dirtPoints.incremental
}
func (home *Home) IsTimeBuyProducts() bool {
	return home.fridge <= 50
}

func (home *Home) IsFamilyAlive() bool {
	return home.wife.IsAlive && home.husband.IsAlive
}

func (home *Home) DirtPoints() int32 {
	return home.dirtPoints.current
}

func (home *Home) eatFromFridge(human *human.Human, foodPoint int32) {
	if foodPoint > home.fridge {
		fmt.Println("Not enough food in fridge")
	}
	home.fridge -= foodPoint
	human.Eat(foodPoint)
}

// Husband Actions

func (home *Home) IsHusbandUnhappy() bool {
	return home.husband.Happiness <= 20
}

func (home *Home) IsHusbandHungry() bool {
	return home.husband.Satiety == 0
}

func (home *Home) HusbandEat(foodPoint int32) {
	home.eatFromFridge(&home.husband.Human, foodPoint)
}

func (home *Home) PlayComputer() {
	home.husband.PlayComputer()
}

func (home *Home) EarnMoney() int32 {
	earnedMoney := home.husband.EarnMoney()
	home.money += earnedMoney
	return earnedMoney
}

// Wife Actions

func (home *Home) IsWifeUnhappy() bool {
	return home.wife.Happiness <= 20
}

func (home *Home) IsWifeHungry() bool {
	return home.wife.Satiety == 0
}

func (home *Home) WifeEat(foodPoint int32) {
	home.eatFromFridge(&home.wife.Human, foodPoint)
}

func (home *Home) BuyProducts(money int32) {
	if money > home.money {
		fmt.Println("Not enough money to buy products")
	}
	home.fridge += home.wife.BuyProducts(money)
	home.money -= money
}

func (home *Home) BuyCoat() {
	home.wife.BuyCoat(&home.money)
}

func (home *Home) CleanHome() {
	home.dirtPoints.current = home.wife.CleanUp(home.dirtPoints.current)
}

// Home Factory

func Factory(husbandName string, wifeName string) Home {
	return Home{
		money:  100,
		fridge: 50,
		dirtPoints: DirtPoints{
			current:     0,
			incremental: 5,
		},
		husband: human.HusbandFactory(husbandName),
		wife:    human.WifeFactory(wifeName),
	}
}
