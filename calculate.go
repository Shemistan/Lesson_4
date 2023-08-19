package main

import (
	"fmt"
	"math"
)

func calc(getRecordType int) {
	var getRecordTypeToStr string
	for day := 0; day <= getRecordType; day++ {
		wifeBool := false    // флаг- для счета действий за день для жены.
		husbandBool := false // флаг- для счета действий за день для мужа.
		dirt = dirt + dirtIncrement

		if bottomLimitSatiation == husbandSatiation {
			husbandSatiation = eat(husbandSatiation)
			husbandBool = true
		}

		if bottomLimitSatiation == wifeSatiation {
			wifeSatiation = eat(wifeSatiation)
			wifeBool = true
		}

		if dirt > dirtEffectiveAmount {
			cleaning()
			wifeBool = true
		}

		if dirt > topDirtLevel {
			wifeHappy = wifeHappy - dailyHappinessLoss
			husbandHappiness = husbandHappiness - dailyHappinessLoss
		}

		if husbandHappiness < bottomHappinessLimit && husbandBool == false {
			play()
			husbandBool = true
		}
		if wifeHappy < bottomHappinessLimit && wifeBool == false {
			buyShuba()
			shubaAmount++
			wifeBool = true
		}

		if food < topFoodStockLevel && wifeBool == false {
			buyProduct(int(math.Min(float64(money), float64(topFoodStockLevel))))
			wifeBool = true
		}

		if !husbandBool {
			job()
		}

		if !wifeBool {
			wifeSatiation = eat(wifeSatiation)
		}
	}
	if getRecordType == year {
		getRecordTypeToStr = "ГОДОВОЙ"
	}
	if getRecordType == semiyear {
		getRecordTypeToStr = "ПОЛУГОДОВОЙ"
	}
	if getRecordType == quarter {
		getRecordTypeToStr = "КВАРТАЛЬНЫЙ"
	}
	fmt.Println("\n===== ", getRecordTypeToStr, "ОТЧЕТ =====")
	//fmt.Println(" Зар. денег: ", money, "\n Cьедено еды: ", totalConsumedMeal, "\n Куплено шуб: ", shubaAmount)
	fmt.Printf("  %-15s  %-15v \n", "Зар. денег:", money)
	fmt.Printf("  %-15s  %-15v \n", "Cьедено еды:", totalConsumedMeal)
	fmt.Printf("  %-15s  %-15v \n", "Куплено шуб::", shubaAmount)
	fmt.Println("==============================")

}
