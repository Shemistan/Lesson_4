package models

import "fmt"

const (
	LimitDirty       = 90  // если в доме грязи больше 90 - у людей падает степень счастья
	dirtyOneDay      = 5   // каждый день прибавляется грязь на 5
	decrHappiness    = 10  // уменьшение счастья
	decrSatiety      = 10  // уменьшение сытости
	addMoney         = 150 // заработок мужа в день
	husbandHappiness = 20  // увеличивание счастья мужа от игры в компьютер
	wifeHappiness    = 60  // увеличивание счастья  жены от покупки шубы
	FoodCount        = 40  // кол-во покупаемой еды
	PriceCoat        = 350 // цена шубы
	cleanDirty       = 100 // минус кол-во грязи за одну уборку
	EatFood          = 30  // кол-во съеденой еды за раз
)

// Family общая структура семьи
type Family struct {
	Husband Person
	Wife    Person
	House   House
}

// Person структура члена семьи
type Person struct {
	Name      string
	Satiety   int
	Happiness int
}

// House структура дома
type House struct {
	Money int
	Food  int
	Dirty int
}

// IncreaseDirtiness увеличивает грязь каждый день на 5
func (f *Family) IncreaseDirtiness() {
	f.House.Dirty += dirtyOneDay
	if f.House.Dirty > LimitDirty {
		f.Wife.Happiness -= decrHappiness
		f.Husband.Happiness -= decrHappiness
	}
}

// GoToWork  муж идет на работу
func (f *Family) GoToWork() {
	f.House.Money += addMoney
	f.Husband.Satiety -= decrSatiety
	fmt.Println(f.Husband.Name, "- сходил на работу, сытость мужа: ", f.Husband.Satiety)
}

// PlayComputer муж играет
func (f *Family) PlayComputer() {
	f.Husband.Happiness += husbandHappiness
	f.Husband.Satiety -= decrSatiety
	fmt.Println(f.Husband.Name, "- поиграл в комьютер, степень счастья мужа: ", f.Husband.Happiness)
}

// BuyEat жена идет за покупками
func (f *Family) BuyEat() {
	f.House.Money -= FoodCount // 1 единица еды стоит 1 единицы денег
	f.House.Food += FoodCount
	f.Wife.Satiety -= decrSatiety
	fmt.Println(f.Wife.Name, "- купила eду, сытость жены: ", f.Wife.Satiety)
}

// BuyCoats жена покупает шубу
func (f *Family) BuyCoats() {
	f.House.Money -= PriceCoat
	f.Wife.Happiness += wifeHappiness
	f.Wife.Satiety -= decrSatiety
	fmt.Println(f.Wife.Name, "- купила шубу, кол-во денег дома: ", f.House.Money)
}

// CleanDirty жена убирается дома
func (f *Family) CleanDirty() {
	f.House.Dirty -= cleanDirty
	f.Wife.Satiety -= decrSatiety
	fmt.Printf("%v - убралась дома, сытость жены: %v, счастье: %v \n", f.Wife.Name, f.Wife.Satiety, f.Wife.Happiness)
}

// EatPerson  есть еду
func (f *Family) EatPerson(who bool) {
	if who {
		f.Husband.Satiety += EatFood
		fmt.Println(f.Husband.Name, "- покушал, кол-во еды дома: ", f.House.Food-EatFood)
	} else {
		f.Wife.Satiety += EatFood
		fmt.Println(f.Wife.Name, "- покушала, кол-во еды дома: ", f.House.Food-EatFood)
	}
	f.House.Food -= EatFood
}
