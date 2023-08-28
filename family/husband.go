package family

import (
	"fmt"

	"github.com/kamolisrailov/Lesson_4.git/home"
	"github.com/kamolisrailov/Lesson_4.git/human"
)

const (
	Money     = 150
	Fullness  = 10
	Happiness = 20
)

type Husband struct {
	human.Human
}

func (h *Husband) GoToWork(home *home.Home) {
	h.Fullnes -= Fullness
	home.Money += Money
	fmt.Printf("%s went to the work and earned %d money\n", h.Name, Money)
	fmt.Printf("%s's fullness: %d\n", h.Name, h.Fullnes)
}

func (h *Husband) PlayComputer() {
	h.Fullnes -= Fullness
	h.Happiness += Happiness
	fmt.Printf("%s has played computer and his Happiness has increased to: %d\n", h.Name, Happiness)
}
