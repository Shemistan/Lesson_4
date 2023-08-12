package main

import (
	"github.com/Shemistan/Lesson_4/home"
)

type Stats struct {
	EarnedMoney int64
	EatenFood   int64
	BoughtCoats int32
}

func main() {
	stats := Stats{
		EarnedMoney: 0,
		EatenFood:   0,
		BoughtCoats: 0,
	}

	home := home.Factory("John", "Marry")

	year(&home, &stats)
}

func year(home *home.Home, stats *Stats) {
	for i := 0; i < 365; i++ {
		home.BuyCoat()
	}
}
