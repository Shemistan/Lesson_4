package main

import "fmt"

type Stats struct {
	earnedMoney int64
	eatenFood   int64
	boughtCoats int32
}

func main() {
	stats := Stats{
		earnedMoney: 0,
		eatenFood:   0,
		boughtCoats: 0,
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
				isAlive:             true,
				satiety:             30,
				happiness:           100,
				maxFoodForEat:       30,
				wastedSatietyForDay: 10,
				name:                "Mary",
			},
		},
		husband: Husband{
			earningsMoney:       150,
			happinessForPlaying: 20,
			Human: Human{
				isAlive:             true,
				satiety:             30,
				happiness:           100,
				maxFoodForEat:       30,
				wastedSatietyForDay: 10,
				name:                "John",
			},
		},
	}
}

func year(home *Home, stats *Stats) {
	for i := 0; i < 365; i++ {
		if home.husband.isAlive == false {
			fmt.Print("husband dead you loose")
			return
		}
	}
}
