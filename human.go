package main

type Human struct {
	isAlive             bool
	satiety             int32
	happiness           int32
	maxFoodForEat       int32
	wastedSatietyForDay int32
	name                string
}

// Начисляем кол-во очков исходя из кол-ва грязи, при недостаточном кол-во очков счастья человек умирает
func (human *Human) calculateHappinessLevel(dirtPoint int32) {
	if dirtPoint > 90 {
		human.happiness -= 10
	}
	if human.happiness < 10 {
		human.isAlive = false
	}
}

func (human *Human) increaseHappiness(happinessPoints int32) {
	human.happiness += happinessPoints
}

func (human *Human) eat(foodPoints int32) {
	if foodPoints <= human.maxFoodForEat {
		human.satiety += foodPoints
		return
	}
	human.satiety += human.maxFoodForEat
}

func (human *Human) wasteSatiety() {
	human.satiety -= human.wastedSatietyForDay
	if human.satiety < 0 {
		human.isAlive = false
	}
}
