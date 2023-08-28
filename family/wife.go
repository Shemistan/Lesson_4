package family

import (
	"fmt"

	"github.com/kamolisrailov/Lesson_4.git/home"
	"github.com/kamolisrailov/Lesson_4.git/human"
)

const (
	Food = 60
	Dust = 100
)

type Wife struct {
	human.Human
}

func (h *Wife) BuyGrocery(home *home.Home) {
	h.Fullnes -= Fullness
	home.Food += Food
	home.Money -= Food
	fmt.Printf("%s went to the market and bought %d food\n", h.Name, Food)
	fmt.Printf("%s's fullness: %d\n", h.Name, h.Fullnes)
}

func (h *Wife) CleanHouse(home *home.Home) {
	h.Fullnes -= Fullness
	home.Dust -= Dust
	fmt.Printf("%s has cleaned house\n", h.Name)
	fmt.Printf("House dirtiness is: : %d\n", home.Dust)
}
