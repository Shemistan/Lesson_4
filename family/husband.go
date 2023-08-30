package family

import (
	"fmt"

	"github.com/kamolisrailov/Lesson_4.git/home"
	"github.com/kamolisrailov/Lesson_4.git/human"
)

const (
	Money            = 150
	Fullness         = 10
	HusbandHappiness = 20
	HappinessMin     = 10
	FoodMin          = 60
	MoneyMin         = 100
)

type Husband struct {
	human.Human
	CountPlayComputer int
}

func (h *Husband) GoToWork(home *home.Home) {
	h.Fullnes -= Fullness
	home.Money += Money
	fmt.Printf("%s went to the work and earned %d money\n", h.Name, Money)
	fmt.Printf("%s's fullness: %d\n", h.Name, h.Fullnes)
}

func (h *Husband) PlayComputer() {
	h.Fullnes -= Fullness
	h.Happiness += HusbandHappiness
	h.CountPlayComputer += 1
	fmt.Printf("%s has played computer and his Happiness has increased to: %d\n", h.Name, HusbandHappiness)
}

func (h *Husband) DailyActions(home *home.Home) int {
	if home.Dust > 90 {
		h.Happiness -= 10
	}
	if h.Fullnes-Fullness < 0 {
		h.Eat(home)
		return 0
	} else if h.Happiness <= 10 {
		h.PlayComputer()
		return 0
	} else if home.Food < FoodMin || home.Money < MoneyMin {
		h.GoToWork(home)
		return 0
	} else {
		h.GoToWork(home)
		return 0
	}

}
