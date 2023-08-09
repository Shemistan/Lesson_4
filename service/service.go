package service

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/models"
)

// YearOfLife год жизни семьи
func YearOfLife(f *models.Family) models.Summary {
	var totalFood, totalBuy int

	for i := 1; i <= 365; i++ {
		fmt.Println("===============================================================")
		fmt.Println("День №", i)
		oneDayLife(f, &totalFood, &totalBuy)
	}

	return models.Summary{
		TotalMoney:     f.House.Money,
		TotalFoodEaten: totalFood,
		TotalBuyCoats:  totalBuy,
	}
}

// oneDayLife один день жизни семьи
func oneDayLife(f *models.Family, totalFood, totalBuy *int) {
	checkLifePerson(f)
	f.IncreaseDirtiness()
	actionHusband(f, totalFood)
	actionWife(f, totalFood, totalBuy)
}

// actionHusband действия мужа в день
func actionHusband(f *models.Family, totalFood *int) {
	if f.Husband.Satiety <= 30 && f.House.Food >= 30 {
		f.EatHusband()
		*totalFood += 30
		return
	}
	if f.Husband.Happiness < 15 {
		f.Game()
		return
	}
	if f.Husband.Satiety > 10 && f.Husband.Happiness >= 10 {
		f.Work()
		return
	}
}

// actionWife действия жены в день
func actionWife(f *models.Family, totalFood, totalBuy *int) {
	if f.Wife.Satiety <= 10 && f.Wife.Happiness >= 10 && f.House.Food >= 20 {
		f.EatWife()
		*totalFood += 20
		return
	}
	if f.Wife.Satiety > 10 && f.Wife.Happiness >= 10 && f.House.Money >= 40 && f.House.Food <= 50 {
		f.BuyEat()
		return
	}
	if f.Wife.Happiness < 15 && f.House.Money >= 390 { // 350 + 40 - чтобы хватило на еду :)
		f.BuyCoats()
		*totalBuy += 1
		return
	}
	if f.House.Dirty > 90 {
		f.CleanDirty()
		return
	}
}

// checkLifePerson  проверяем сытость и счастливость
func checkLifePerson(f *models.Family) {
	if f.Wife.Happiness < 10 || f.Wife.Satiety <= 0 {
		fmt.Println(f.Wife.Happiness)
		fmt.Println(f.Wife.Satiety)
		panic("Game over, Mary is dead")
	}
	if f.Husband.Happiness < 10 || f.Husband.Satiety <= 0 {
		fmt.Println(f.Husband.Happiness)
		fmt.Println(f.Husband.Satiety)
		panic("Game over, Chel is dead")
	}
}
