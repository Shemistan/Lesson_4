package women

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/constants"
	"github.com/Shemistan/Lesson_4/home"
)

type Women struct {
	home.Human
}

// Жена покупает продукты на 60 единиц, если денег в доме больше на 60.
// Если дома количество денег меньше 60 то она покупает продукты ровно столько есть дома денег

func (w *Women) BuyProducts() {
	if w.Human.Home.Money >= constants.CountMoneyForProducts {
		w.Human.Home.Meal += constants.CountMeal
		w.Human.Home.Money -= constants.CountMoneyForProducts
		fmt.Println(w.Human.Name, "Купила продукты для дома", constants.CountMoneyForProducts)
	}
	if w.Human.Home.Money < constants.CountMoneyForProducts {
		w.Human.Home.Meal += w.Human.Home.Money
		fmt.Println(w.Human.Name, "Купила продукты для дома", w.Human.Home.Money)
		w.Human.Home.Money -= w.Human.Home.Money
	}
	fmt.Println(w.Human.Name, "Не может купить продукты так как не хватае денег")
	w.Human.Fullness -= constants.DecreaseHappiness

}

// Жена чистить дом если грязи нет то дом чистый

func (w *Women) CleanHome() {
	if w.Human.Home.Grime > constants.MaxOfGrime {
		w.Human.Home.Grime -= constants.MaxOfGrime
		w.Human.Fullness -= constants.DecreaseHappiness
	} else {
		w.Human.Fullness -= constants.DecreaseHappiness
		w.Human.Home.Grime = constants.StartGrime
	}
	fmt.Println(w.Human.Name, "Закончила уборку дома")
}

// Жена покупает себе шубу, если количество денег мало то показывем что не хватает денег

func (w *Women) BuyCoat() {
	if w.Human.Home.Money >= constants.CoatPrice {
		w.Human.Home.Money -= constants.CoatPrice
		w.Human.Happiness += constants.IncreaseHappinessWomen
		w.Human.Fullness -= constants.DecreaseHappiness
		w.Human.Home.CountCoat += 1
		fmt.Println(w.Human.Name, "Купила себе шубу")

	} else {
		fmt.Println("Вы не сможете купить шубу, у вас дома", w.Human.Home.Money, "денег")
	}
}
