package models

const (
	Salary    = 150
	furPrice  = 350
	foodPrice = 1
	foodCount = 60

	DecrSatiety        = 10
	IncSatietyLimit    = 30
	LimitSatietyToDead = 0

	incHappinessByPC     = 20
	incHappinessByFur    = 60
	DecHappinessByDirty  = 10
	HappinessLimitToDead = 10

	incDirty     = 5
	maxCleanHome = 100
	limitDirty   = 90

	//Useless cat
	DecrCatSatiety        = 10
	catFoodCount          = 10
	catFoodPrice          = 1
	catSatietyByFood      = 2
	CatEatLimit           = 10
	CatlimitSatietyToDead = 0
	incDirtyByCat         = 5
)

type Creature struct {
	Name    string
	Satiety int
}

type Person struct {
	Creature
	Happiness int
}

type House struct {
	MoneyInNightstand int
	FoodInFridge      int
	CatFood           int
	DirtyInHouse      int
}

type Family struct {
	Husband Person
	Wife    Person
	Cat     Creature
	Home    House
}

type TotalCount struct {
	EarnedMoney int
	EatenFood   int
	BoughtFur   int
	DeadPerson  []string
}

func (f *Family) EatFood(isHusband bool) {
	if isHusband {
		f.Husband.Satiety += IncSatietyLimit
	} else {
		f.Wife.Satiety += IncSatietyLimit
	}

	f.Home.FoodInFridge -= IncSatietyLimit
}

func (f *Family) PlayPC() {
	f.Husband.Happiness += incHappinessByPC
	f.Husband.Satiety -= DecrSatiety
}

func (f *Family) GoToWork() {
	f.Home.MoneyInNightstand += Salary
	f.Husband.Satiety -= DecrSatiety
}

func (f *Family) BuyProducts() {
	f.Home.FoodInFridge += foodCount
	f.Home.MoneyInNightstand -= foodCount * foodPrice
	f.Wife.Satiety -= DecrSatiety

}

func (f *Family) BuyFur() {
	f.Home.MoneyInNightstand -= furPrice
	f.Wife.Happiness += incHappinessByFur
	f.Wife.Satiety -= DecrSatiety

}

func (f *Family) CleanHome() {
	if f.Home.DirtyInHouse > maxCleanHome {
		f.Home.DirtyInHouse -= maxCleanHome

	} else {
		f.Home.DirtyInHouse = 0
	}

	f.Wife.Satiety -= DecrSatiety
}

func (f *Family) DirtyEveryDay() {
	f.Home.DirtyInHouse += incDirty

	if f.Home.DirtyInHouse >= limitDirty {
		f.Wife.Happiness -= DecHappinessByDirty
		f.Husband.Happiness -= DecHappinessByDirty
	}
}

func (f *Family) EatFoodByCat() {
	f.Cat.Satiety += CatEatLimit * catSatietyByFood
	f.Home.CatFood -= CatEatLimit
}

func (f *Family) Sleep() {
	f.Cat.Satiety -= 10
}

func (f *Family) TearUpWalpaper() {
	f.Cat.Satiety -= 10
	f.Home.DirtyInHouse += 5
}

func (f *Family) PetCat(person *Person) {
	person.Satiety -= DecrSatiety
	person.Happiness += 5
}

func (f *Family) BuyCatFood() {
	f.Home.CatFood += foodCount
	f.Home.MoneyInNightstand -= foodCount * foodPrice
	f.Wife.Satiety -= DecrSatiety
}

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

func (f *Family) IsAliveCat() bool {
	if f.Cat.Satiety <= CatlimitSatietyToDead {
		return false
	}

	return true
}
