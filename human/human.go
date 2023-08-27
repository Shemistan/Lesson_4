package human

import (
	"fmt"
	"math/rand"

	"github.com/kamolisrailov/Lesson_4.git/home"
)

const (
	MIN = 1
)

type Human struct {
	Name      string
	Fullnes   int16
	Happiness int16
}

func (human *Human) Eat(humanName string, home *home.Home) bool {
	human.Fullnes += 1
	food := get_food(home.Food)
	home.Food -= food
	fmt.Printf("food to eat: %d\n", food)
	fmt.Printf("Food has eaten. %s Fullness: %d\n", humanName, human.Fullnes)
	fmt.Printf("Food at home: %d\n", home.Food)
	return true

}

func get_food(food int) int {

	amount := rand.Intn(food)
	if amount == 0 {
		get_food(food)
	}
	return amount
}
