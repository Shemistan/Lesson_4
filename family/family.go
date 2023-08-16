package family

import (
	"github.com/Shemistan/Lesson_4/house"
	"github.com/Shemistan/Lesson_4/husband"
	"github.com/Shemistan/Lesson_4/person"
	"github.com/Shemistan/Lesson_4/wife"
)

type Family struct {
	Husband *husband.Husband
	Wife    *wife.Wife
	House   *house.House
	Stats   map[string]int64
}

func (f *Family) Eat(person *person.Person) {
	mealEaten := person.Eat()
	f.Stats["mealEaten"] += mealEaten
}

func (f *Family) BuyCoat() {
	isBought := f.Wife.BuyCoat()
	switch {
	case isBought:
		f.Stats["coatsBought"] += 1
	default:
		f.Stats["coatsBought"] += 0
	}
}

func (f *Family) Work() {
	revenue := f.Husband.Work()
	f.Stats["revenue"] += revenue
}

func (f *Family) PlayGames() {
	f.Husband.PlayGames()
}

func (f *Family) BuyGroceries() {
	f.Wife.Shop()
}

func (f *Family) Clean() {
	f.Wife.Clean()
}
