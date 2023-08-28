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

	for i := 1; i <= 10; i++ {
		fmt.Printf("Start of the %d day----------------------\n\n ", i)
		home.AddDirtiness()
		fmt.Println(husband.Name, wife.Name)
		husband.Eat(&home)
		wife.Eat(&home)
		wife.BuyGrocery(&home)
		wife.CleanHouse(&home)
		husband.GoToWork(&home)
		husband.PlayComputer()
		fmt.Println()
		fmt.Printf("End of the %d day----------------------\n\n", i)
	}
	fmt.Printf("Food: %d,   Money: %d,   Dust: %d\n", home.Food, home.Money, home.Dust)
}
