package human

import "fmt"

type Human struct {
	IsAlive             bool
	Satiety             int32
	Happiness           int32
	maxFoodForEat       int32
	wastedSatietyForDay int32
	Name                string
}

// CalculateHappinessLevel Начисляем кол-во очков исходя из кол-ва грязи, при недостаточном кол-во очков счастья человек умирает
func (human *Human) CalculateHappinessLevel(dirtPoint int32) {
	if dirtPoint > 90 {
		human.Happiness -= 10
	}
	if human.Happiness < 10 {
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

func (human *Human) wasteSatiety() {
	human.Satiety -= human.wastedSatietyForDay
	if human.Satiety < 0 {
		fmt.Print(human.Name, ": Dead, reason satiety")
		human.IsAlive = false
	}
}

func Factory(name string) Human {
	return Human{
		IsAlive:             true,
		Satiety:             30,
		Happiness:           100,
		maxFoodForEat:       30,
		wastedSatietyForDay: 10,
		Name:                name,
	}
}
