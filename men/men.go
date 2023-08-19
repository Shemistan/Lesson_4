package men

import (
	"fmt"
	"github.com/Shemistan/Lesson_4/home"
)

const SALARY = 150

type Men struct {
	home.Human
}

func (m *Men) PlayComputer() {
	m.Human.Happiness += 20
	m.Human.Fullness -= 10
	fmt.Print(m.Human.Name, " Играл компьютерные игры \n")
}

func (m *Men) GoWork() {
	m.Human.Home.Money += SALARY
	m.Human.Fullness -= 10
	m.Human.Home.CountMoney += SALARY
	fmt.Print(m.Human.Name, " Ходил на работу \n")
}
