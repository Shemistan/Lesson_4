package human

import (
	"fmt"
	"math/rand"

	"github.com/kamolisrailov/Lesson_4.git/home"
)

const (
	MIN = 10
	MAX = 25
)

type Human struct {
	Name      string
	Fullnes   int
	Happiness int
}

func (human *Human) Eat(home *home.Home) bool {

	food := get_food(home.Food)
	home.Food -= food
	human.Fullnes += food
	fmt.Printf("food to eat: %d\n", food)
	fmt.Printf("Food has eaten. %s's Fullness: %d\n", human.Name, human.Fullnes)
	fmt.Printf("Food at home: %d\n", home.Food)
	return true

}

func get_food(food int) int {
	return rand.Intn(MAX-MIN) + MIN
}
