package interfaces

import (
	"EnchanceSimulator/models"
	"io"
)

type HuffmanCoder interface {
	MakeTree(freqs map[byte]uint32) *models.Node
	Decode(tree *models.Node, packed []byte, bits uint32) string
	ReadUint32(r io.Reader) uint32
	ReadByte(r io.Reader) byte
}
