package main

const shubaPrice = 350            // цена одной шубы
const year = 365                  // дней в году
const semiyear = 183              // дней- 6 месяцев
const quarter = 91                // дней- в квартале
const bottomLimitSatiation = 0    // Нижный порог сытости, ниже которого человек умирает.
const dirtIncrement = 5           // Накапливаемая грязь в доме в день.
const dirtEffectiveAmount = 100   // Мин. кол. грязи при котором уборка дома эффективна.
const topDirtLevel = 90           // Макс. кол. грязи после которого у человека уменьшается степень счастья.
const dailyHappinessLoss = 10     // Дневная потеря степени счастья.
const bottomHappinessLimit = 20   // Нижняя гран степени счастья после которого человек умирает, с запасом 20 единиц.
const topFoodStockLevel = 300     // Верхний потолок запаса еды.
const foodExpensePerMeal = 30     // Расход еды - человек/день.
const depletionOfSatiation = 10   // Уменьшение сытости за одно действие.
const boostHappinessPlay = 20     // Увеличение степени счастья при игре на комп.
const boostHappinessBuyShuba = 60 // Увеличение степени счастья при покупке шубы.
const dailyIncomeFromJob = 150    // Доход от дневной работы.

var (
	food              = 50  // Еда в начале периода
	money             = 100 // Деньги в начале периода
	husbandSatiation  = 30  // Сытость мужа в начале перида
	wifeSatiation     = 30  // Сытость жены в начале перида
	husbandHappiness  = 100 // Степень счастья мужа в начале перида
	wifeHappy         = 100 // Степень счастья жены в начале перида
	dirt              = 0   // Степень загрязненности дома в начале периода
	shubaAmount       int   // Количество шуб преобретенных в течении периода
	totalConsumedMeal int   // Общее количество съеденной еды в течении периода
)

func main() {
	var runProgram = true
	for runProgram {
		getRecordType := display()
		calc(getRecordType)
		getRes := repeatReport()
		if getRes == 2 {
			runProgram = false
		}
	}
}
