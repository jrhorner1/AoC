package day16

import (
	"math"
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	message := Decode(HexToBinary(string(*input)))
	if part2 {
		return message.content
	}
	return message.versionSum
}

func HexToBinary(input string) []rune {
	hexToBinary := map[rune]string{'0': "0000", '1': "0001", '2': "0010", '3': "0011", '4': "0100", '5': "0101", '6': "0110", '7': "0111", '8': "1000", '9': "1001", 'A': "1010", 'B': "1011", 'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111"}
	binaryString := ""
	for _, r := range strings.TrimSpace(input) {
		binaryString += hexToBinary[r]
	}
	return []rune(binaryString)
}

type Packet struct {
	bits       []rune
	index      int
	version    int
	versionSum int
	typeID     int
	content    int
	subpackets []Packet
}

func (p *Packet) Read(c int) (int, error) {
	v, err := strconv.ParseInt(string(p.Pop(c)), 2, 0)
	return int(v), err
}

func (p *Packet) Pop(c int) []rune {
	s := p.bits[p.index : p.index+c]
	p.index += c
	return s
}

func Decode(bin []rune) Packet {
	var p Packet
	p.bits = bin
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	switch p.typeID {
	case 4: // literal
		p.content, _ = p.Literal()
	default: // Operators
		p.Subpackets()
		switch p.typeID {
		case 0:
			p.content = p.Sum()
		case 1:
			p.content = p.Product()
		case 2:
			p.content = p.Min()
		case 3:
			p.content = p.Max()
		case 5:
			p.content = p.Greater()
		case 6:
			p.content = p.Less()
		case 7:
			p.content = p.Equal()
		}
	}
	p.versionSum = p.version
	for _, sub := range p.subpackets {
		p.versionSum += sub.versionSum
	}
	return p
}

func (p *Packet) Subpackets() {
	lengthTypeID, _ := p.Read(1)
	if lengthTypeID == 0 {
		subPacketsLength, _ := p.Read(15)
		remainingLength := len(p.bits[p.index:]) - subPacketsLength
		for remainingLength < len(p.bits[p.index:]) {
			p.subpackets = append(p.subpackets, Decode(p.bits[p.index:]))
			p.index += p.subpackets[len(p.subpackets)-1].index
		}
	} else { // lengthTypeID == 1
		subPacketsCount, _ := p.Read(11)
		for i := 0; i < subPacketsCount; i++ {
			p.subpackets = append(p.subpackets, Decode(p.bits[p.index:]))
			p.index += p.subpackets[len(p.subpackets)-1].index
		}
	}
}

func (p *Packet) Sum() int {
	value := 0
	for _, sub := range p.subpackets {
		value += sub.content
	}
	return value
}

func (p *Packet) Product() int {
	value := 1
	for _, sub := range p.subpackets {
		value *= sub.content
	}
	return value
}

func (p *Packet) Min() int {
	min := math.MaxInt64
	for _, sub := range p.subpackets {
		if sub.content < min {
			min = sub.content
		}
	}
	return min
}

func (p *Packet) Max() int {
	max := 0
	for _, sub := range p.subpackets {
		if sub.content > max {
			max = sub.content
		}
	}
	return max
}

func (p *Packet) Literal() (int, error) {
	bValue := []rune{}
	for {
		header, _ := p.Read(1)
		bValue = append(bValue, p.Pop(4)...)
		if header == 0 {
			break
		}
	}
	value, err := strconv.ParseInt(string(bValue), 2, 0)
	return int(value), err
}

func (p *Packet) Greater() int {
	if p.subpackets[0].content > p.subpackets[1].content {
		return 1
	}
	return 0
}

func (p *Packet) Less() int {
	if p.subpackets[0].content < p.subpackets[1].content {
		return 1
	}
	return 0
}

func (p *Packet) Equal() int {
	if p.subpackets[0].content == p.subpackets[1].content {
		return 1
	}
	return 0
}
