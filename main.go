package main

import (
	"fmt"

	f "github.com/Shemistan/Lesson_4/family"
	hs "github.com/Shemistan/Lesson_4/house"
	h "github.com/Shemistan/Lesson_4/husband"
	"github.com/Shemistan/Lesson_4/person"
	w "github.com/Shemistan/Lesson_4/wife"
)

func main() {
	day := 1

	house := hs.House{
		Money: 100,
		Meal:  50,
		Grime: 0,
	}

	john := person.Person{
		Name:      "John",
		Fullness:  30,
		Happiness: 100,
		House:     &house,
	}

	helen := person.Person{
		Name:      "Helen",
		Fullness:  30,
		Happiness: 100,
		House:     &house,
	}

	husband := h.Husband{
		Person: &john,
	}

	wife := w.Wife{
		Person: &helen,
	}

	family := f.Family{
		Husband: &husband,
		Wife:    &wife,
		House:   &house,
		Stats:   map[string]int64{"mealEaten": 0, "coatsBought": 0, "revenue": 0},
	}

	for day < 366 {
		fmt.Printf("--------\nDay - %d\n--------\n", day)

		if isDead(&john) || isDead(&helen) {
			fmt.Printf("Total revenue - %v\nMeal eaten - %v\nCoats bought - %v\n", family.Stats["revenue"], family.Stats["mealEaten"], family.Stats["coatsBought"])
			break
		}

		switch {
		case family.Husband.Person.Fullness < 20:
			family.Eat(&john)
		case family.Wife.Person.Fullness < 20:
			family.Eat(&helen)
		case family.House.Grime > 90:
			family.Husband.Person.Happiness -= 10
			family.Wife.Person.Happiness -= 10
			family.Clean()
		case family.House.Meal < 30:
			family.BuyGroceries()
		case family.House.Money < 200:
			family.Work()
		case family.Husband.Person.Happiness < 15:
			family.PlayGames()
		case family.Wife.Person.Happiness < 15:
			family.BuyCoat()
		}
		house.Grime += 5
		day++
	}
	fmt.Printf("Total revenue - %v\nMeal eaten - %v\nCoats bought - %v\n", family.Stats["revenue"], family.Stats["mealEaten"], family.Stats["coatsBought"])
}

func isDead(person *person.Person) bool {
	//fmt.Printf("Person - %v\nFullness - %v\nHappiness - %v\n", person.Name, person.Fullness, person.Happiness)
	switch {
	case person.Fullness <= 0 || person.Happiness < 10:
		fmt.Printf("%v is dead\n", person.Name)
		return true
	default:
		return false
	}
}
