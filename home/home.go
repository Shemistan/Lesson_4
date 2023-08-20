package home

import (
	"errors"
	"github.com/Shemistan/Lesson_4/human"
)

const (
	initialMoney int32 = 100
	initialFood  int32 = 50
	initialDirt  int32 = 0

	notEnoughFoodInFridge           = "not enough food in fridge"
	notEnoughMoneyToProductsMessage = "not enough money to buy products"
)

type Home struct {
	money   int32
	food    int32
	dirt    int32
	wife    human.Wife
	husband human.Husband
}

// General Family Actions

func (home *Home) CalculateFamilyPropertiesForToday(dirt int32) error {
	err := home.husband.CalculateHappinessForToday(home.dirt)

	if err != nil {
		return err
	}

	err = home.wife.CalculateHappinessForToday(home.dirt)

	if err != nil {
		return err
	}

	home.dirt += dirt

	return nil
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

func (home *Home) eatFromFridge(human *human.Human, food int32) error {
	if food > home.food {
		return errors.New(notEnoughFoodInFridge)
	}

	home.food -= food
	human.Eat(food)

	return nil
}

// Husband Actions

func (home *Home) IsHusbandUnhappy() bool {
	return home.husband.IsFamilyMemberUnhappy()
}

func (home *Home) IsHusbandHungry() bool {
	return home.husband.Satiety == 0
}

func (home *Home) HusbandEat(food int32) error {
	return home.eatFromFridge(&home.husband.Human, food)
}

func (home *Home) PlayComputer() error {
	return home.husband.PlayComputer()
}

func (home *Home) EarnMoney() (err error, earnedMoney int32) {
	err, earnedMoney = home.husband.EarnMoney()

	if err != nil {
		return err, 0
	}

	home.money += earnedMoney

	return err, earnedMoney
}

// Wife Actions

func (home *Home) IsWifeUnhappy() bool {
	return home.wife.IsFamilyMemberUnhappy()
}

func (home *Home) IsWifeHungry() bool {
	return home.wife.Satiety == 0
}

func (home *Home) WifeEat(food int32) error {
	return home.eatFromFridge(&home.wife.Human, food)
}

func (home *Home) BuyProducts(money int32) error {
	if money > home.money {
		return errors.New(notEnoughMoneyToProductsMessage)
	}

	err, products := home.wife.BuyProducts(money)

	if err != nil {
		return err
	}

	home.food += products
	home.money -= money

	return nil
}

func (home *Home) BuyCoat() error {
	err, remainingMoney := home.wife.BuyCoat(home.money)

	if err != nil {
		return err
	}

	home.money = remainingMoney

	return nil
}

func (home *Home) CleanHome() error {
	err, dirtAfterClean := home.wife.CleanUp(home.dirt)

	if err != nil {
		return err
	}

	home.dirt = dirtAfterClean

	return nil
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
