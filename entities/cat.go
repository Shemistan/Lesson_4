package entities

import "math/rand"

type Cat struct {
	Entity
}

func (c *Cat) Eat(amount int) {
	if c.houseResources.Food < amount {
		amount = c.houseResources.Food
	}
	c.Satiety += amount * 2
	c.houseResources.CatFoodOutcome(amount)
}

func (c *Cat) Sleep() {
	c.Satiety -= 10
}

func (c *Cat) TearWall() {
	c.Satiety -= 10
	c.houseResources.DirtIncome(5)
}

func (c *Cat) CatLogic() {
	c.Check()

	if c.IsAlive {
		if c.Satiety < 90 && c.houseResources.CatFood > 0 {
			c.Eat(10)
			c.toDo = "Eat"
		} else {
			switch rand.Intn(2) {
			case 0:
				{
					c.Sleep()
					c.toDo = "Sleep"
				}
			case 1:
				{
					c.TearWall()
					c.toDo = "TearWall"
				}
			}
		}
	}
}
