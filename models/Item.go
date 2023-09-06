package models

import "EnchanceSimulator/interfaces"

type SimulatedItem struct {
	marketPrice float64
	cost        float64
	level       int
	grade       string
	itemType    string
	group       string
}

func (s SimulatedItem) MarketPrice() float64 {
	return s.marketPrice
}

func (s SimulatedItem) Cost() float64 {
	return s.cost
}

func (s SimulatedItem) Level() int {
	return s.level
}

func (s SimulatedItem) Grade() string {
	return s.grade
}

func (s SimulatedItem) Type() string {
	return s.itemType
}

func (s SimulatedItem) Group() string {
	return s.group
}

func (s SimulatedItem) SuccesProbability() float64 {
	// Todo
	return 0.0
}

func (s SimulatedItem) Downgrades() bool {
	// Todo
	return false
}

func (s SimulatedItem) Explodes() bool {
	// Todo
	return false
}

func (s SimulatedItem) Cronnable() bool {
	// Todo
	return false
}

func NewSimulatedItem(item interfaces.Item) interfaces.Item {
	return SimulatedItem{
		marketPrice: item.MarketPrice(),
		cost:        item.Cost(),
		level:       item.Level(),
		grade:       item.Grade(),
		itemType:    item.Type(),
		group:       item.Group(),
	}
}
