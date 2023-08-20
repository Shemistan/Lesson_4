package home

import (
	"fmt"
	"github.com/Shemistan/Lesson_4/human"
)

const (
	initialMoney int32 = 100
	initialFood  int32 = 50
	initialDirt  int32 = 0
)

type Home struct {
	money   int32
	food    int32
	dirt    int32
	wife    human.Wife
	husband human.Husband
}

// General Family Actions

func (home *Home) CalculateFamilyPropertiesForToday(dirt int32) {
	home.husband.CalculateHappinessForToday(home.dirt)
	home.wife.CalculateHappinessForToday(home.dirt)
	home.dirt += dirt
}
func (home *Home) IsTimeBuyProducts() bool {
	return home.food <= 50
}

func (home *Home) IsFamilyAlive() bool {
	return home.wife.IsAlive && home.husband.IsAlive
}

func (home *Home) GetDirtPoints() int32 {
	return home.dirt
}

func (home *Home) eatFromFridge(human *human.Human, food int32) {
	if food > home.food {
		fmt.Println("Not enough food in fridge")
	}
	home.food -= food
	human.Eat(food)
}

// Husband Actions

func (home *Home) IsHusbandUnhappy() bool {
	return home.husband.IsFamilyMemberUnhappy()
}

func (home *Home) IsHusbandHungry() bool {
	return home.husband.Satiety == 0
}

func (home *Home) HusbandEat(food int32) {
	home.eatFromFridge(&home.husband.Human, food)
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
	return home.wife.IsFamilyMemberUnhappy()
}

func (home *Home) IsWifeHungry() bool {
	return home.wife.Satiety == 0
}

func (home *Home) WifeEat(food int32) {
	home.eatFromFridge(&home.wife.Human, food)
}

func (home *Home) BuyProducts(money int32) {
	if money > home.money {
		fmt.Println("Not enough money to buy products")
	}
	home.food += home.wife.BuyProducts(money)
	home.money -= money
}

func (home *Home) BuyCoat() error {
	err, remainingMoney := home.wife.BuyCoat(home.money)
	if err != nil {
		return err
	}
	home.money = remainingMoney
	return nil
}

func (home *Home) CleanHome() {
	home.dirt = home.wife.CleanUp(home.dirt)
}

// Home CreateHome

func CreateHome(husbandName string, wifeName string) Home {
	return Home{
		money:   initialMoney,
		food:    initialFood,
		dirt:    initialDirt,
		husband: human.CreateHusband(husbandName),
		wife:    human.CreateWife(wifeName),
	}
}
