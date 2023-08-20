package person

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/house"
	"github.com/Shemistan/Lesson_4/utils"
)

type Person struct {
	Name      string
	Fullness  int64
	Happiness int64
	House     *house.House
}

// Eat Person употребляет 30 едениц продуктов если в доме больше 30 едениц еды. Иначе, употребляет продкутов равное
// количеству доступной еды в доме
func (p *Person) Eat() int64 {
	switch {
	case p.House.Meal > utils.MealThreshold:
		p.Fullness += 30
		p.House.Meal -= 30
		fmt.Printf("%v had a meal\n", p.Name)
		return 30
	case p.House.Meal < utils.MealThreshold:
		mealEaten := p.House.Meal
		p.Fullness += p.House.Meal
		p.House.Meal -= mealEaten
		fmt.Printf("%v had a meal\n", p.Name)
		return mealEaten
	default:
		fmt.Printf("%v couldn't eat because the is no food", p.Name)
		return 0
	}
}
