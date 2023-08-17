package entities

import "fmt"

type Statistics struct {
	MoneyEarned  uint32
	MoneySpent   uint32
	FoodEaten    uint32
	CatFoodEaten uint32
	CoatBought   uint32
	CompPlayed   uint32
	WallTeared   uint32
}

func (s *Statistics) PrintStatistics() {
	fmt.Println("MoneyEarned: ", s.MoneyEarned)
	fmt.Println("MoneySpent: ", s.MoneySpent)
	fmt.Println("FoodEaten: ", s.FoodEaten)
	fmt.Println("CatFoodEaten: ", s.CatFoodEaten)
	fmt.Println("FurCoatBought: ", s.CoatBought)
	fmt.Println("CompPlayed: ", s.CompPlayed)
	fmt.Println("WallTeared: ", s.WallTeared)
}
