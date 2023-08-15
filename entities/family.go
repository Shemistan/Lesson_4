package entities

import "fmt"

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

func (f *Family) Logic() {

	for i := 1; i <= 365; i++ {

		fmt.Println("Day: ", i)
		f.resources.DirtIncome(5)
		f.husband.HusbandLogic()
		f.wife.WifeLogic()
		f.cat.CatLogic()

		if !f.husband.IsAlive || !f.wife.IsAlive || !f.cat.IsAlive {

			fmt.Println("Someone is dead")
			break
		}

		f.husband.GetData()
		f.wife.GetData()
		f.cat.GetData()
		f.resources.GetResources()

	}

	fmt.Println("All alive\nStatistics: \n")
	f.resources.Totals.GetStatistics()
}
