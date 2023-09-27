package service

import (
	"fmt"

	"github.com/Shemistan/Lesson_4/models"
)

const (
	daysOfLife = 365

	timeToEat     = models.LimitSatietyToDead + models.DecrSatiety + 1
	timeToFun     = models.HappinessLimitToDead + 2*models.DecHappinessByDirty + 1
	timeToFeedCat = models.CatlimitSatietyToDead + models.DecrCatSatiety + 1
)

func TryNotToDie(family *models.Family) *models.TotalCount {
	currentResult := &models.TotalCount{
		EarnedMoney: 0,
		EatenFood:   0,
		BoughtFur:   0,
		DeadPerson:  make([]string, 0, 3),
	}

	curCatAction := 0 // Пусть кот выполняет действия циклично -> Поел, поспал, погадил (как собственно они и поступают обычно)

	for i := 0; i < daysOfLife; i++ {
		printDayStat(family, i+1)

		checkCreatureStatus(family, currentResult, i)

		anotherDayOfLife(family, currentResult, curCatAction)

		curCatAction = (curCatAction + 1) % 3
	}

	return currentResult
}

func checkCreatureStatus(f *models.Family, result *models.TotalCount, day int) {
	if !isAddedToList(result.DeadPerson, "husband") && !f.IsAliveHusband() {
		result.DeadPerson = append(result.DeadPerson, "husband")
		fmt.Printf("Муж умер на  %d день\n", day)
	}
	if !isAddedToList(result.DeadPerson, "wife") && !f.IsAliveWife() {
		result.DeadPerson = append(result.DeadPerson, "wife")
		fmt.Printf("Жена умерла на  %d день\n", day)

	}
	if !isAddedToList(result.DeadPerson, "cat") && !f.IsAliveCat() {
		result.DeadPerson = append(result.DeadPerson, "cat")
		fmt.Printf("Кот умер на  %d день\n", day)

	}
}

func isAddedToList(deadList []string, creature string) bool {
	for i := 0; i < len(deadList); i++ {
		if deadList[i] == creature {
			return true
		}
	}

	return false
}

func anotherDayOfLife(f *models.Family, result *models.TotalCount, curCatAction int) {
	f.DirtyEveryDay()

	if f.IsAliveHusband() {
		husbandAction(f, result)

	}

	if f.IsAliveWife() {
		wifeAction(f, result)
	}

	if f.IsAliveCat() {
		catAction(f, result, curCatAction)
	}
}

func husbandAction(f *models.Family, r *models.TotalCount) {
	if f.Husband.Satiety < timeToEat {
		f.EatFood(true)
		r.EatenFood += models.IncSatietyLimit
	} else if f.Husband.Happiness < timeToFun {
		f.PlayPC()
	} else {
		f.GoToWork()
		r.EarnedMoney += models.Salary
	}
}

func wifeAction(f *models.Family, r *models.TotalCount) {
	if f.Wife.Satiety < timeToEat {
		f.EatFood(false)
		r.EatenFood += models.IncSatietyLimit
	} else if f.Wife.Happiness < timeToFun {
		f.BuyFur()
		r.BoughtFur++
	} else if f.Home.FoodInFridge < 2*models.IncSatietyLimit {
		f.BuyProducts()
	} else if f.Home.CatFood < models.CatEatLimit {
		f.BuyCatFood()
	} else {
		f.CleanHome()
	}

}

func catAction(f *models.Family, r *models.TotalCount, action int) {
	switch action {
	case 0:
		f.TearUpWalpaper()
	case 1:
		f.Sleep()
	case 2:
		f.EatFoodByCat()
	}
}

func printDayStat(f *models.Family, day int) {
	fmt.Printf("\nНаступил %d день: \n", day)
	fmt.Printf("Статистика мужа: Счастье %d единиц, Голод %d единиц \n", f.Husband.Happiness, f.Husband.Satiety)
	fmt.Printf("Статистика Жены: Счастье %d единиц, Голод %d единиц \n", f.Wife.Happiness, f.Wife.Satiety)
	fmt.Printf("Статистика Кота: Голод %d единиц  \n", f.Cat.Satiety)
	fmt.Printf("Статистика Дома: Денег %d единиц, Еды %d единиц, Грязь %d единиц, Кошачьей еды %d единиц \n", f.Home.MoneyInNightstand, f.Home.FoodInFridge, f.Home.DirtyInHouse, f.Home.CatFood)

}
