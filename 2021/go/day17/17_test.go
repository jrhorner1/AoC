package day17

import (
	"testing"
)

func Test_maxHeight(t *testing.T) {
	input := "target area: x=20..30, y=-10..-5"
	target := getTarget(input)
	velocityHeights, _ := getVelocities(&target)
	maxHeight := 0
	for _, height := range velocityHeights {
		if height > maxHeight {
			maxHeight = height
		}
	}
	answer := 45
	if maxHeight != answer {
		t.Errorf("got %d, wanted %d", maxHeight, answer)
	}
}

func Test_distinctInitialVelocities(t *testing.T) {
	input := "target area: x=20..30, y=-10..-5"
	target := getTarget(input)
	_, velocities := getVelocities(&target)
	distinct := len(velocities)
	answer := 112
	if distinct != answer {
		t.Errorf("got %d, wanted %d", distinct, answer)
	}
}
