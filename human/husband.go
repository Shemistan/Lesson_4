package human

type Husband struct {
	Human
	earningsMoney       int32
	happinessForPlaying int32
}

func (husband *Husband) PlayComputer() {
	husband.Human.wasteSatiety()
	husband.Human.increaseHappiness(husband.happinessForPlaying)
}

// EarnMoney Вовзращает кол-во заработаных денег
func (husband *Husband) EarnMoney() int32 {
	husband.Human.wasteSatiety()
	return husband.earningsMoney
}

func HusbandFactory(name string) Husband {
	return Husband{
		earningsMoney:       150,
		happinessForPlaying: 20,
		Human:               Factory(name),
	}
}
