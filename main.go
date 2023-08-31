package main

import (
	"fmt"

	d "github.com/Shemistan/Lesson_4/dependencies"
)

const (
	minimumPossibleFood             int16 = 70
	minimumPossibleMoney            int16 = 300
	minimumPossibleHappiness        int16 = 20
	minimumPossibleSatiety          int16 = 10
	deadlyLevelOfHappiness          int16 = 10
	dailyHappinessDecreaseDueToDirt int16 = 10
	maximumDirtAmount               int16 = 90
	deadlyLevelOfSatiety            int16 = 0
	dailyDirtIncreaseRate           int16 = 5
)

func main() {

	var (
		foodStat, coatNumber, moneyEarned int
		foodToEat                         int16 = 30
		foodToBuy                         int16 = 90
	)

	house := d.House{
		Money: 100,
		Food:  30,
		Dirt:  0,
	}
	husband := d.Husband{
		Person: d.Person{
			Name:      "John",
			Satiety:   30,
			Happiness: 100,
		},
		House: &house,
	}
	wife := d.Wife{
		Person: d.Person{
			Name:      "Jennifer",
			Satiety:   30,
			Happiness: 100,
		},
		House: &house,
	}
	for i := 1; i <= 365; i++ {
		house.Dirt += dailyDirtIncreaseRate
		if husband.Satiety <= deadlyLevelOfSatiety {
			fmt.Println("Husband has died of hunger")
			break
		}
		if wife.Satiety <= deadlyLevelOfSatiety {
			fmt.Println("Wife has died of hunger")
			break
		}
		if husband.Happiness <= deadlyLevelOfHappiness {
			fmt.Println("Husband has died of depression")
			fmt.Print(i)
			break
		}
		if wife.Happiness <= deadlyLevelOfHappiness {
			fmt.Println("Wife has died of depression")
			fmt.Print(i)
			break
		}
		switch {
		case husband.Satiety <= minimumPossibleSatiety && husband.Happiness > deadlyLevelOfHappiness:
			husband.HusbandEat(foodToEat)
			foodStat += int(foodToEat)
		case wife.Satiety <= minimumPossibleSatiety && wife.Happiness > deadlyLevelOfHappiness:
			wife.WifeEat(foodToEat)
			foodStat += int(foodToEat)
		case house.Dirt >= maximumDirtAmount:
			husband.Happiness -= dailyHappinessDecreaseDueToDirt
			wife.Happiness -= dailyHappinessDecreaseDueToDirt
			wife.CleanHouse()
		case house.Money < minimumPossibleMoney && husband.Happiness > minimumPossibleHappiness:
			husband.GoWork()
			moneyEarned += int(d.SalaryDaily)
		case house.Food <= minimumPossibleFood && wife.Happiness > minimumPossibleHappiness:
			wife.BuyGroceries(foodToBuy)
		case husband.Happiness <= minimumPossibleHappiness: //&& husband.Satiety > 20:
			husband.PlayComputer()
		case wife.Happiness <= minimumPossibleHappiness && house.Money >= d.PriceOfCoat: //&& wife.Satiety > 20:
			wife.BuyCoat()
			coatNumber++
		}
	}
	fmt.Printf("In a year the family ate %d units of food\n", foodStat)
	fmt.Printf("In a year the family earned %d units of money\n", moneyEarned)
	fmt.Printf("By the end of the year they have %d units of money\n", house.Money)
	fmt.Printf("In a year the family bought %d coats", coatNumber)
}
