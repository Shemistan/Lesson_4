package human

import (
	"errors"
	"fmt"
)

const (
	unhappyFromDirtEdge             int32 = 90
	unhappyEdge                     int32 = 20
	decreaseHappinessFromDirtPerDay int32 = 10
	deadFromUnhappyEdge             int32 = 10
	deadFromSatietyEdge             int32 = 0
	maxFoodForEat                   int32 = 30
	wastedSatietyForDay             int32 = 10

	initialIsAlive   bool  = true
	initialSatiety   int32 = 30
	initialHappiness int32 = 100

	happinessDeadMessage = "%s: dead from happiness"
	satietyDeadMessage   = "%s: dead from satiety"
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
func (human *Human) CalculateHappinessForToday(dirtPoint int32) error {
	if dirtPoint > unhappyFromDirtEdge {
		human.Happiness -= decreaseHappinessFromDirtPerDay
	}
	if human.Happiness < deadFromUnhappyEdge {
		human.IsAlive = false
		errorMessage := fmt.Sprintf(happinessDeadMessage, human.Name)
		return errors.New(errorMessage)
	}
	return nil
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

func (human *Human) increaseHappiness(happinessPoints int32) {
	human.Happiness += happinessPoints
}

func (human *Human) wasteSatiety() error {
	human.Satiety -= human.wastedSatietyForDay
	if human.Satiety < deadFromSatietyEdge {
		human.IsAlive = false
		errorMessage := fmt.Sprintf(satietyDeadMessage, human.Name)
		return errors.New(errorMessage)
	}

	return nil
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
