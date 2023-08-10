package service

import (
	"errors"
	"fmt"

	"github.com/Shemistan/Lesson_4/models"
)

const (
	LIMIT_DIRTY       = 90  // предел грязи
	TOTAL_DAY_OF_LIFE = 365 // цикл жизни
	LIMIT_HAPPINESS   = 15  // предел счастья
	LIMIT_SATIETES    = 10  // предел сытости
	FOOD_COUNT        = 40  // кол-во покумаемой еды
	PRICE_COAT        = 350 // цена шубы
	EAT_FOOD_HUSBAND  = 30  //  кол-во еды за раз у мужа
	EAT_FOOD_WIFE     = 20  // кол-во еды за раз у жены
	SEPARATOR         = "==============================================================="
)

var (
	ErrHusbandIsDeadOfEat    = errors.New("game over, Husband died of starvation")
	ErrHusbandIsDeadNotHappy = errors.New("game over, Husband died unhappy")
	ErrWifeIsDeadOfEat       = errors.New("game over, Wife died of starvation")
	ErrWifeIsDeadNotHappy    = errors.New("game over, Wife died unhappy")
)

// YearOfLife год жизни семьи
func YearOfLife(f *models.Family) (models.Summary, error) {
	var totalFood, totalBuy int

	for i := 1; i <= TOTAL_DAY_OF_LIFE; i++ {
		fmt.Println(SEPARATOR)
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
	if f.Husband.Satiety <= EAT_FOOD_HUSBAND && f.House.Food >= EAT_FOOD_HUSBAND {
		f.EatHusband()
		*totalFood += EAT_FOOD_HUSBAND
		return
	}
	if f.Husband.Happiness < LIMIT_HAPPINESS {
		f.PlayComputer()
		return
	}
	if f.Husband.Satiety > LIMIT_SATIETES && f.Husband.Happiness >= LIMIT_SATIETES {
		f.GoToWork()
		return
	}
}

// actionWife действия жены в день
func actionWife(f *models.Family, totalFood, totalBuy *int) {
	if f.Wife.Satiety <= LIMIT_SATIETES && f.Wife.Happiness >= LIMIT_SATIETES && f.House.Food >= EAT_FOOD_WIFE {
		f.EatWife()
		*totalFood += EAT_FOOD_WIFE
		return
	}
	if f.Wife.Satiety > LIMIT_SATIETES && f.Wife.Happiness >= LIMIT_SATIETES && f.House.Money >= FOOD_COUNT && f.House.Food <= FOOD_COUNT+10 {
		f.BuyEat()
		return
	}
	if f.Wife.Happiness < LIMIT_HAPPINESS && f.House.Money >= PRICE_COAT+FOOD_COUNT { // 350 + 40 - чтобы хватило на еду :)
		f.BuyCoats()
		*totalBuy += 1
		return
	}
	if f.House.Dirty > LIMIT_DIRTY {
		f.CleanDirty()
		return
	}
}

// checkLifePerson  проверяем сытость и счастливость
// проверяю раздельно чтобы знать кто именно и от чего умер...
func checkLifePerson(f *models.Family) error {
	switch {
	case f.Husband.Happiness < LIMIT_HAPPINESS:
		return ErrHusbandIsDeadNotHappy
	case f.Husband.Satiety <= 0:
		return ErrHusbandIsDeadOfEat
	case f.Wife.Happiness < LIMIT_HAPPINESS:
		return ErrWifeIsDeadNotHappy
	case f.Wife.Satiety <= 0:
		return ErrWifeIsDeadOfEat
	default:
		return nil
	}
}
