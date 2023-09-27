package main

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/models"
	"github.com/Shemistan/Lesson_4/service"
)

const (
	startMoneyInNightstand = 100
	startFoodInFridge      = 50
	startDirtyInHouse      = 0
	startSatiety           = 30
	startHappiness         = 100

	startCatSatiety = 30
	startCatFood    = 30

	husbandName = "Sam"
	wifeName    = "Alice"
	catName     = "Kitty"
)

func main() {
	fmt.Println("Создаем семью и смотрим итоги прожитого года")

	husband := models.Person{
		Creature: models.Creature{
			Name:    husbandName,
			Satiety: startSatiety,
		},
		Happiness: startHappiness,
	}

	wife := models.Person{
		Creature: models.Creature{
			Name:    wifeName,
			Satiety: startSatiety,
		},
		Happiness: startHappiness,
	}

	home := models.House{
		MoneyInNightstand: startMoneyInNightstand,
		FoodInFridge:      startFoodInFridge,
		DirtyInHouse:      startDirtyInHouse,
		CatFood:           startCatFood,
	}

	cat := models.Creature{
		Name:    catName,
		Satiety: startCatSatiety,
	}

	family := &models.Family{
		Husband: husband,
		Wife:    wife,
		Home:    home,
		Cat:     cat,
	}

	result := service.TryNotToDie(family)

	fmt.Println("\n\nИтоги года:")
	fmt.Printf("Заработано денег за год: %d	единиц\n", result.EarnedMoney)
	fmt.Printf("Съедено еды за год: %d единиц\n", result.EatenFood)
	fmt.Printf("Куплено шуб за год: %d единиц\n", result.BoughtFur)
	fmt.Printf("Количество погибших за год: %d\n", len(result.DeadPerson))

	if len(result.DeadPerson) > 0 {
		fmt.Println("Список погибших: ")
		for dead := range result.DeadPerson {
			fmt.Println(dead)
		}

		fmt.Println("Задание не выполнено.")
	} else {
		fmt.Println("Все выжили, задание выполнено.")
	}

}
