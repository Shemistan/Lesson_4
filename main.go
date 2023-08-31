package main

import (
	"fmt"

	"github.com/kamolisrailov/Lesson_4.git/family"
	"github.com/kamolisrailov/Lesson_4.git/home"
	"github.com/kamolisrailov/Lesson_4.git/human"
)

const (
	BaseFullness  = 30
	BaseHappiness = 100
	BaseMoney     = 100
	BaseFood      = 50
	BaseDust      = 0
)

func main() {

	var year int
	fmt.Print("Type number of days to live in range (1 , 365):  ")
	fmt.Scan(&year)

	var husband = family.Husband{
		Human: human.Human{Name: "Mike", Fullnes: BaseFullness, Happiness: BaseHappiness},
	}

	var wife = family.Wife{
		Human: human.Human{Name: "Sara", Fullnes: BaseFullness, Happiness: BaseHappiness},
	}
	var home = home.Home{Money: BaseMoney, Food: BaseFood, Dust: BaseDust}

	for i := 1; i <= year; i++ {
		fmt.Printf("Start of the %d day----------------------\n\n ", i)

		husband.LivingDay(&home)
		wife.LivingDay(&home)
		home.AddDirtiness()
		fmt.Println()
		fmt.Printf("End of the %d day----------------------\n\n", i)
	}
	fmt.Printf("Sara`s Fullness: %d,   Happiness: %d,   Bought  Coat: %d\n", wife.Fullnes, wife.Happiness, wife.CountBuyCoat)
	fmt.Printf("Mike`s Fullness: %d,   Happiness: %d,   Played Computer: %d\n", husband.Fullnes, husband.Happiness, husband.CountPlayComputer)
	fmt.Printf("Food: %d,   Money: %d,   Dust: %d\n", home.Food, home.Money, home.Dust)
}
