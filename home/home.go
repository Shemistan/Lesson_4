package home

import "fmt"

type Home struct {
	Money      int64
	Meal       int64
	Grime      int64
	CountMoney int64
	CountCoat  int64
	CountMeal  int64
}

type Human struct {
	Name      string
	Fullness  int64
	Happiness int64
	Home      *Home
}

func IsDead(h *Human) bool {
	// Возвращает то что человек умер или нет
	if h.Fullness <= 0 {
		fmt.Println("Умер", h.Name, "от голода")
		return true
	}
	if h.Happiness < 10 {
		fmt.Println("Умер", h.Name, "от депресси")
		return true
	}
	return false
}

func (h *Human) HumanEat() {
	if h.Home.Meal > 30 {
		h.Fullness += 30
		h.Home.Meal -= 30
		h.Home.CountMeal += 30
		fmt.Println(h.Name, "Употреблял еду")
	} else if h.Home.Meal < 30 {
		h.Fullness += h.Home.Meal
		h.Home.CountMeal += h.Home.Meal
		h.Home.Meal -= h.Home.Meal
		fmt.Println(h.Name, "Употреблял еду")
	} else {
		fmt.Println("В доме не осталась еда")
	}
}
