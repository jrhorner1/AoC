package day9

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/jrhorner1/AoC/pkg/intcode"
)

func TestExample1(t *testing.T) {
	filename := "2019/examples/9.1"
	program := Input(filename)
	computer := intcode.NewComputer(&program)
	go computer.Run()

	ok := true
	var out int
	var example1 []int
	for ok {
		out, ok = <-computer.GetOutput()
		example1 = append(example1, out)
	}

	correct := true
	if len(example1) == len(program) {
		for i := range example1 {
			if example1[i] != program[i] {
				correct = false
			}
		}
	}
	if correct {
		fmt.Println("Example 1: YEET!")
	}
}

func TestExample2(t *testing.T) {
	filename := "2019/examples/9.2"
	program := Input(filename)
	computer := intcode.NewComputer(&program)
	go computer.Run()
	var example2 int
	example2 = <-computer.GetOutput()
	correct := true
	if len(strconv.Itoa(example2)) != 16 {
		correct = false
	}
	if correct {
		fmt.Println("Example 2: YEET!")
	}
}

func TestExample3(t *testing.T) {
	filename := "2019/examples/9.3"
	program := Input(filename)
	computer := intcode.NewComputer(&program)
	go computer.Run()
	var example3 int
	example3 = <-computer.GetOutput()
	correct := true
	if example3 != program[1] {
		correct = false
	}
	if correct {
		fmt.Println("Example 3: YEET!")
	}
}
