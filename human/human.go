package human

import "fmt"

const (
	unhappyFromDirtEdge             int32 = 90
	unhappyEdge                     int32 = 20
	decreaseHappinessFromDirtPerDay int32 = 10
	deadFromUnhappyEdge             int32 = 10
	maxFoodForEat                   int32 = 30
	wastedSatietyForDay             int32 = 10

	initialIsAlive   bool  = true
	initialSatiety   int32 = 30
	initialHappiness int32 = 100
)

type Human struct {
	IsAlive             bool
	Satiety             int32
	Happiness           int32
	maxFoodForEat       int32
	wastedSatietyForDay int32
	Name                string
}

// CalculateHappinessForToday Начисляем кол-во очков исходя из кол-ва грязи, при недостаточном кол-во очков счастья человек умирает
func (human *Human) CalculateHappinessForToday(dirtPoint int32) {
	if dirtPoint > unhappyFromDirtEdge {
		human.Happiness -= decreaseHappinessFromDirtPerDay
	}
	if human.Happiness < deadFromUnhappyEdge {
		fmt.Print(human.Name, ": Dead, reason happiness")
		human.IsAlive = false
	}
}

func (human *Human) increaseHappiness(happinessPoints int32) {
	human.Happiness += happinessPoints
}

func (human *Human) Eat(foodPoints int32) {
	if foodPoints <= human.maxFoodForEat {
		human.Satiety += foodPoints
		return
	}
	human.Satiety += human.maxFoodForEat
}

func (human *Human) IsFamilyMemberUnhappy() bool {
	return human.Happiness <= unhappyEdge
}

func (human *Human) wasteSatiety() {
	human.Satiety -= human.wastedSatietyForDay
	if human.Satiety < 0 {
		fmt.Print(human.Name, ": Dead, reason satiety")
		human.IsAlive = false
	}
}

func CreateHuman(name string) Human {
	return Human{
		IsAlive:             initialIsAlive,
		Satiety:             initialSatiety,
		Happiness:           initialHappiness,
		maxFoodForEat:       maxFoodForEat,
		wastedSatietyForDay: wastedSatietyForDay,
		Name:                name,
	}
}
