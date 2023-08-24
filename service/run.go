package service

import (
	"errors"
	"fmt"
	"github.com/Shemistan/Lesson_4/models"
)

const (
	totalDayOfLife           = 365 // цикл жизни
	limitHappiness           = 15  // предел счастья
	limitSatiety             = 10  // предел сытости
	Separator                = "================================================"
	errHusbandIsDeadOfEat    = "GG! Husband died of starvation"
	errHusbandIsDeadNotHappy = "GG! Husband died unhappy"
	errWifeIsDeadOfEat       = "GG! Wife died of starvation"
	errWifeIsDeadNotHappy    = "GG! Wife died unhappy"
)

// YearOfLife год жизни семьи
func YearOfLife(f *models.Family) (models.Summary, error) {
	var totalFood, totalBuy int

	for i := 1; i <= totalDayOfLife; i++ {
		fmt.Println(Separator)
		fmt.Println("День №", i)
		err := checkLifePerson(f)
		if err != nil {
			return models.Summary{}, err
		}
		oneDayLife(f, &totalFood, &totalBuy)
	}

	return models.Summary{
		TotalMoney:     f.House.Money,
		TotalFoodEaten: totalFood,
		TotalBuyCoats:  totalBuy,
	}, nil
}

// oneDayLife один день жизни семьи
func oneDayLife(f *models.Family, totalFood, totalBuy *int) {
	f.IncreaseDirtiness()
	actionHusband(f, totalFood)
	actionWife(f, totalFood, totalBuy)
}

// actionHusband действия мужа в день
func actionHusband(f *models.Family, totalFood *int) {
	if f.Husband.Satiety <= models.EatFood && f.House.Food >= models.EatFood {
		f.EatPerson(true)
		*totalFood += models.EatFood
		return
	}
	if f.Husband.Happiness < limitHappiness {
		f.PlayComputer()
		return
	}
	if f.Husband.Satiety > limitSatiety && f.Husband.Happiness >= limitSatiety {
		f.GoToWork()
		return
	}
}

// actionWife действия жены в день
func actionWife(f *models.Family, totalFood, totalBuy *int) {
	if f.Wife.Satiety <= limitSatiety && f.Wife.Happiness >= limitSatiety && f.House.Food >= models.EatFood {
		f.EatPerson(false)
		*totalFood += models.EatFood
		return
	}
	if f.Wife.Satiety > limitSatiety && f.Wife.Happiness >= limitSatiety && f.House.Money >= models.FoodCount &&
		f.House.Food <= models.FoodCount+10 {
		f.BuyEat()
		return
	}
	if f.Wife.Happiness < limitHappiness && f.House.Money >= models.PriceCoat+models.FoodCount { // 350 + 40 - чтобы хватило на еду :)
		f.BuyCoats()
		*totalBuy += 1
		return
	}
	if f.House.Dirty > models.LimitDirty {
		f.CleanDirty()
		return
	}
}

// checkLifePerson  проверяем сытость и счастливость
// проверяю раздельно чтобы знать кто именно и от чего умер...
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
