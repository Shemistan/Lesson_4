package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_4/models"
	"github.com/Shemistan/Lesson_4/service"
)

const (
	startMoney     = 100 // начальное кол-во капитала
	startHappiness = 100 // начальное кол-во радости
	startSatiety   = 30  // начальное кол-во сытости
	startFood      = 50  // начальное кол-во еды
	startDirty     = 0   // начальное кол-во грязи
)

func main() {
	fmt.Printf("\n\n   a year of family life: \n\n")

	family := &models.Family{
		Husband: models.Person{
			Name:      "Neo",
			Satiety:   startSatiety,
			Happiness: startHappiness,
		},
		Wife: models.Person{
			Name:      "Trinity",
			Satiety:   startSatiety,
			Happiness: startHappiness,
		},
		House: models.House{
			Money: startMoney,
			Food:  startFood,
			Dirty: startDirty,
		},
	}

	totally, err := service.YearOfLife(family)
	if err != nil {
		fmt.Println(err)
	} else {
		resultsOfTheYear(totally, family)
	}
}

// findResultYear подводим итоги года
func resultsOfTheYear(s models.Summary, f *models.Family) {
	fmt.Println(service.Separator)
	fmt.Println("Итоги года:")
	fmt.Printf("Заработано денег: %d\n", s.TotalMoney-startMoney)
	fmt.Printf("Съедено еды: %d\n", s.TotalFoodEaten)
	fmt.Printf("Куплено шуб: %d\n\n", s.TotalBuyCoats)
	fmt.Printf("Стата мужа: сытость - %v, счастье - %v\n", f.Husband.Satiety, f.Husband.Happiness)
	fmt.Printf("Стата жены: сытость - %v, счастье - %v\n", f.Wife.Satiety, f.Wife.Happiness)
}
