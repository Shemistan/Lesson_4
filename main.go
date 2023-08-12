package main

import (
	"fmt"
	home2 "github.com/Shemistan/Lesson_4/home"
	"github.com/Shemistan/Lesson_4/human"
)

type Stats struct {
	EarnedMoney int64
	EatenFood   int64
	BoughtCoats int32
}

func main() {
	stats := Stats{
		EarnedMoney: 0,
		EatenFood:   0,
		BoughtCoats: 0,
	}

	home := home2.Home{
		money:  100,
		fridge: 50,
		dirtPoints: home2.DirtPoints{
			current:     0,
			incremental: 5,
		},
		wife: human.Wife{
			foodToMoneyRatio: 1,
			cleanDirtPoints:  100,
			coat: human.CoatSpecs{
				HappinessPoints: 60,
				Price:           350,
			},
			Human: human.Human{
				IsAlive:             true,
				satiety:             30,
				happiness:           100,
				maxFoodForEat:       30,
				wastedSatietyForDay: 10,
				Name:                "Mary",
			},
		},
		husband: human.Husband{
			earningsMoney:       150,
			happinessForPlaying: 20,
			Human: human.Human{
				IsAlive:             true,
				satiety:             30,
				happiness:           100,
				maxFoodForEat:       30,
				wastedSatietyForDay: 10,
				Name:                "John",
			},
		},
	}
}

func year(home *home2.Home, stats *Stats) {
	for i := 0; i < 365; i++ {
		if home.husband.IsAlive == false {
			fmt.Print("husband dead you loose")
			return
		}
	}
}
