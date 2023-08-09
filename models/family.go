package models

import "fmt"

type Family struct {
	Husband Person
	Wife    Person
	House   House
}

type Person struct {
	Name      string
	Satiety   int
	Happiness int
}

type House struct {
	Money int
	Food  int
	Dirty int
}

// IncreaseDirtiness увеличивает грязь каждый день на 5
func (f *Family) IncreaseDirtiness() {
	f.House.Dirty += 5
	if f.House.Dirty > 90 {
		f.Wife.Happiness -= 10
		f.Husband.Happiness -= 10
	}
}

// Work  муж идет на работу
func (f *Family) Work() {
	f.House.Money += 150
	f.Husband.Satiety -= 10
	fmt.Println(f.Husband.Name, "- сходил на работу")
}

// Game муж играет
func (f *Family) Game() {
	f.Husband.Happiness += 20
	f.Husband.Satiety -= 10
	fmt.Println(f.Husband.Name, "- поиграл в комьютер")
}

// BuyEat жена идет за покупками
func (f *Family) BuyEat() {
	f.House.Money -= 40
	f.House.Food += 40
	f.Wife.Satiety -= 10
	fmt.Println(f.Wife.Name, "- купила eду")
}

// BuyCoats жена покупает шубу
func (f *Family) BuyCoats() {
	f.House.Money -= 350
	f.Wife.Happiness += 60
	f.Wife.Satiety -= 10
	fmt.Println(f.Wife.Name, "- купила шубу")
}

// CleanDirty жена убирается дома
func (f *Family) CleanDirty() {
	f.House.Dirty -= 100
	f.Wife.Satiety -= 10
	fmt.Println(f.Wife.Name, "- убралась дома")
}

// EatWife жена ест
func (f *Family) EatWife() {
	f.Wife.Satiety += 20
	f.House.Food -= 20
	fmt.Println(f.Wife.Name, "- покушала")
}

// EatHusband муж ест
func (f *Family) EatHusband() {
	f.Husband.Satiety += 30
	f.House.Food -= 30
	fmt.Println(f.Husband.Name, "- покушал")
}
