package main

import "fmt"

const (
	salary              = 150
	coat_cost           = 350
	cleaning_count      = 100
	game_ball           = 20
	min_dirty_limit     = 90
	min_satiety_limit   = 0
	min_happyness_limit = 10
	coat_ball           = 60
	dirty_min           = 5
	eat_min             = 30
	any_action_point    = 10
)

type House struct {
	man     *Man
	woman   *Woman
	money   int
	product int
	dirty   int
}
type Man struct {
	house     *House
	name      string
	satiety   int
	happyness int
}

func (h *House) add_day(day int) {
	h.dirty = h.dirty + day*dirty_min
	if h.man.satiety < 0 {
		fmt.Println(h.man.name, " died from starvation")
	} else if h.woman.satiety < 0 {
		fmt.Println(h.woman.name, " died from starvation")
	} else if h.man.happyness < 10 {
		fmt.Println(h.man.name, " died from depression")
	} else if h.woman.happyness < 10 {
		fmt.Println(h.woman.name, " died from depression")
	}
	if h.dirty > dirty_min {
		h.man.happyness = h.man.happyness - day*any_action_point
		h.woman.happyness = h.woman.happyness - day*any_action_point
	}
}

func (m *Man) eat(count int) { //max 30
	if count > eat_min {
		fmt.Println(" You are cannot eat so much in one time")
	} else if count <= 0 {
		fmt.Println("You are eat nothing")
	} else {
		m.house.product = m.house.product - count
		m.satiety = m.satiety + count
	}

}
func (m *Man) play_game() {
	m.satiety = m.satiety - any_action_point
	m.happyness = m.happyness + game_ball
}
func (m *Man) work() {
	m.house.money = m.house.money + salary
	m.satiety = m.satiety - any_action_point
}

type Woman struct {
	house     *House
	name      string
	satiety   int
	happyness int
}

func (w *Woman) eat(count int) { //max 30
	if count > eat_min {
		fmt.Println(" You are cannot eat so much in one time")
	} else if count <= 0 {
		fmt.Println("You are eat nothing")
	} else {
		w.house.product = w.house.product - count
		w.satiety = w.satiety + count
	}
}
func (w *Woman) buy_product(count int) {
	w.house.product = w.house.product + count
	if w.house.money < count {
		fmt.Println("You dont have enough money. You remain is :", w.house.money)
		w.satiety = w.satiety - any_action_point
	} else {
		w.house.money = w.house.money - count
		w.house.product = w.house.product + count
		w.satiety = w.satiety - any_action_point
	}
}
func (w *Woman) buy_coat() {
	if w.house.money < coat_cost {
		fmt.Println("You dont have enough money. You remain is :", w.house.money)
		w.satiety = w.satiety - any_action_point
	} else {
		w.house.money = w.house.money - coat_cost
		w.happyness = w.happyness + coat_ball
		w.satiety = w.satiety - any_action_point
	}
}
func (w *Woman) clean_home() {
	w.satiety = w.satiety - any_action_point
	if w.house.dirty > min_dirty_limit {
		w.house.dirty = w.house.dirty - cleaning_count
		w.satiety = w.satiety - any_action_point
	} else {
		fmt.Print("Your home is clean, Index cleaness is : ", w.house.dirty)
	}

}
func main() {
	house := House{
		money:   100,
		product: 50,
		dirty:   0,
	}
	man := Man{
		name:      "James",
		satiety:   20,
		happyness: 35,
	}
	woman := Woman{
		name:      "Anna",
		satiety:   25,
		happyness: 15,
	}
	man.work()
	man.play_game()
	man.eat(10)
	woman.eat(8)
	woman.buy_product(5)
	house.add_day(1)
}
