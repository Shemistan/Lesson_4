package main

import "fmt"

func main() {
	//var a Animal
	house := House{
		cleanliness: 0,
		money:       100,
		food:        50,
		moneystat:   0,
		foodstat:    0,
		shubstat:    0,
	}
	man := Male{
		People{
			name:      "Pavel",
			fullness:  30,
			happiness: 100,
		},
	}
	woman := Female{
		People{
			name:      "Katya",
			fullness:  30,
			happiness: 100,
		},
	}
	for i := 1; i <= 365; i++ {
		fmt.Println("----------")
		fmt.Println("----------Пошел день", i)
		Day(&man, &woman, &house)
		PrintStat(&man, &woman, &house)
		if IsDead(&man, &woman) == true {
			break
		}
	}
	fmt.Println("Статистика за игру: Денег заработано:", house.moneystat,
		", Шуб куплено:", house.shubstat, ", еды съедено:", house.foodstat)
}
