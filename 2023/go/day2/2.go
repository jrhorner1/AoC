package day2

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type cube struct {
	color    string
	quantity int
}

func Puzzle(input *[]byte, part2 bool) int {
	//log.SetLevel(log.DebugLevel)
	totalCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var possible, powers []int
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		content := strings.Split(strings.TrimSpace(line), ":")
		gameID, err := strconv.Atoi(strings.Split(strings.TrimSpace(content[0]), " ")[1])
		if err != nil {
			log.Error(err)
		}
		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		impossible := false
		sets := strings.Split(strings.TrimSpace(content[1]), ";")
		for _, set := range sets {
			for _, spec := range strings.Split(strings.TrimSpace(set), ",") {
				value := strings.Split(strings.TrimSpace(spec), " ")
				cube := cube{
					quantity: 0,
					color:    value[1],
				}
				cube.quantity, err = strconv.Atoi(value[0])
				if err != nil {
					log.Error(err)
				}
				if !impossible && cube.quantity > totalCubes[cube.color] {
					log.Debug("Game ", gameID, " is impossible.")
					impossible = true
				}
				if minCubes[cube.color] < cube.quantity {
					minCubes[cube.color] = cube.quantity
				}
			}
		}
		if !impossible {
			possible = append(possible, gameID)
		}
		log.Debug("Game ", gameID, " minimum cubes required: red=", minCubes["red"], " green=", minCubes["green"], " blue=", minCubes["blue"])
		power := 1
		for _, v := range minCubes {
			power *= v
		}
		powers = append(powers, power)
		log.Debug("Game ", gameID, " power: ", power)
	}
	sum := 0
	for _, id := range possible {
		sum += id
	}
	if part2 {
		sum = 0
		for _, power := range powers {
			sum += power
		}
	}
	return sum
}
