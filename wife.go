package main

import "fmt"

type Wife struct {
	Human
	coat CoatSpecs
	// Кол-во убираемой грязи
	cleanDirtPoints int32
	// Коэфицент Еды к Деньгам
	foodToMoneyRatio int32
}

type CoatSpecs struct {
	// Кол-во очков счастья за шубу
	HappinessPoints int32
	//Стоимость шубы
	Price int32
}

// BuyProducts Сделано так чтобы можно было легко изменять формулу расчета денег к продукту
func (wife *Wife) BuyProducts(money int32) int32 {
	wife.wasteSatiety()
	products := money * wife.foodToMoneyRatio
	return products
}

// BuyCoat Возвращает стоимость шубы
func (wife *Wife) BuyCoat(money int32) int32 {
	if money < wife.coat.Price {
		fmt.Print("Not enough money for coat")
		return money
	}
	wife.wasteSatiety()
	wife.increaseHappiness(wife.coat.HappinessPoints)
	return money - wife.coat.Price
}

// CleanUp Возвращает кол-во очков грязи после уборки
func (wife *Wife) CleanUp(dirtPoints int32) int32 {
	wife.wasteSatiety()
	if dirtPoints > wife.cleanDirtPoints {
		return dirtPoints - wife.cleanDirtPoints
	}
	return 0
}
