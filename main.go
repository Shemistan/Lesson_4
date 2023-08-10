package main

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/models"
	"github.com/Shemistan/Lesson_4/service"
)

const (
	START_MONEY    = 100 // начальный капитал
	START_HAPPINES = 100 // начальная радость
	START_SATIETES = 30  // начальная сытость
	START_FOOD     = 50  // начальная еда
	START_DIRTY    = 0   // начальная грязь
	SEPARATOR      = "==============================================================="
)

func main() {
	fmt.Println("Год жизни не большой семьи")
	fmt.Println(SEPARATOR)

	family := &models.Family{
		Husband: models.Person{
			Name:      "Chel",
			Satiety:   START_SATIETES,
			Happiness: START_HAPPINES,
		},
		Wife: models.Person{
			Name:      "Mary",
			Satiety:   START_SATIETES,
			Happiness: START_HAPPINES,
		},
		House: models.House{
			Money: START_MONEY,
			Food:  START_FOOD,
			Dirty: START_DIRTY,
		},
	}

	summary, err := service.YearOfLife(family)
	if err != nil {
		fmt.Println(err)
	} else {
		findResultYear(summary)
	}

}

// findResultYear подводим итоги года
func findResultYear(s models.Summary) {
	fmt.Println(SEPARATOR)
	fmt.Println(SEPARATOR)
	fmt.Println("Итоги года:")
	fmt.Printf("Заработано денег: %d\n", s.TotalMoney-START_MONEY)
	fmt.Printf("Съедено еды: %d\n", s.TotalFoodEaten)
	fmt.Printf("Куплено шуб: %d\n", s.TotalBuyCoats)
}
