package main

import (
	c "github.com/Shemistan/Lesson_4/constants"
	h "github.com/Shemistan/Lesson_4/home"
	m "github.com/Shemistan/Lesson_4/men"
	w "github.com/Shemistan/Lesson_4/women"
)

func Run() {
	var firstDay int = 1

	home := h.Home{
		Money:      c.StartMoney,
		Meal:       c.StartMeal,
		Grime:      c.StartGrime,
		CountMoney: 0,
		CountCoat:  0,
		CountMeal:  0,
	}

	aleks := h.Human{
		Name:      "Aleksandr",
		Fullness:  c.StartFulness,
		Happiness: c.StartMoney,
		Home:      &home,
	}

	men := m.Men{Human: aleks}

	nastya := h.Human{
		Name:      "Anastasiya",
		Fullness:  c.StartFulness,
		Happiness: c.StartMoney,
		Home:      &home,
	}

	women := w.Women{Human: nastya}


	for i := firstDay; i < c.DaysInYear; i++ {
		switch {
		case men.Human.Fullness < c.MinFulnessBorder:
			men.Human.HumanEat()
		case women.Human.Fullness < c.MinFulnessBorder:
			women.Human.HumanEat()
		case home.Grime >= c.MinGrimeBorder:
			men.Human.Happiness -= c.DecreaseHappinessByGrime 
			women.Human.Happiness -= c.DecreaseHappinessByGrime 
			women.CleanHome()
		case home.Money < c.MinMoneyBorder:
			men.GoWork()
		case home.Meal <= c.MinMealBorder:
			women.BuyProducts()
		case men.Human.Happiness < c.MinHappinessBorder:
			men.PlayComputer()
		case women.Human.Happiness < c.MinHappinessBorder:
			women.BuyCoat()
		}
		if h.IsDead(&aleks) || h.IsDead(&nastya) {
			break
		}
		home.Grime += c.GrimIncrement
	}

	home.ShowStatistic()
}