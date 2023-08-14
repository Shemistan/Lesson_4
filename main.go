package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_4/home"
)

type Stats struct {
	EarnedMoney int64
	EatenFood   int64
	BoughtCoats int32
}

func main() {
	stats := Stats{
		EarnedMoney: 0,
		EatenFood:   0,
		BoughtCoats: 0,
	}

	myHome := home.Factory("John", "Marry")

	year(&myHome, &stats)
}

func year(home *home.Home, stats *Stats) {
	for i := 0; i < 365; i++ {
		if !home.IsFamilyAlive() {
			fmt.Println("Family died on...", i)
			break
		}

		wifeActionToday(home, stats)
		husbandActionToday(home, stats)

		home.NextDay()
	}
	showStats(*stats)
}

func showStats(stats Stats) {
	fmt.Println("Earned money:", stats.EarnedMoney)
	fmt.Println("Eaten food: ", stats.EatenFood)
	fmt.Println("Bought coats: ", stats.BoughtCoats)
}

func wifeActionToday(home *home.Home, stats *Stats) {
	switch {
	case home.IsWifeHungry():
		var food int32 = 30
		stats.EatenFood += int64(food)
		home.WifeEat(food)
	case home.IsTimeBuyProducts():
		home.BuyProducts(50)
	case home.IsWifeUnhappy():
		stats.BoughtCoats += 1
		home.BuyCoat()
	default:
		home.CleanHome()
	}
}

func husbandActionToday(home *home.Home, stats *Stats) {
	switch {
	case home.IsHusbandHungry():
		var food int32 = 20
		stats.EatenFood += int64(food)
		home.HusbandEat(food)
	case home.IsHusbandUnhappy():
		home.PlayComputer()
	default:
		stats.EarnedMoney += int64(home.EarnMoney())
	}
}
