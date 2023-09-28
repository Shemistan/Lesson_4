package service

import (
	"errors"
	"fmt"

	"github.com/Shemistan/Lesson_4/models"
)

const (
	totalDayOfLife           = 365 // цикл жизни
	limitHappiness           = 15  // предел счастья
	limitSatietes            = 10  // предел сытости
	Separator                = "==============================================================="
	errHusbandIsDeadOfEat    = "game over, Husband died of starvation"
	errHusbandIsDeadNotHappy = "game over, Husband died unhappy"
	errWifeIsDeadOfEat       = "game over, Wife died of starvation"
	errWifeIsDeadNotHappy    = "game over, Wife died unhappy"
)

func YearOfLife(f *models.Family) (models.Result, error) {
	var totalFood, totalBuy int

	for i := 1; i <= totalDayOfLife; i++ {
		fmt.Println("----------")
		fmt.Println("День №", i)
		err := checkLifePerson(f)
		if err != nil {
			return models.Result{}, err
		}
		oneDayLife(f, &totalFood, &totalBuy)
	}

	return models.Result{
		TotalMoney:     f.House.Budget,
		TotalFoodEaten: totalFood,
		TotalBuyCoats:  totalBuy,
	}, nil
}

func oneDayLife(f *models.Family, totalFood, totalBuy *int) {
	f.IncreaseDirtiness()
	actionHusband(f, totalFood)
	actionWife(f, totalFood, totalBuy)
}

func actionHusband(f *models.Family, totalFood *int) {
	if f.Husband.Satiety <= models.MaxFoodConsumption && f.House.FoodSupplies >= models.MaxFoodConsumption {
		f.Eat(true)
		*totalFood += models.MaxFoodConsumption
		return
	}
	if f.Husband.Happiness < limitHappiness {
		f.PlayComputerGames()
		return
	}
	if f.Husband.Satiety > limitSatietes && f.Husband.Happiness >= limitSatietes {
		f.GoWork()
		return
	}
}

func actionWife(f *models.Family, totalFood, totalBuy *int) {
	if f.Wife.Satiety <= limitSatietes && f.Wife.Happiness >= limitSatietes && f.House.FoodSupplies >= models.MaxFoodConsumption {
		f.Eat(false)
		*totalFood += models.MaxFoodConsumption
		return
	}
	if f.Wife.Satiety > limitSatietes && f.Wife.Happiness >= limitSatietes && f.House.Budget >= models.MaxFoodConsumption &&
		f.House.FoodSupplies <= models.MaxFoodConsumption+10 {
		f.BuyGroceries()
		return
	}
	if f.Wife.Happiness < limitHappiness && f.House.Budget >= models.CoatCost+models.MaxFoodConsumption {
		f.BuyCoats()
		*totalBuy += 1
		return
	}
	if f.House.DirtInHouse > models.DirtyThreshold {
		f.CleanHouse()
		return
	}
}

func checkLifePerson(f *models.Family) error {
	switch {
	case f.Husband.Happiness < limitHappiness:
		return errors.New(errHusbandIsDeadNotHappy)
	case f.Husband.Satiety <= 0:
		return errors.New(errHusbandIsDeadOfEat)
	case f.Wife.Happiness < limitHappiness:
		return errors.New(errWifeIsDeadNotHappy)
	case f.Wife.Satiety <= 0:
		return errors.New(errWifeIsDeadOfEat)
	default:
		return nil
	}
}
