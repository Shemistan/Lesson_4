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
	fmt.Println(husband.Name, wife.Name)
	husband.Eat(husband.Name, &home)
	wife.Eat(wife.Name, &home)
	fmt.Println(home.Food)

}
