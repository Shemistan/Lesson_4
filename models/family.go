package models

import "fmt"

const (
	LimitDirty       = 90  // Если в доме грязи больше 90 - у людей падает степень счастья
	dirtyOneDay      = 5   // каждый день прибавляется грязь на 5
	decrHappiness    = 10  // уменьшаем счастье
	decrSatietes     = 10  // уменьшаем сытость
	addMoney         = 150 // заработок мужа в день
	husbandHappiness = 20  // увеличивает счастье мужа от игры в комп
	wifeHappiness    = 60  // увеличивает счастье  жены от покупки шубы
	FoodCount        = 40  // кол-во покупаемой еды
	PriceCoat        = 350 // цена шубы
	cleanDirty       = 100 // миинус кол-во грязи за одну уборку
	EatFood          = 30  // кол-во еды за раз
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
	f.Husband.Satiety -= decrSatietes
	fmt.Println(f.Husband.Name, "- сходил на работу")
}

// PlayComputer муж играет
func (f *Family) PlayComputer() {
	f.Husband.Happiness += husbandHappiness
	f.Husband.Satiety -= decrSatietes
	fmt.Println(f.Husband.Name, "- поиграл в комьютер")
}

// BuyEat жена идет за покупками
func (f *Family) BuyEat() {
	f.House.Money -= FoodCount // 1 единица еды стоит 1 единицы денег
	f.House.Food += FoodCount
	f.Wife.Satiety -= decrSatietes
	fmt.Println(f.Wife.Name, "- купила eду")
}

// BuyCoats жена покупает шубу
func (f *Family) BuyCoats() {
	f.House.Money -= PriceCoat
	f.Wife.Happiness += wifeHappiness
	f.Wife.Satiety -= decrSatietes
	fmt.Println(f.Wife.Name, "- купила шубу")
}

// CleanDirty жена убирается дома
func (f *Family) CleanDirty() {
	f.House.Dirty -= cleanDirty
	f.Wife.Satiety -= decrSatietes
	fmt.Println(f.Wife.Name, "- убралась дома")
}

// EatPerson  есть еду
func (f *Family) EatPerson(who bool) {
	if who {
		f.Husband.Satiety += EatFood
		fmt.Println(f.Husband.Name, "- покушал")
	} else {
		f.Wife.Satiety += EatFood
		fmt.Println(f.Wife.Name, "- покушала")
	}
	f.House.Food -= EatFood
}
