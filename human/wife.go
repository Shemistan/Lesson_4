package human

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

// BuyProducts Возвращает кол-во продуктов после похода в магазин
// Сделано так чтобы можно было легко изменять формулу расчета денег к продукту
func (wife *Wife) BuyProducts(money int32) int32 {
	wife.wasteSatiety()
	products := money * wife.foodToMoneyRatio
	return products
}

// BuyCoat Изменяет кол-во денег после покупки шубы
func (wife *Wife) BuyCoat(money *int32) {
	if *money < wife.coat.Price {
		fmt.Print("Not enough money for coat")
		return
	}
	wife.wasteSatiety()
	wife.increaseHappiness(wife.coat.HappinessPoints)
	*money -= wife.coat.Price
}

// CleanUp Возвращает кол-во очков грязи после уборки
func (wife *Wife) CleanUp(dirtPoints int32) int32 {
	wife.wasteSatiety()
	if dirtPoints > wife.cleanDirtPoints {
		return dirtPoints - wife.cleanDirtPoints
	}
	return 0
}

func WifeFactory(name string) Wife {
	return Wife{
		foodToMoneyRatio: 1,
		cleanDirtPoints:  100,
		coat: CoatSpecs{
			HappinessPoints: 60,
			Price:           350,
		},
		Human: Factory(name),
	}
}
