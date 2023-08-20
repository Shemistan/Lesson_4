package wife

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/person"
	"github.com/Shemistan/Lesson_4/utils"
)

type Wife struct {
	Person *person.Person
}

// Shop Wife покупает 60 едениц продуктов если количество денег в доме превышает 60 монет.
// Иначе покупает количество еды равное количеству денег в доме
func (w *Wife) Shop() {
	switch {
	case w.Person.House.Money >= utils.MoneyForGroceries:
		w.Person.Fullness -= utils.FullnessToSubtract
		w.Person.House.Money -= utils.MoneyForGroceries
		w.Person.House.Meal += utils.MealUnits
		fmt.Printf("%v bought %v groceries\n", w.Person.Name, 60)
	case w.Person.House.Money < utils.MoneyForGroceries:
		w.Person.Fullness -= utils.FullnessToSubtract
		w.Person.House.Meal += w.Person.House.Money
		fmt.Printf("%v bought %v groceries\n", w.Person.Name, w.Person.House.Money)
		w.Person.House.Money -= w.Person.House.Money
	default:
		fmt.Printf("%v couldn'y buy croceries. Not enough money\n", w.Person.Name)
	}
}

// BuyCoat Wife проверяет достаточно ли денег в доме. Если да, то покупает шубу и возвращает true, иначе false
func (w *Wife) BuyCoat() bool {
	switch {
	case w.Person.House.Money < utils.MoneyForCoat:
		fmt.Printf("Coat costs %v, and you have %v\n", utils.MoneyForCoat, w.Person.House.Money)
		return false
	default:
		w.Person.Fullness -= utils.FullnessToSubtract
		w.Person.House.Money -= utils.MoneyForCoat
		w.Person.Happiness += utils.WifeHappinessUnits
		fmt.Printf("%v bought a coat\n", w.Person.Name)
		return true
	}
}

// Clean Wife проверяет насколко загрязнен дом. Если уровень загрязнения больше или равно 90, то жена очищает 60
// едениц грязи. Если уровень заргязнения меньше 60 едениц, то она очищает ровно столько насколько зарязнен дом.
func (w *Wife) Clean() {
	switch {
	case w.Person.House.Grime >= utils.GrimeThreshold:
		w.Person.Fullness -= utils.FullnessToSubtract
		w.Person.House.Grime -= utils.GrimeToSubtract
	case w.Person.House.Grime < utils.GrimeThreshold:
		w.Person.Fullness -= utils.FullnessToSubtract
		w.Person.House.Grime -= w.Person.House.Grime
	}
	fmt.Printf("%v cleaned the place\n", w.Person.Name)
}
