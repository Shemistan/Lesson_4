package main

import (
	"fmt"

	"github.com/kamolisrailov/Lesson_4.git/family"
	"github.com/kamolisrailov/Lesson_4.git/home"
	"github.com/kamolisrailov/Lesson_4.git/human"
)

func main() {

	var husband = family.Husband{
		Human: human.Human{Name: "Mike", Fullnes: 30, Happiness: 100},
	}

	var wife = family.Wife{
		Human: human.Human{Name: "Sara", Fullnes: 30, Happiness: 100},
	}
	var home = home.Home{Money: 100, Food: 50, Dust: 0}

	for i := 1; i <= 365; i++ {
		fmt.Printf("Start of the %d day----------------------\n\n ", i)

		husband.DailyActions(&home)
		wife.DailyActions(&home)
		home.AddDirtiness()
		fmt.Println()
		fmt.Printf("End of the %d day----------------------\n\n", i)
	}
	fmt.Printf("Sara`s Fullness: %d,   Happiness: %d,   Bought  Coat: %d\n", wife.Fullnes, wife.Happiness, wife.CountBuyCoat)
	fmt.Printf("Mike`s Fullness: %d,   Happiness: %d,   Played Computer: %d\n", husband.Fullnes, husband.Happiness, husband.CountPlayComputer)
	fmt.Printf("Food: %d,   Money: %d,   Dust: %d\n", home.Food, home.Money, home.Dust)
}
