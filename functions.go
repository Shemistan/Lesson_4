package main

// Расчет еды
func eat(satiationLevel int) int {
	food = food - foodExpensePerMeal
	totalConsumedMeal += foodExpensePerMeal
	return satiationLevel + foodExpensePerMeal
}

// Расчет степени счастья и сытости от игры в комп.
func play() {
	husbandHappiness = husbandHappiness + boostHappinessPlay
	husbandSatiation = husbandSatiation - depletionOfSatiation
}

// Расчет дохода от работы мужа.
func job() {
	money = money + dailyIncomeFromJob
	husbandSatiation = husbandSatiation - depletionOfSatiation
}

// Расчет запаса продукта, расхода от покупки продуктов и уменьшения степени сытости жены.
func buyProduct(buyProdAmount int) {
	food = food + buyProdAmount
	wifeSatiation = wifeSatiation - depletionOfSatiation
	money = money - buyProdAmount
}

// Расчет денежного баланса, степени счастья и сытости жены при покупке шубы.
func buyShuba() {
	money = money - shubaPrice
	wifeHappy = wifeHappy + boostHappinessBuyShuba
	wifeSatiation = wifeSatiation - depletionOfSatiation
}

// Расчет количества грязи при уборке домаб а также степени сытости жены.
func cleaning() {
	dirt = dirt - dirtEffectiveAmount
	wifeSatiation = wifeSatiation - depletionOfSatiation
}
