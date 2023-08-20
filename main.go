package main

import (
	"errors"
	"fmt"
	"github.com/Shemistan/Lesson_4/home"
)

const (
	countOfDays                  int32 = 365
	countProductsPurchasedAtTime int32 = 50
	countOfFoodForWife           int32 = 30
	countOfFoodForHusband        int32 = 20
	dirtPointsPerDay             int32 = 5

	nameOfHusband string = "John"
	nameOfWife    string = "Marry"

	earnedMoneyMessage        string = "Earned money: "
	eatenFoodMessage          string = "Eaten food: "
	boughtCoatsMessage        string = "Bought coats: "
	countOfPassedDaysMessage  string = "Count of passed days: "
	somethingWentWrongMessage string = "Something went wrong in simulation: "
)

type Stats struct {
	EarnedMoney       int32
	EatenFood         int32
	BoughtCoats       int32
	CountOfPassedDays int32
}

func main() {
	stats := Stats{
		EarnedMoney:       0,
		EatenFood:         0,
		BoughtCoats:       0,
		CountOfPassedDays: 0,
	}

	myHome := home.CreateHome(nameOfHusband, nameOfWife)

	err := RunSimulation(&myHome, &stats, countOfDays)

	if err != nil {
		DisplayError(err)
	}

	ShowStats(stats)
}

func RunSimulation(home *home.Home, stats *Stats, countOfDays int32) error {
	for i := int32(0); i < countOfDays; i++ {
		stats.CountOfPassedDays++

		if !home.IsFamilyAlive() {
			return errors.New("Family died")
		}

		err := doWifeActionForToday(home, stats)

		if err != nil {
			return err
		}

		err = doHusbandActionForToday(home, stats)

		if err != nil {
			return err
		}

		err = home.CalculateFamilyPropertiesForToday(dirtPointsPerDay)

		if err != nil {
			return err
		}

	}
	return nil
}

func ShowStats(stats Stats) {
	fmt.Println(earnedMoneyMessage, stats.EarnedMoney)
	fmt.Println(eatenFoodMessage, stats.EatenFood)
	fmt.Println(boughtCoatsMessage, stats.BoughtCoats)
	fmt.Println(countOfPassedDaysMessage, stats.CountOfPassedDays)
}

func DisplayError(err error) {
	fmt.Println(somethingWentWrongMessage, err)
}

func doWifeActionForToday(home *home.Home, stats *Stats) error {
	switch {
	case home.IsWifeHungry():
		stats.EatenFood += countOfFoodForWife
		return home.WifeEat(countOfFoodForWife)
	case home.IsTimeBuyProducts():
		return home.BuyProducts(countProductsPurchasedAtTime)
	case home.IsWifeUnhappy():
		stats.BoughtCoats++
		return home.BuyCoat()
	default:
		return home.CleanHome()
	}
}

func doHusbandActionForToday(home *home.Home, stats *Stats) error {
	switch {
	case home.IsHusbandHungry():
		stats.EatenFood += countOfFoodForHusband
		return home.HusbandEat(countOfFoodForHusband)
	case home.IsHusbandUnhappy():
		return home.PlayComputer()
	default:
		err, money := home.EarnMoney()
		stats.EarnedMoney += money
		return err
	}
}
