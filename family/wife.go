package family

import (
	"fmt"

	"github.com/kamolisrailov/Lesson_4.git/home"
	"github.com/kamolisrailov/Lesson_4.git/human"
)

const (
	Food          = 60
	Dust          = 100
	Coat          = 350
	WifeHappiness = 60
)

type Wife struct {
	human.Human
	CountBuyCoat int
}

func (w *Wife) BuyGrocery(home *home.Home) {
	w.Fullnes -= Fullness
	home.Food += Food
	home.Money -= Food
	fmt.Printf("%s went to the market and bought %d food\n", w.Name, Food)
	fmt.Printf("%s's fullness: %d\n", w.Name, w.Fullnes)
}

func (w *Wife) CleanHouse(home *home.Home) {
	w.Fullnes -= Fullness
	home.Dust -= Dust
	fmt.Printf("%s has cleaned house\n", w.Name)
	fmt.Printf("House dirtiness is: : %d\n", home.Dust)
}

func (w *Wife) BuyCoat(home *home.Home) {
	w.Happiness += WifeHappiness
	home.Money -= Coat
	w.CountBuyCoat += 1
	fmt.Printf("%s has bought coat for $%d and her Happiness increased to %d:\n", w.Name, Coat, WifeHappiness)
	fmt.Printf("Amount of money: %d\n", home.Money)
}

func (w *Wife) DailyActions(home *home.Home) int {
	if w.Fullnes-Fullness < 0 {
		w.Eat(home)
		return 0
	}
	if home.Food < FoodMin || home.Money < MoneyMin {
		w.BuyGrocery(home)
		return 0
	}

	if w.Happiness <= 10 {
		w.BuyCoat(home)
		return 0
	}

	if home.Dust > 90 {
		w.Happiness -= 10
		w.CleanHouse(home)
		return 0
	}
	return 0
}
