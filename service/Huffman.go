package services

import (
	"EnchanceSimulator/models"
	"bytes"
	"encoding/binary"
	"io"
)

type HuffmanService struct{}

func makeTree(freqs map[byte]uint32) *models.Node {
	h := &models.MinHeap{}
	for c, f := range freqs {
		h.Push(&models.Node{Char: c, Freq: f})
	}
	for len(h.Arr) > 1 {
		a := h.Pop()
		b := h.Pop()
		h.Push(&models.Node{Char: 0, Freq: a.Freq + b.Freq, Left: a, Right: b})
	}
	return h.Pop()
}

func decode(tree *models.Node, packed []byte, bits uint32) string {
	var unpacked string
	pos := uint32(0)
	var mask byte = 128

	for pos < bits {
		node := tree
		for node.Left != nil || node.Right != nil {
			if pos >= bits {
				panic("out of message bounds")
			}
			b := packed[pos/8] & mask
			if b == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			pos++
			mask >>= 1
			if mask == 0 {
				mask = 128
			}
		}
		unpacked += string(node.Char)
	}
	return unpacked
}

func (s *HuffmanService) ReadUint32(r io.Reader) uint32 {
	var ret uint32
	err := binary.Read(r, binary.LittleEndian, &ret)
	if err != nil {
		return 0
	}
	return ret
}
func readUint32(r io.Reader) uint32 {
	var ret uint32
	err := binary.Read(r, binary.LittleEndian, &ret)
	if err != nil {
		return 0
	}
	return ret
}

func readByte(r io.Reader) byte {
	var ret byte
	err := binary.Read(r, binary.LittleEndian, &ret)
	if err != nil {
		return 0
	}
	return ret
}

func (s *HuffmanService) ReadByte(r io.Reader) byte {
	var ret byte
	err := binary.Read(r, binary.LittleEndian, &ret)
	if err != nil {
		return 0
	}
	return ret
}
func (s *HuffmanService) UnpackFile(inputData []byte) string {
	r := bytes.NewReader(inputData)

	_, _, charsCount := readUint32(r), readUint32(r), readUint32(r)

	freqs := make(map[byte]uint32)
	for i := uint32(0); i < charsCount; i++ {
		count := readUint32(r)
		char := readByte(r)
		freqs[char] = count
	}

	tree := makeTree(freqs)
	packedBits := readUint32(r)
	packedBytes := readUint32(r)
	_ = readUint32(r)

	packed := make([]byte, packedBytes)
	_, err := io.ReadFull(r, packed)
	if err != nil {
		return ""
	}
	unpacked := decode(tree, packed, packedBits)
	return unpacked
}
