package day16

import (
	"testing"
)

func Test_literal(t *testing.T) {
	p := Packet{}
	hexString := "D2FE28"
	p.bits = HexToBinary(hexString)
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	if p.typeID != 4 {
		t.Errorf("TypeID 4 mismatch. %d detected.", p.typeID)
	}
	var err error
	p.content, err = p.Literal()
	if err != nil {
		t.Error(err)
	}
	answer := 2021
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}

/* Examples to unit test:
Subpackets:
	Operator w/ length type 0: 38006F45291200
	Operator w/ length type 1: EE00D40C823060
Version number sums:
	8A004A801A8002F478 : 16
	620080001611562C8802118E34 : 12
	C0015000016115A2E0802F182340 : 23
	A0016C880162017C3686B18A3D4780 : 31
*/

func Test_sum(t *testing.T) {
	p := Packet{}
	hexString := "C200B40A82"
	p.bits = HexToBinary(hexString)
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	if p.typeID != 0 {
		t.Errorf("TypeID 0 mismatch. %d detected.", p.typeID)
	}
	p.Subpackets()
	p.content = p.Sum()
	answer := 3
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}

func Test_product(t *testing.T) {
	p := Packet{}
	hexString := "04005AC33890"
	p.bits = HexToBinary(hexString)
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	if p.typeID != 1 {
		t.Errorf("TypeID 1 mismatch. %d detected.", p.typeID)
	}
	p.Subpackets()
	p.content = p.Product()
	answer := 54
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}

func Test_minimum(t *testing.T) {
	p := Packet{}
	hexString := "880086C3E88112"
	p.bits = HexToBinary(hexString)
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	if p.typeID != 2 {
		t.Errorf("TypeID 2 mismatch. %d detected.", p.typeID)
	}
	p.Subpackets()
	p.content = p.Min()
	answer := 7
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}

func Test_maximum(t *testing.T) {
	p := Packet{}
	hexString := "CE00C43D881120"
	p.bits = HexToBinary(hexString)
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	if p.typeID != 3 {
		t.Errorf("TypeID 3 mismatch. %d detected.", p.typeID)
	}
	p.Subpackets()
	p.content = p.Max()
	answer := 9
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}

func Test_less(t *testing.T) {
	p := Packet{}
	hexString := "D8005AC2A8F0"
	p.bits = HexToBinary(hexString)
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	if p.typeID != 6 {
		t.Errorf("TypeID 6 mismatch. %d detected.", p.typeID)
	}
	p.Subpackets()
	p.content = p.Less()
	answer := 1
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}

func Test_greater(t *testing.T) {
	p := Packet{}
	hexString := "F600BC2D8F"
	p.bits = HexToBinary(hexString)
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	if p.typeID != 5 {
		t.Errorf("TypeID 5 mismatch. %d detected.", p.typeID)
	}
	p.Subpackets()
	p.content = p.Greater()
	answer := 0
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}

func Test_equal(t *testing.T) {
	p := Packet{}
	hexString := "9C005AC2F8F0"
	p.bits = HexToBinary(hexString)
	p.index = 0
	p.version, _ = p.Read(3)
	p.typeID, _ = p.Read(3)
	if p.typeID != 7 {
		t.Errorf("TypeID 7 mismatch. %d detected.", p.typeID)
	}
	p.Subpackets()
	p.content = p.Equal()
	answer := 0
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}

func Test_nested(t *testing.T) {
	hexString := "9C0141080250320F1802104A08"
	p := Decode([]rune(HexToBinary(hexString)))
	answer := 1
	if p.content != answer {
		t.Errorf("got %d, wanted %d", p.content, answer)
	}
}
