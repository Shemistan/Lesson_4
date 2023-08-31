package home

import (
	"fmt"
)

const (
	Dirtyness     = 90
	IncreaseDirty = 5
)

type Home struct {
	Money, Dust, Food int
}

func (h *Home) AddDirtiness() {
	h.Dust += IncreaseDirty
	fmt.Printf("Dust in the house: %d\n", h.Dust)
}

func (h *Home) DecreaseHappiness() {

	if h.Dust > Dirtyness {
	}
	fmt.Printf("Dust in the house is more than: %d\n and family happiness is decreased", h.Dust)
}
