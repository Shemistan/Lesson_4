package main

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

// Сделано так чтобы можно было легко изменять формулу расчета денег к продукту
func (wife *Wife) buyProducts(money int32) int32 {
	wife.wasteSatiety()
	products := money * wife.foodToMoneyRatio
	return products
}

// Возвращает стоимость шубы
func (wife *Wife) buyCoatAndGetPrice() int32 {
	wife.wasteSatiety()
	wife.increaseHappiness(wife.coat.HappinessPoints)
	return wife.coat.Price
}

// Возвращает кол-во очков грязи после уборки
func (wife *Wife) cleanUp(dirtPoints int32) int32 {
	wife.wasteSatiety()
	if dirtPoints > wife.cleanDirtPoints {
		return dirtPoints - wife.cleanDirtPoints
	}
	return 0
}
