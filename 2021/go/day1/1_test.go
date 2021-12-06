package day1

import (
	"testing"
)

func TestExample1(t *testing.T) {
	input := Input("2021/examples/1")
	result := DepthScan(input, false)
	if result != 7 {
		t.Error("some err")
	}
}

func TestExample2(t *testing.T) {
	input := Input("2021/examples/1")
	result := DepthScan(input, true)
	if result != 5 {
		t.Error("some err")
	}
}
