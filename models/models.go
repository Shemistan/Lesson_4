package models

import "fmt"

const (
	DefaultHungerDecrease         = 10
	DailyDirtIncrease             = 5
	MaxFoodConsumption            = 30
	DirtyThreshold                = 90
	HappinessDecrease             = 10
	MoneyPerWorkday               = 150
	FoodCost                      = 40
	FoodUnitsPerCost              = 10
	ComputerGameHappinessIncrease = 20
	CoatPurchaseHappinessIncrease = 60
	CoatCost                      = 350
	MaxCleaningCapacity           = 100
	LimitSatietyToDead            = 0
	LimitHappinessToDead          = 10
)

type Human struct {
	Name      string
	Satiety   int
	Happiness int
}
type House struct {
	Budget       int
	FoodSupplies int
	DirtInHouse  int
}
type Family struct {
	Husband Human
	Wife    Human
	House   House
}

type Result struct {
	TotalMoney     int
	TotalFoodEaten int
	TotalBuyCoats  int
}

func (f *Family) IncreaseDirtiness() {
	f.House.DirtInHouse += DailyDirtIncrease
	if f.House.DirtInHouse > DirtyThreshold {
		f.Wife.Happiness -= HappinessDecrease
		f.Husband.Happiness -= HappinessDecrease
	}
}

func (f *Family) GoWork() {
	f.House.Budget += MoneyPerWorkday
	f.Husband.Satiety -= DefaultHungerDecrease
	fmt.Printf("%v сходил на работу", f.Husband.Name)
}

func (f *Family) PlayComputerGames() {
	f.Husband.Happiness += ComputerGameHappinessIncrease
	f.Husband.Satiety -= DefaultHungerDecrease
	fmt.Printf("%v сыграл в компьютерные игры", f.Husband.Name)
}

func (f *Family) BuyCoats() {
	f.House.Budget -= CoatCost
	f.Wife.Happiness += CoatPurchaseHappinessIncrease
	f.Wife.Satiety -= DefaultHungerDecrease
	fmt.Printf("%v купила шубу", f.Wife.Name)
}

func (f *Family) CleanHouse() {
	f.House.DirtInHouse -= MaxCleaningCapacity
	f.Wife.Satiety -= DefaultHungerDecrease
	fmt.Printf("%v убрала дом", f.Wife.Name)
}

func (f *Family) BuyGroceries() {
	f.House.Budget -= FoodCost
	f.House.FoodSupplies += FoodCost
	f.Wife.Satiety -= DefaultHungerDecrease

	fmt.Printf("%v купила продукты", f.Wife.Name)
}

func (f *Family) Eat(person bool) {
	if person {
		f.Husband.Satiety += MaxFoodConsumption
		fmt.Printf("%v поел", f.Husband.Name)
	} else {
		f.Wife.Satiety += MaxFoodConsumption
		fmt.Printf("%v поелa", f.Wife.Name)
	}
	f.House.FoodSupplies -= MaxFoodConsumption
}

/*
func (f *Family) IsAliveHusband() bool {
	if f.Husband.Satiety <= LimitSatietyToDead || f.Husband.Happiness <= HappinessLimitToDead {
		return false
	}

	return true
}

func (f *Family) IsAliveWife() bool {
	if f.Wife.Satiety <= LimitSatietyToDead || f.Wife.Happiness <= HappinessLimitToDead {
		return false
	}

	return true
}
*/
