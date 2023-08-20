package human

import (
	"errors"
)

const (
	foodToMoneyRatio               = 1
	cleanDirtPoints                = 100
	coatHappinessPoints            = 60
	coatPrice                      = 350
	notEnoughMoneyCoatErrorMessage = "not enough money for coat"
)

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

// BuyProducts Возвращает кол-во продуктов после похода в магазин
// Сделано так чтобы можно было легко изменять формулу расчета денег к продукту
func (wife *Wife) BuyProducts(money int32) int32 {
	wife.wasteSatiety()
	products := money * wife.foodToMoneyRatio
	return products
}

// BuyCoat Возвращает кол-во денег после покупки шубы
func (wife *Wife) BuyCoat(money int32) (err error, remainingMoney int32) {
	if money < wife.coat.Price {
		return errors.New(notEnoughMoneyCoatErrorMessage), remainingMoney
	}
	wife.wasteSatiety()
	wife.increaseHappiness(wife.coat.HappinessPoints)
	return nil, money - wife.coat.Price
}

// CleanUp Возвращает кол-во очков грязи после уборки
func (wife *Wife) CleanUp(dirtPoints int32) int32 {
	wife.wasteSatiety()
	if dirtPoints > wife.cleanDirtPoints {
		return dirtPoints - wife.cleanDirtPoints
	}
	return 0
}

func CreateWife(name string) Wife {
	return Wife{
		foodToMoneyRatio: foodToMoneyRatio,
		cleanDirtPoints:  cleanDirtPoints,
		coat: CoatSpecs{
			HappinessPoints: coatHappinessPoints,
			Price:           coatPrice,
		},
		Human: CreateHuman(name),
	}
}
