package entities

import (
	"errors"
	"fmt"
)

type Family struct {
	husband   Husband
	wife      Wife
	cat       Cat
	resources Resources
	total     Statistics
}

func (f *Family) Init() {
	f.resources.Init(&f.total)
	f.husband.Init("Uasya", &f.resources)
	f.wife.Init("Uaselisa", &f.resources)
	f.cat.Init("Sarapka", &f.resources)
}

func (f *Family) RunSimulation() error {

	for i := 1; i <= SIMULATION_DURATION; i++ {

		fmt.Println("Day: ", i)
		f.resources.DirtIncrease(EVERY_DAY_DIRNESS_INCREASE)
		f.husband.HusbandSimulation()
		f.wife.WifeSimulation()
		f.cat.CatSimulation()

		if !f.husband.IsAlive || !f.wife.IsAlive || !f.cat.IsAlive {
			return errors.New("Someone is dead")
		}

		f.husband.PrintData()
		f.wife.PrintData()
		f.cat.PrintData()
		f.resources.PrintResources()

	}

	fmt.Println("All alive\nStatistics: \n")
	f.resources.Totals.PrintStatistics()

	return nil
}
