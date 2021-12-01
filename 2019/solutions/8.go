package main

import (
    "bufio"
    "fmt"
    "os"
    // "strconv"
    // "strings"
    // "time"
)

const (
	x int = 25
	y int = 6
)

func silver(layers *[][]int) {
	var target,zctemp int = 0,x*y
	for i := range *layers {
		layer := (*layers)[i]
		zcount := 0
		for j := range layer {
			if layer[j] == 0 {
				zcount++
			}
		}
		if zcount < zctemp {
			target = i
			zctemp = zcount
		}
	}
	// fmt.Println(target)
	var one,two int = 0,0
	for i := range (*layers)[target] {
		switch (*layers)[target][i] {
		case 1: 
			one++
		case 2: 
			two++
		default:
			continue
		}
	}
	fmt.Println("Solution 1:",one*two)
}

func gold(layers *[][]int) {
	image := make([]int, x*y)
	for i := range image {
		image[i] = 2
	}
	for i := range *layers {
		layer := (*layers)[i]
		for j := range layer {
			if image[j] == 2 {
				switch layer[j] {
				case 0: // black
					image[j] = layer[j]
				case 1: // white
					image[j] = layer[j]
				case 2: // transparent
					continue
				}
			}
		}
	}
	fmt.Println("Solution 2:")
	for i := range image {
		for j := 1; j <= y; j++ {
			if i == x*j {
				fmt.Println()
			}
		}
		if image[i] == 0 {
			fmt.Print(" ")
		} else {
			fmt.Print(image[i])
		}
	}
}

func main() {
	// change to request the input file from user
    file, _ := os.Open("input")
    defer file.Close()

    var layers [][]int

    scanner := bufio.NewScanner(file)
    scanner.Scan()
	input := []rune(scanner.Text())

	adv := x * y
	for i := 0; i < len(input); i += adv {
		var layer []int
		layer_runes := input[i:i+adv]
		for j := 0; j < len(layer_runes); j++ {
			layer = append(layer, int(layer_runes[j])-48)
		}
		layers = append(layers, layer)
	}
	// fmt.Println(layers)
	silver(&layers)
	gold(&layers)
}	