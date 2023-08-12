package main

type Husband struct {
	Human
	earningsMoney       int32
	happinessForPlaying int32
}

func (husband *Husband) playComputer() {
	husband.Human.wasteSatiety()
	husband.Human.increaseHappiness(husband.happinessForPlaying)
}

// Вовзращает кол-во заработаных денег
func (husband *Husband) earnMoney() int32 {
	husband.Human.wasteSatiety()
	return husband.earningsMoney
}
