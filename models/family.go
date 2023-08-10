package models

import "fmt"

const (
	LIMIT_DIRTY      = 90  // Если в доме грязи больше 90 - у людей падает степень счастья
	DIRTY_ONE_DAY    = 5   // каждый день прибавляктся грязь на 5
	DECR_HAPPINES    = 10  // уменьшаем счастье
	DECR_SATIETES    = 10  // уменьшаем сытость
	ADD_MONEY        = 150 // заработок мужа в день
	HUSBAND_HAPPINES = 20  // увеличивает счастье мужа от игры в комп
	WIFE_HAPPINES    = 60  // увеличивает счастье  жены от покупки шубы
	FOOD_COUNT       = 40  // кол-во покумаемой еды
	PRICE_COAT       = 350 // цена шубы
	CLEAN_DIRTY      = 100 // миинус кол-во грязи за одну уборку
	EAT_FOOD_HUSBAND = 30  //  кол-во еды за раз у мужа
	EAT_FOOD_WIFE    = 20  // кол-во еды за раз у жены
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
	f.House.Dirty += DIRTY_ONE_DAY
	if f.House.Dirty > LIMIT_DIRTY {
		f.Wife.Happiness -= DECR_HAPPINES
		f.Husband.Happiness -= DECR_HAPPINES
	}
}

// GoToWork  муж идет на работу
func (f *Family) GoToWork() {
	f.House.Money += ADD_MONEY
	f.Husband.Satiety -= DECR_SATIETES
	fmt.Println(f.Husband.Name, "- сходил на работу")
}

// PlayComputer муж играет
func (f *Family) PlayComputer() {
	f.Husband.Happiness += HUSBAND_HAPPINES
	f.Husband.Satiety -= DECR_SATIETES
	fmt.Println(f.Husband.Name, "- поиграл в комьютер")
}

// BuyEat жена идет за покупками
func (f *Family) BuyEat() {
	f.House.Money -= FOOD_COUNT // 1 еденица еды стоит 1 еденицы денег
	f.House.Food += FOOD_COUNT
	f.Wife.Satiety -= DECR_SATIETES
	fmt.Println(f.Wife.Name, "- купила eду")
}

// BuyCoats жена покупает шубу
func (f *Family) BuyCoats() {
	f.House.Money -= PRICE_COAT
	f.Wife.Happiness += WIFE_HAPPINES
	f.Wife.Satiety -= DECR_SATIETES
	fmt.Println(f.Wife.Name, "- купила шубу")
}

// CleanDirty жена убирается дома
func (f *Family) CleanDirty() {
	f.House.Dirty -= CLEAN_DIRTY
	f.Wife.Satiety -= DECR_SATIETES
	fmt.Println(f.Wife.Name, "- убралась дома")
}

// EatWife жена ест
func (f *Family) EatWife() {
	f.Wife.Satiety += EAT_FOOD_WIFE
	f.House.Food -= EAT_FOOD_WIFE
	fmt.Println(f.Wife.Name, "- покушала")
}

// EatHusband муж ест
func (f *Family) EatHusband() {
	f.Husband.Satiety += EAT_FOOD_HUSBAND
	f.House.Food -= EAT_FOOD_HUSBAND
	fmt.Println(f.Husband.Name, "- покушал")
}
