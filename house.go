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
func DangerWomanTasks(woman *Female, house *House) bool {
	var isDo bool
	switch {
	case house.food <= 25:
		{
			woman.buyEat(house)
			isDo = true
		}
	case woman.fullness <= 20:
		{
			woman.eat(house)
			isDo = true
		}
	case woman.happiness < 41:
		{
			woman.buyCoat(house)
			isDo = true
		}
	case house.cleanliness >= 80:
		{
			woman.clean(house)
			isDo = true
		}
	default:
		isDo = false
	}
	return isDo
}

func DangerManTasks(man *Male, woman *Female, house *House) bool {
	var isDo bool
	switch {
	case man.fullness <= 20:
		{
			man.eat(house)
			isDo = true
		}
	case man.happiness < 21:
		{
			man.play()
			isDo = true
		}
	case house.money <= 100:
		{
			man.earn(house)
			isDo = true
		}
	case woman.happiness < 41:
		{
			man.earn(house)
			isDo = true
		}
	default:
		isDo = false
	}
	return isDo
}

func Day(man *Male, woman *Female, house *House) {
	isManDo := DangerManTasks(man, woman, house)
	isWomanDo := DangerWomanTasks(woman, house)
	i := 1 + rand.Intn(3)
	j := 1 + rand.Intn(3)
	switch {
	case isManDo == true:
		fmt.Println("Сработал Danger")
	case i == 1:
		man.eat(house)
	case i == 2:
		man.earn(house)
	case i == 3:
		man.play()
	}
	switch {
	case isWomanDo == true:
		fmt.Println("Сработал Danger")
	case j == 1:
		woman.eat(house)
	case j == 2:
		woman.buyCoat(house)
	case j == 3:
		woman.buyEat(house)
	}
	house.cleanliness += 5
}
