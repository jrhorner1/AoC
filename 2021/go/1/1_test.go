package d121

import (
	"testing"
)

func TestExample1(t *testing.T) {
	example := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	result := DepthScan(example, false)
	if result != 7 {
		t.Error("some err")
	}
}

func TestPart1(t *testing.T) {
	input := Input("2021/input/1")
	result := DepthScan(input, false)
	if result != 1681 {
		t.Error("some err")
	}
}

func TestExample2(t *testing.T) {
	example := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	result := DepthScan(example, true)
	if result != 5 {
		t.Error("some err")
	}
}

func TestPart2(t *testing.T) {
	input := Input("2021/input/1")
	result := DepthScan(input, true)
	if result != 1704 {
		t.Error("some err")
	}
}
