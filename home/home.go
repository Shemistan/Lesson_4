package home

import (
	"fmt"
	"github.com/Shemistan/Lesson_4/human"
)

type DirtPoints struct {
	current     int32
	incremental int32
}

type Home struct {
	money      int32
	fridge     int32
	dirtPoints DirtPoints
	wife       human.Wife
	husband    human.Husband
}

func (home *Home) nextDay() {
	home.husband.CalculateHappinessLevel(home.dirtPoints.current)
	home.wife.CalculateHappinessLevel(home.dirtPoints.current)
	home.dirtPoints.current += home.dirtPoints.incremental
}

func (home *Home) cleanHome() {
	home.dirtPoints.current = home.wife.CleanUp(home.dirtPoints.current)
}

func (home *Home) earnMoney() {
	home.money += home.husband.EarnMoney()
}

func (home *Home) eatFromFridge(human *human.Human, foodPoint int32) {
	if foodPoint > home.fridge {
		fmt.Print("Not enough food in Fridge")
	}
	home.fridge -= foodPoint
	human.Eat(foodPoint)
}
