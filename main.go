package main

import (
	"fmt"
	h "github.com/Shemistan/Lesson_4/home"
	m "github.com/Shemistan/Lesson_4/men"
	w "github.com/Shemistan/Lesson_4/women"
)

func main() {
	home := h.Home{
		Money:      100,
		Meal:       50,
		Grime:      0,
		CountMoney: 0,
		CountCoat:  0,
		CountMeal:  0,
	}

	aleks := h.Human{
		Name:      "Aleksandr",
		Fullness:  30,
		Happiness: 100,
		Home:      &home,
	}

	men := m.Men{Human: aleks}

	nastya := h.Human{
		Name:      "Anastasiya",
		Fullness:  30,
		Happiness: 100,
		Home:      &home,
	}

	women := w.Women{Human: nastya}
	for i := 1; i < 365; i++ {
		switch {
		case men.Human.Fullness < 20:
			men.Human.HumanEat()
		case women.Human.Fullness < 20:
			women.Human.HumanEat()
		case home.Grime >= 80:
			men.Human.Happiness -= 10
			women.Human.Happiness -= 10
			women.CleanHome()
		case home.Money < 350:
			men.GoWork()
		case home.Meal <= 25:
			women.BuyProducts()
		case men.Human.Happiness < 15:
			men.PlayComputer()
		case women.Human.Happiness < 15:
			women.BuyCoat()
		}
		if h.IsDead(&aleks) || h.IsDead(&nastya) {
			break
		}
		home.Grime += 5
	}

	fmt.Println("Употребление еды: ", home.CountMeal)
	fmt.Println("Количество денег: ", home.CountMoney)
	fmt.Println("Количество шубы: ", home.CountCoat)
}
