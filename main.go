package main

import "fmt"

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

	home := Home{
		money:  100,
		fridge: 50,
		dirtPoints: DirtPoints{
			current:     0,
			incremental: 5,
		},
		wife: Wife{
			foodToMoneyRatio: 1,
			cleanDirtPoints:  100,
			coat: CoatSpecs{
				HappinessPoints: 60,
				Price:           350,
			},
			Human: Human{
				IsAlive:             true,
				satiety:             30,
				happiness:           100,
				maxFoodForEat:       30,
				wastedSatietyForDay: 10,
				Name:                "Mary",
			},
		},
		husband: Husband{
			earningsMoney:       150,
			happinessForPlaying: 20,
			Human: Human{
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

func year(home *Home, stats *Stats) {
	for i := 0; i < 365; i++ {
		if home.husband.IsAlive == false {
			fmt.Print("husband dead you loose")
			return
		}
	}
}
