package men

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/constants"
	"github.com/Shemistan/Lesson_4/home"
)

type Men struct {
	home.Human
}

func (m *Men) PlayComputer() {
	m.Human.Happiness += constants.IncreaseHappinessMen
	m.Human.Fullness -= constants.DecreaseHappiness
	fmt.Print(m.Human.Name, " Играл компьютерные игры \n")
}

func (m *Men) GoWork() {
	m.Human.Home.Money += constants.EarnedMoneyByMen
	m.Human.Fullness -= constants.DecreaseHappiness
	m.Human.Home.CountMoney += constants.EarnedMoneyByMen
	fmt.Print(m.Human.Name, " Ходил на работу \n")
}
