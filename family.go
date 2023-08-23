package main

import "fmt"

type People struct {
	name      string
	fullness  int
	happiness int
}
type Male struct {
	People
}
type Female struct {
	People
}

// кушают все, если есть еда в холодильнике
func (a *People) eat(house *House) {
	if house.food > 2*fullnessChange { //если хватает обоих загрузить едой
		a.fullness += fullnessChange //зависит от наличия еды в холодильнике
		house.food -= fullnessChange
		house.foodstat += fullnessChange
	} else {
		a.fullness += house.food / 2
		house.foodstat += house.food / 2
		house.food /= 2
	}
	fmt.Println(a.name, "Поел еды")
}

// игра на компьютере мужа
func (a *Male) play() {
	a.happiness += happinessChangeMan
	a.fullness -= fullnessDecrease
	fmt.Println(a.name, "Поиграл на компе")
}

// заработок
func (a *Male) earn(house *House) {
	house.money += moneyChange
	house.moneystat += moneyChange
	a.fullness -= fullnessDecrease
	fmt.Println(a.name, "Поработал, Заработал денег" , moneyChange)
}

// жена чистит дом
func (a *Female) clean(house *House) {
	if house.cleanliness >= cleanChange {
		house.cleanliness -= cleanChange
	} else {
		house.cleanliness = 0
	}
	a.fullness -= fullnessDecrease
	fmt.Println(a.name, "Почистила квартиру")
}

// Жена покупает шубу
func (a *Female) buyCoat(house *House) {
	if house.money >= coatPrice {
		house.money -= coatPrice
		a.happiness += happinessChangeWoman
		a.fullness -= fullnessDecrease
		fmt.Println(a.name, "Купила шубу")
		house.shubstat += 1
	} else {
		fmt.Println(a.name, "Не купила шубу, нет денег")
	}
}

// Жена покупает продукты, 60 или меньше, если денег недостаточно
func (a *Female) buyEat(house *House) {
	if house.money > foodChange {
		house.money -= foodChange
		house.food += foodChange
		fmt.Println(a.name, "Купила еды: ", foodChange)
	} else {
		house.food += house.money
		fmt.Println(a.name, "Купила еды: ", house.money)
		house.money = 0
	}
	a.fullness -= fullnessDecrease
}
