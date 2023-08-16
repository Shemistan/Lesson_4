package husband

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/person"
)

type Husband struct {
	Person *person.Person
}

func (h *Husband) PlayGames() {
	h.Person.Fullness -= 10
	h.Person.Happiness += 20
	fmt.Printf("%v played video games\n", h.Person.Name)
}

func (h *Husband) Work() int64 {
	h.Person.Fullness -= 10
	h.Person.House.Money += 150
	fmt.Printf("%v went to work\n", h.Person.Name)
	return 150
}
