package main

import (
	"fmt"
	"math/rand"
	_ "math/rand"
)

type House struct {
	cleanliness int
	money       int
	food        int
	moneystat   int // Лень было выделять в отдельную структуру, поэтому сложил в дом, вроде тут тоже по логике может лежать
	foodstat    int
	shubstat    int
}

func PrintStat(man *Male, woman *Female, house *House) {
	fmt.Println(
		man.name, "сытость:", man.fullness, "счастье:", man.happiness, "--------",
		woman.name, "сытость:", woman.fullness, "счастье:", woman.happiness,
		"-------- HOUSE", "еда:", house.food, "деньги:", house.money, "чистота:", house.cleanliness)
}

func IsDead(man *Male, woman *Female) bool {
	if man.fullness <= 0 || woman.fullness <= 0 {
		fmt.Println("Смерть от голода")
		return true
	} else if man.happiness <= 0 || woman.happiness <= 0 {
		fmt.Println("Смерть без счастья")
		return true
	} else {
		return false
	}
}

// проверяем условия близкие к вымиранию - изначально потребности закрываем
func DoBeforeDeadWoman(woman *Female, house *House) bool {
	isDo:=false
	switch {
	case house.food <= dangerHouseFood:
		woman.buyEat(house)
		isDo = true
	case woman.fullness <= dangerFullness:
		woman.eat(house)
		isDo = true
	case woman.happiness <= dangerHappiness:
		woman.buyCoat(house)
		isDo = true
	case house.cleanliness >= dangerCleanliness:
		woman.clean(house)
		isDo = true
	default:
		isDo = false
	}
	return isDo
}

func DoBeforeDeadMan(man *Male, woman *Female, house *House) bool {
	isDo:=false
	switch {
	case man.fullness <= dangerFullness:
		man.eat(house)
		isDo = true
	case man.happiness <= dangerHappiness:
		man.play()
		isDo = true
	case house.money <= 100:
		man.earn(house)
		isDo = true
	case woman.happiness <= dangerHappiness:
		man.earn(house)
		isDo = true
	default:
		isDo = false
	}
	return isDo
}

func DoInDay(man *Male, woman *Female, house *House) {
	isManDo := DoBeforeDeadMan(man, woman, house)
	isWomanDo := DoBeforeDeadWoman(woman, house)

	//рандомный таск, если не было Danger Man
	randomTaskMan := 1 + rand.Intn(3)
	switch {
	case isManDo == true:
		fmt.Println("Сработал Danger")
	case randomTaskMan == 1:
		man.eat(house)
	case randomTaskMan == 2:
		man.earn(house)
	case randomTaskMan == 3:
		man.play()
	}

	//рандомный таск, если не было Danger Woman
	randomTaskWoman := 1 + rand.Intn(3)
	switch {
	case isWomanDo == true:
		fmt.Println("Сработал Danger")
	case randomTaskWoman == 1:
		woman.eat(house)
	case randomTaskWoman == 2:
		woman.buyCoat(house)
	case randomTaskWoman == 3:
		woman.buyEat(house)
	}

	house.cleanliness += 5
}
