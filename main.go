package main

import (
	"fmt"
)

type Person struct {
	Name      string
	Satiation int
	Happiness int
}

func (p *Person) Eat(food *int) {
	if *food >= 30 {
		p.Satiation += 30
		*food -= 30
	} else {
		p.Satiation += *food
		*food = 0
	}
	if p.Satiation > 100 {
		p.Satiation = 100
	}
}

func (p *Person) Work(money *int) {
	p.Satiation -= 10
	*money += 150
}

type House struct {
	Money int
	Food  int
	Dirt  int
}

func (h *House) BuyFood(money *int, food *int) {
	if *money >= 10 {
		*money -= 10
		h.Food += 10
	} else {
		fmt.Println("Недостаточно денег для покупки еды")
	}
}

func (h *House) BuyFurCoat(money *int, wife *Person) {
	if *money >= 350 {
		*money -= 350
		wife.Happiness += 60
	} else {
		fmt.Println("Недостаточно денег для покупки шубы")
	}
}

func (h *House) CleanHouse(wife *Person) {
	if h.Dirt > 0 {
		cleaned := 100
		if h.Dirt < 100 {
			cleaned = h.Dirt
		}
		h.Dirt -= cleaned
		wife.Happiness += 10
	}
}

func main() {
	// Инициализация семьи и дома
	money := 100
	food := 50
	dirt := 0
	Husband := Person{Name: "Муж", Satiation: 30, Happiness: 100}
	Wife := Person{Name: "Жена", Satiation: 30, Happiness: 100}
	home := House{Money: money, Food: food, Dirt: dirt}

	// Моделирование жизни в течение года
	for day := 1; day <= 365; day++ {
		fmt.Printf("День %d:\n", day)

		// Муж и жена едят
		Husband.Eat(&home.Food)
		Wife.Eat(&home.Food)

		// Муж ходит на работу
		Husband.Work(&home.Money)

		// Жена покупает продукты
		home.BuyFood(&home.Money, &home.Food)

		// Жена убирается в доме
		home.CleanHouse(&Wife)

		// Грязь добавляется каждый день
		home.Dirt += 5

		// Проверка на степень счастья и сытости
		if home.Dirt > 90 {
			Husband.Happiness -= 10
			Wife.Happiness -= 10
		}

		if Husband.Satiation <= 0 || Wife.Satiation <= 0 {
			fmt.Println("Кто-то умер от голода")
			break
		}

		if Husband.Happiness <= 10 || Wife.Happiness <= 10 {
			fmt.Println("Кто-то умер от депрессии")
			break
		}
	}

	// Подвести итоги
	fmt.Printf("Деньги заработано: %d\n", home.Money-100)
	fmt.Printf("Съедено еды: %d\n", 50-home.Food)
	fmt.Printf("Куплено шуб: %d\n", 0)
}
