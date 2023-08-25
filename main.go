package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stats struct {
	totalMoney uint32
	totalFood  uint32
	totalCoats uint32
}

type House struct {
	Money uint16
	Food  uint16
	Dirt  uint16
}

type Person struct {
	Name      string
	Fullness  int8
	Happiness uint8
}

const (
	// house
	defaultMoney   = 100
	defaultFood    = 50
	defaultDirt    = 0
	plusMoney      = 150
	FurCoatPrice   = 350
	foodPrice      = 10
	plusDirt       = 5
	maxMinusDirt   = 100
	thresholdDirt  = 90
	minusHappiness = 10
	// person
	defaultFullness  = 30
	maxHappiness     = 100
	minHappiness     = 10
	hasbandHappiness = 20
	wifeHappiness    = 60
	minusFullness    = 10
	maxPlusFullness  = 30
	minFullness      = 0
	maxFullness      = 100
	// simulation
	DaysInYear = 365
)

// husband

func (h *Person) play() {
	if h.Happiness < maxHappiness {
		h.Happiness += hasbandHappiness
		h.Fullness -= minusFullness
	}
}

func (h *Person) work(house *House, stats *Stats) {
	house.Money += plusMoney
	h.Fullness -= minusFullness
	stats.totalMoney += plusMoney
}

// wife

func (w *Person) buyFood(house *House) {
	if house.Money >= 10 {
		house.Money -= 10
		house.Food += 10
		w.Fullness -= minusFullness
	}
}

func (w *Person) buyFurCoat(house *House, stats *Stats) {
	if w.Happiness < maxHappiness {
		if house.Money >= FurCoatPrice {
			house.Money -= FurCoatPrice
			stats.totalCoats += 1
			w.Happiness += wifeHappiness
			w.Fullness -= minusFullness
			if w.Happiness > maxHappiness {
				w.Happiness = maxHappiness
			}
		}
	}
}

func (w *Person) cleanHouse(house *House, amount uint16) {
	if house.Dirt != 0 && amount <= maxMinusDirt {
		house.Dirt -= amount
		w.Fullness -= minusFullness
	}
}

// both

func (p *Person) eat(house *House, amount uint16) {
	if house.Food >= amount && amount <= maxPlusFullness {
		house.Food -= amount
		p.Fullness += int8(amount)
	}
}

func main() {
	house := House{
		Money: defaultMoney,
		Food:  defaultFood,
		Dirt:  defaultDirt,
	}

	husband := Person{
		Name:      "Husband",
		Fullness:  defaultFullness,
		Happiness: maxHappiness,
	}

	wife := Person{
		Name:      "Wife",
		Fullness:  defaultFullness,
		Happiness: maxHappiness,
	}

	// Create an instance of Stats to track the simulation's progress
	stats := Stats{
		totalMoney: 0,
		totalFood:  0,
		totalCoats: 0,
	}

	rand.Seed(time.Now().Unix())

	husbandActions := []func(){
		husband.play,
		func() { husband.work(&house, &stats) },
		func() { husband.eat(&house, uint16(rand.Intn(int(maxFullness)+1))) }}
	wifeActions := []func(){
		func() { wife.buyFood(&house) },
		func() { wife.buyFurCoat(&house, &stats) },
		func() { wife.cleanHouse(&house, uint16(rand.Intn(int(maxMinusDirt)+1))) },
		func() { wife.eat(&house, uint16(rand.Intn(int(maxFullness)+1))) }}

	for day := 0; day < DaysInYear; day++ {

		husbandAction := rand.Intn(3) // 0: play, 1: work, 2: eat

		wifeAction := rand.Intn(4) // 0: buyFood, 1: buyFurCoat, 2: cleanHouse, 3: eat

		husbandActions[husbandAction]()
		wifeActions[wifeAction]()
		if house.Dirt > thresholdDirt {
			husband.Happiness -= minusHappiness
			wife.Happiness -= minusHappiness
		}

		if (husband.Fullness < minFullness || wife.Fullness < minFullness) || (husband.Happiness < minHappiness || wife.Happiness < minHappiness) {
			fmt.Println("Game over!")
			fmt.Printf("Степень счастья мужа: %d\n", husband.Happiness)
			fmt.Printf("Степень счастья жены: %d\n", wife.Happiness)
			fmt.Printf("Степень сытости мужа: %d\n", husband.Fullness)
			fmt.Printf("Степень сытости жены: %d\n", wife.Fullness)
			fmt.Printf("Все деньги: %d\n", stats.totalMoney)
			fmt.Printf("Вся еда: %d\n", stats.totalFood)
			fmt.Printf("Все шубы: %d\n", stats.totalCoats)
			break
		}

		if day == DaysInYear-1 {
			fmt.Println("Game over!")
			fmt.Printf("Степень счастья мужа: %d\n", husband.Happiness)
			fmt.Printf("Степень счастья жены: %d\n", wife.Happiness)
			fmt.Printf("Степень сытости мужа: %d\n", husband.Fullness)
			fmt.Printf("Степень сытости жены: %d\n", wife.Fullness)
			fmt.Printf("Все деньги: %d\n", stats.totalMoney)
			fmt.Printf("Вся еда: %d\n", stats.totalFood)
			fmt.Printf("Все шубы: %d\n", stats.totalCoats)
			break
		}
	}
}
