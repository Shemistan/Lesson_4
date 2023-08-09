package main

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/models"
	"github.com/Shemistan/Lesson_4/service"
)

func main() {
	fmt.Println("Год жизни не большой семьи")
	family := &models.Family{
		Husband: models.Person{
			Name:      "Chel",
			Satiety:   30,
			Happiness: 100,
		},
		Wife: models.Person{
			Name:      "Mary",
			Satiety:   30,
			Happiness: 100,
		},
		House: models.House{
			Money: 100,
			Food:  50,
			Dirty: 0,
		},
	}
	summary := service.YearOfLife(family)
	fmt.Println("===============================================================")
	fmt.Println("===============================================================")
	fmt.Println("Итоги года:")
	fmt.Printf("Заработано денег: %d\n", family.House.Money-100) // Минус начальные 100
	fmt.Printf("Съедено еды: %d\n", summary.TotalFoodEaten)
	fmt.Printf("Куплено шуб: %d\n", summary.TotalBuyCoats)
}
