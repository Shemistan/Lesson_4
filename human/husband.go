package human

type Husband struct {
	Human
	earningsMoney       int32
	happinessForPlaying int32
}

const (
	salary              int32 = 150
	happinessForPlaying int32 = 20
)

func (husband *Husband) PlayComputer() error {
	err := husband.Human.wasteSatiety()

	if err != nil {
		return err
	}

	husband.Human.increaseHappiness(husband.happinessForPlaying)

	return nil
}

// EarnMoney Вовзращает кол-во заработаных денег
func (husband *Husband) EarnMoney() (err error, earnedMoney int32) {
	err = husband.Human.wasteSatiety()

	if err != nil {
		return err, 0
	}

	return err, husband.earningsMoney
}

func CreateHusband(name string) Husband {
	return Husband{
		earningsMoney:       salary,
		happinessForPlaying: happinessForPlaying,
		Human:               CreateHuman(name),
	}
}
