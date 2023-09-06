package models

type Node struct {
	Char  byte
	Freq  uint32
	Left  *Node
	Right *Node
}
