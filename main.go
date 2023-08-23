package main

import "fmt"

const(
//стартовые константы
startMoney=100
startFood=50
startFullness=30
startHappiness=100
scriptDays=365
//изменения значений
fullnessChange=30 //едят эту величину
moneyChange=150
happinessChangeMan=20
happinessChangeWoman=60
foodChange=60 //стандартная величина закупки еды
cleanChange=100 // максимальная очистка комнаты
coatPrice=350
fullnessDecrease=10 //сытость уменьшается на эту величину, если делают другие действия
//Danger values - условия для запуска действий, чтобы не умереть
dangerHouseFood=25
dangerFullness=20
dangerHappiness=40
dangerCleanliness=80
)

func main() {
	//var a Animal
	house := House{
		cleanliness: 0,
		money:       startMoney,
		food:        startFood,
		moneystat:   0,
		foodstat:    0,
		shubstat:    0,
	}
	man := Male{
		People{
			name:      "Pavel",
			fullness:  startFullness,
			happiness: startHappiness,
		},
	}
	woman := Female{
		People{
			name:      "Katya",
			fullness:  startFullness,
			happiness: startHappiness,
		},
	}
	for i := 1; i <= scriptDays; i++ {
		fmt.Println("----------")
		fmt.Println("----------Пошел день", i)
		DoInDay(&man, &woman, &house)
		PrintStat(&man, &woman, &house)
		if IsDead(&man, &woman) == true {
			break
		}
	}
	fmt.Println("Статистика за игру: Денег заработано:", house.moneystat,
		", Шуб куплено:", house.shubstat, ", еды съедено:", house.foodstat)
}
