package day1

import (
	"io/ioutil"
	"testing"
)

func TestExample1(t *testing.T) {
	input, _ := ioutil.ReadFile("2021/examples/1")
	result := Puzzle(&input, false)
	if result != 7 {
		t.Error("some err")
	}
}

func TestExample2(t *testing.T) {
	input, _ := ioutil.ReadFile("2021/examples/1")
	result := Puzzle(&input, true)
	if result != 5 {
		t.Error("some err")
	}
}
