package day16

import (
	"sort"
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	hex2Bin := map[rune]string{'0': "0000", '1': "0001", '2': "0010", '3': "0011", '4': "0100", '5': "0101", '6': "0110", '7': "0111", '8': "1000", '9': "1001", 'A': "1010", 'B': "1011", 'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111"}
	binaryString := ""
	for _, r := range strings.TrimSpace(string(*input)) {
		binaryString += hex2Bin[r]
	}
	message := decode([]rune(binaryString))
	if part2 {
		return message.content
	}
	return message.versionSum
}

type packet struct {
	bits       []rune
	index      int
	version    int
	versionSum int
	typeID     int
	content    int
	subpackets []packet
}

func (p *packet) read(c int) int {
	s := p.bits[p.index : p.index+c]
	v, _ := strconv.ParseInt(string(s), 2, 0)
	p.index += c
	return int(v)
}

func decode(bin []rune) packet {
	var p packet
	p.bits = bin
	p.index = 0
	p.version = p.read(3)
	p.typeID = p.read(3)
	switch p.typeID {
	case 4: // literal
		value := 0
		for {
			header := p.read(1)
			value += p.read(4)
			if header == 0 {
				break
			}
		}
		p.content = value
	default:
		lengthTypeID := p.read(1)
		if lengthTypeID == 0 {
			subPacketsLength := p.read(15)
			remainingLength := len(p.bits[p.index:]) - subPacketsLength
			for remainingLength != len(p.bits[p.index:]) {
				p.subpackets = append(p.subpackets, decode(p.bits[p.index:]))
				p.index += p.subpackets[len(p.subpackets)-1].index
			}
		} else { // lengthTypeID == 1
			subPacketsCount := p.read(11)
			for i := 0; i < subPacketsCount; i++ {
				p.subpackets = append(p.subpackets, decode(p.bits[p.index:]))
				p.index += p.subpackets[len(p.subpackets)-1].index
			}
		}
		subpacketContents := []int{}
		for _, sub := range p.subpackets {
			subpacketContents = append(subpacketContents, sub.content)
		}
		if len(subpacketContents) == 1 {
			p.content = subpacketContents[0]
		}
		switch p.typeID {
		case 0: // sum
			value := 0
			for _, c := range subpacketContents {
				value += c
			}
			p.content = value
		case 1: // product
			value := 1
			for _, c := range subpacketContents {
				value *= c
			}
			p.content = value
		case 2: // minimum
			sort.Ints(subpacketContents)
			p.content = subpacketContents[0]
		case 3: // maximum
			sort.Ints(subpacketContents)
			p.content = subpacketContents[len(subpacketContents)-1]
		case 5: // greater than
			if subpacketContents[0] > subpacketContents[1] {
				p.content = 1
			} else {
				p.content = 0
			}
		case 6: // less than
			if subpacketContents[0] < subpacketContents[1] {
				p.content = 1
			} else {
				p.content = 0
			}
		case 7: // equal to
			if subpacketContents[0] == subpacketContents[1] {
				p.content = 1
			} else {
				p.content = 0
			}
		}
	}
	p.versionSum = p.version
	for _, sub := range p.subpackets {
		p.versionSum += sub.versionSum
	}
	return p
}
