package wife

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/person"
)

type Wife struct {
	Person *person.Person
}

// Shop Wife покупает 60 едениц продуктов если количество денег в доме превышает 60 монет.
// Иначе покупает количество еды равное количеству денег в доме
func (w *Wife) Shop() {
	switch {
	case w.Person.House.Money >= 60:
		w.Person.Fullness -= 10
		w.Person.House.Money -= 60
		w.Person.House.Meal += 60
		fmt.Printf("%v bought %v groceries\n", w.Person.Name, 60)
	case w.Person.House.Money < 60:
		w.Person.Fullness -= 10
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
	case w.Person.House.Money < 150:
		fmt.Printf("Coat costs %v, and you have %v\n", 150, w.Person.House.Money)
		return false
	default:
		w.Person.Fullness -= 10
		w.Person.House.Money -= 150
		w.Person.Happiness += 60
		fmt.Printf("%v bought a coat\n", w.Person.Name)
		return true
	}
}

// Clean Wife проверяет насколко загрязнен дом. Если уровень загрязнения больше или равно 90, то жена очищает 60
// едениц грязи. Если уровень заргязнения меньше 60 едениц, то она очищает ровно столько насколько зарязнен дом.
func (w *Wife) Clean() {
	switch {
	case w.Person.House.Grime >= 90:
		w.Person.Fullness -= 10
		w.Person.House.Grime -= 60
	case w.Person.House.Grime < 60:
		w.Person.Fullness -= 10
		w.Person.House.Grime -= w.Person.House.Grime
	}
	fmt.Printf("%v cleaned the place\n", w.Person.Name)
}
