package main

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/models"
	"github.com/Shemistan/Lesson_4/service"
)

const (
	startMoney        = 100 // начальный капитал
	startHappiness    = 100 // начальная радость
	startSatietes     = 30  // начальная сытость
	startFoodSupplies = 50  // начальная еда
	startDirt         = 0   // начальная грязь
)

func main() {
	fmt.Println("Год жизни не большой семьи.")
	fmt.Println(service.Separator)

	family := &models.Family{
		Husband: models.Human{
			Name:      "Chel",
			Satiety:   startSatietes,
			Happiness: startHappiness,
		},
		Wife: models.Human{
			Name:      "Mary",
			Satiety:   startSatietes,
			Happiness: startHappiness,
		},
		House: models.House{
			Budget:       startMoney,
			FoodSupplies: startFoodSupplies,
			DirtInHouse:  startDirt,
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
func findResultYear(s models.Result) {
	fmt.Println(service.Separator)
	fmt.Println(service.Separator)
	fmt.Println("Итоги года:")
	fmt.Printf("Заработано денег: %d\n", s.TotalMoney-startMoney)
	fmt.Printf("Съедено еды: %d\n", s.TotalFoodEaten)
	fmt.Printf("Куплено шуб: %d\n", s.TotalBuyCoats)
}
