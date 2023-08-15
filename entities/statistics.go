package entities

import "fmt"

type Statistics struct {
	MoneyEarned   int
	MoneySpent    int
	FoodEaten     int
	CatFoodEaten  int
	FurCoatBought int
	CompPlayed    int
	WallTeared    int
}

func (s *Statistics) GetStatistics() {
	fmt.Println("MoneyEarned: ", s.MoneyEarned)
	fmt.Println("MoneySpent: ", s.MoneySpent)
	fmt.Println("FoodEaten: ", s.FoodEaten)
	fmt.Println("CatFoodEaten: ", s.CatFoodEaten)
	fmt.Println("FurCoatBought: ", s.FurCoatBought)
	fmt.Println("CompPlayed: ", s.CompPlayed)
	fmt.Println("WallTeared: ", s.WallTeared)
}
