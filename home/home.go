package home

import "fmt"

type Home struct {
	Money, Dust, Food int
}

func (h *Home) AddDirtiness() {
	h.Dust += 5
	fmt.Printf("Dust in the house: %d\n", h.Dust)
}
