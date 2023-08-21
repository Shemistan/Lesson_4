package main

import (
	"fmt"

	d "github.com/Shemistan/Lesson_4/dependencies"
)

func main() {
	var foodStat, coatNumber, moneyEarned int
	var foodToEat, foodToBuy int16
	foodToEat = 30
	foodToBuy = 90
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
		house.Dirt += 5
		if husband.Satiety <= 0 {
			fmt.Println("Husband has died of hunger")
			break
		}
		if wife.Satiety <= 0 {
			fmt.Println("Wife has died of hunger")
			break
		}
		if husband.Happiness <= 10 {
			fmt.Println("Husband has died of depression")
			fmt.Print(i)
			break
		}
		if wife.Happiness <= 10 {
			fmt.Println("Wife has died of depression")
			fmt.Print(i)
			break
		}
		switch {
		case husband.Satiety <= 10 && husband.Happiness > 10:
			husband.HusbandEat(foodToEat)
			foodStat += 30
		case wife.Satiety <= 10 && wife.Happiness > 10:
			wife.WifeEat(foodToEat)
			foodStat += 30
		case house.Dirt >= 90:
			husband.Happiness -= 10
			wife.Happiness -= 10
			wife.CleanHouse()
		case house.Money < 300 && husband.Happiness > 20:
			husband.GoWork()
			moneyEarned += 150
		case house.Food <= 70 && wife.Happiness > 20:
			wife.BuyGroceries(foodToBuy)
		case husband.Happiness <= 20: //&& husband.Satiety > 20:
			husband.PlayComputer()
		case wife.Happiness <= 20 && house.Money >= 350: //&& wife.Satiety > 20:
			wife.BuyCoat()
			coatNumber++
		}
	}
	fmt.Printf("In a year the family ate %d units of food\n", foodStat)
	fmt.Printf("In a year the family earned %d units of money\n", moneyEarned)
	fmt.Printf("By the end of the year they have %d units of money\n", house.Money)
	fmt.Printf("In a year the family bought %d coats", coatNumber)
}
