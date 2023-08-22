package run

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/family"
	"github.com/Shemistan/Lesson_4/house"
	"github.com/Shemistan/Lesson_4/husband"
	"github.com/Shemistan/Lesson_4/person"
	"github.com/Shemistan/Lesson_4/utils"
	"github.com/Shemistan/Lesson_4/wife"
)

func Run() {
	var day int64 = 1

	houseObject := house.House{
		Money: utils.InitialMoney,
		Meal:  utils.InitialMeal,
		Grime: utils.InitialGrime,
	}

	john := person.Person{
		Name:      "John",
		Fullness:  utils.InitialFullness,
		Happiness: utils.InitialHappiness,
		House:     &houseObject,
	}

	helen := person.Person{
		Name:      "Helen",
		Fullness:  utils.InitialFullness,
		Happiness: utils.InitialHappiness,
		House:     &houseObject,
	}

	husbandObject := husband.Husband{
		Person: &john,
	}

	wifeObject := wife.Wife{
		Person: &helen,
	}

	familyObject := family.Family{
		Husband: &husbandObject,
		Wife:    &wifeObject,
		House:   &houseObject,
		Stats:   map[string]int64{"mealEaten": 0, "coatsBought": 0, "revenue": 0},
	}

	for day < utils.Days {
		fmt.Printf("--------\nDay - %d\n--------\n", day)

		err := isDead(&john)

		if err != nil {
			fmt.Printf("Total revenue - %v\nMeal eaten - %v\nCoats bought - %v\n", familyObject.Stats["revenue"], familyObject.Stats["mealEaten"], familyObject.Stats["coatsBought"])
			break
		}

		err = isDead(&helen)

		if err != nil {
			fmt.Printf("Total revenue - %v\nMeal eaten - %v\nCoats bought - %v\n", familyObject.Stats["revenue"], familyObject.Stats["mealEaten"], familyObject.Stats["coatsBought"])
			break
		}

		switch {
		case familyObject.Husband.Person.Fullness < utils.FullnessThreshold:
			familyObject.Eat(&john)
		case familyObject.Wife.Person.Fullness < utils.FullnessThreshold:
			familyObject.Eat(&helen)
		case familyObject.House.Grime > utils.GrimeThreshold:
			familyObject.Husband.Person.Happiness -= utils.HappinessToSubtract
			familyObject.Wife.Person.Happiness -= utils.HappinessToSubtract
			familyObject.Clean()
		case familyObject.House.Meal < utils.MealThreshold:
			familyObject.BuyGroceries()
		case familyObject.House.Money < utils.MoneyThreshold:
			familyObject.Work()
		case familyObject.Husband.Person.Happiness < utils.HappinessThreshold:
			familyObject.PlayGames()
		case familyObject.Wife.Person.Happiness < utils.HappinessThreshold:
			familyObject.BuyCoat()
		}
		houseObject.Grime += utils.GrimeUnits
		day++
	}
	fmt.Printf("Total revenue - %v\nMeal eaten - %v\nCoats bought - %v\n", familyObject.Stats["revenue"], familyObject.Stats["mealEaten"], familyObject.Stats["coatsBought"])
}

func isDead(person *person.Person) error {
	switch {
	case person.Fullness <= utils.LowFullness || person.Happiness < utils.LowHappiness:
		return fmt.Errorf("person is dead...\nName: %v\n", person.Name)
	default:
		return nil
	}
}
