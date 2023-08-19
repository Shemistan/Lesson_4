package women

import (
	"fmt"
	"github.com/Shemistan/Lesson_4/home"
)

const PriceCoat = 350
const CountGrime = 90

type Women struct {
	home.Human
}

// Жена покупает продукты на 60 единиц, если денег в доме больше на 60.
// Если дома количество денег меньше 60 то она покупает продукты ровно столько есть дома денег

func (w *Women) BuyProducts() {
	if w.Human.Home.Money >= 60 {
		w.Human.Home.Meal += 60
		w.Human.Home.Money -= 60
		fmt.Println(w.Human.Name, "Купила продукты для дома", 60)
	}
	if w.Human.Home.Money < 60 {
		w.Human.Home.Meal += w.Human.Home.Money
		fmt.Println(w.Human.Name, "Купила продукты для дома", w.Human.Home.Money)
		w.Human.Home.Money -= w.Human.Home.Money
	}
	fmt.Println(w.Human.Name, "Не может купить продукты так как не хватае денег")
	w.Human.Fullness -= 10

}

// Жена чистить дом если грязи нет то дом чистый

func (w *Women) CleanHome() {
	if w.Human.Home.Grime > CountGrime {
		w.Human.Home.Grime -= CountGrime
		w.Human.Fullness -= 10
	} else {
		w.Human.Fullness -= 10
		w.Human.Home.Grime = 0
	}
	fmt.Println(w.Human.Name, "Закончила уборку дома")
}

// Жена покупает себе шубу, если количество денег мало то показывем что не хватает денег

func (w *Women) BuyCoat() {
	if w.Human.Home.Money >= PriceCoat {
		w.Human.Home.Money -= PriceCoat
		w.Human.Happiness += 60
		w.Human.Fullness -= 10
		w.Human.Home.CountCoat += 1
		fmt.Println(w.Human.Name, "Купила себе шубу")

	} else {
		fmt.Println("Вы не сможете купить шубу, у вас дома", w.Human.Home.Money, "денег")
	}
}
