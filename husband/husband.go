package husband

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/person"
	"github.com/Shemistan/Lesson_4/utils"
)

type Husband struct {
	Person *person.Person
}

func (h *Husband) PlayGames() {
	h.Person.Fullness -= utils.FullnessToSubtract
	h.Person.Happiness += utils.HusbandHappinessUnits
	fmt.Printf("%v played video games\n", h.Person.Name)
}

func (h *Husband) Work() int64 {
	h.Person.Fullness -= utils.FullnessToSubtract
	h.Person.House.Money += utils.EarnedMoney
	fmt.Printf("%v went to work\n", h.Person.Name)
	return utils.EarnedMoney
}
