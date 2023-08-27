package home

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/constants"
)

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
	if h.Fullness <= constants.FulnessThreshold {
		fmt.Println("Умер", h.Name, "от голода")
		return true
	}
	if h.Happiness < constants.HappinessThreshold {
		fmt.Println("Умер", h.Name, "от депресси")
		return true
	}
	return false
}

func (h *Human) HumanEat() {
	if h.Home.Meal > constants.StartFulness {
		h.Fullness += constants.IncreaseFullnessByEat
		h.Home.Meal -= constants.StartFulness
		h.Home.CountMeal += constants.IncreaseFullnessByEat
		fmt.Println(h.Name, "Употреблял еду")
	} else if h.Home.Meal < constants.StartFulness {
		h.Fullness += h.Home.Meal
		h.Home.CountMeal += h.Home.Meal
		h.Home.Meal -= h.Home.Meal
		fmt.Println(h.Name, "Употреблял еду")
	} else {
		fmt.Println("В доме не осталась еда")
	}
}

func (h *Home)ShowStatistic() {
	fmt.Println("Употребление еды: ", h.CountMeal)
	fmt.Println("Количество денег: ", h.CountMoney)
	fmt.Println("Количество шубы: ", h.CountCoat)
}

