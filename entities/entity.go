package entities

import "fmt"

type Entity struct {
	Name           string
	Satiety        int
	IsAlive        bool
	toDo           string
	houseResources *Resources
}

func (e *Entity) Init(name string, houseResources *Resources) {
	e.Name = name
	e.Satiety = 30
	e.IsAlive = true
	e.houseResources = houseResources
}

func (e *Entity) Check() {
	if e.Satiety <= 0 {
		e.IsAlive = false
	}
}

func (e *Entity) GetData() {
	fmt.Printf("[%v]\n", e.Name)
	fmt.Println("	Today's Business: ", e.toDo)
	fmt.Println("	Satiety: ", e.Satiety)
	fmt.Println("	IsAlive: ", e.IsAlive)
}
