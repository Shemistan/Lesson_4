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
	if house.food > 60 { //если хватает обоих загрузить едой
		a.fullness += 30 //зависит от наличия еды в холодильнике
		house.food -= 30
		house.foodstat += 30
	} else {
		a.fullness += house.food / 2
		house.foodstat += house.food / 2
		house.food /= 2
	}
	fmt.Println(a.name, "Поел еды")
}

// игра на компьютере мужа
func (a *Male) play() {
	a.happiness += 20
	a.fullness -= 10
	fmt.Println(a.name, "Поиграл на компе")
}

// заработок
func (a *Male) earn(house *House) {
	house.money += 150
	house.moneystat += 150
	a.fullness -= 10
	fmt.Println(a.name, "Заработал денег")
}

// жена чистит дом
func (a *Female) clean(house *House) {
	if house.cleanliness >= 100 {
		house.cleanliness -= 100
	} else {
		house.cleanliness = 0
	}
	a.fullness -= 10
	fmt.Println(a.name, "Почистила квартиру")
}

// Жена покупает шубу
func (a *Female) buyCoat(house *House) {
	if house.money >= 350 {
		house.money -= 350
		a.happiness += 60
		a.fullness -= 10
		fmt.Println(a.name, "Купила шубу")
		house.shubstat += 1
	} else {
		fmt.Println(a.name, "Не купила шубу, нет денег")
	}
}

// Жена покупает продукты, 60 или меньше, если денег недостаточно
func (a *Female) buyEat(house *House) {
	if house.money > 60 {
		house.money -= 60
		house.food += 60
		fmt.Println(a.name, "Купила еды: ", 60)
	} else {
		house.food += house.money
		fmt.Println(a.name, "Купила еды: ", house.money)
		house.money = 0
	}
	a.fullness -= 10
}
