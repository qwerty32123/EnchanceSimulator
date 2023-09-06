package interfaces

// Item interfaces defines the properties and methods for an item.
// Main item that gonna be use to run all the simulations enchancing simulations
type Item interface {
	MarketPrice() float64       // Method to get the market price of the item.
	Cost() float64              // Method to get the cost of the item.
	Level() int                 // Method to get the level of the item (from 0 to 20).
	Grade() string              // Method to get the grade of the item (e.g., Blue, Green, Yellow, Orange).
	Type() string               // Method to get the type of the item.
	Group() string              // Method to get the group of the item.
	SuccesProbability() float64 // Method that returns succes probability till next stage, based on Grade, Type and Level
	Downgrades() bool           //Method that tell is if the item can downgrade, based on Grade  Type and level
	Explodes() bool             // Method that tells us if the item explodes if not backed up with crons Grade Type and
	Cronnable() bool            //Method that tells us if the item explodes if not backed up with crons Grade Type and level
}
