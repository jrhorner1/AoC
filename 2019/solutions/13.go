package main

import (
	"fmt"
	"../../utils"
)

type Tile struct {
	x,y int
	id int
}

func silver() {
	intcode := utils.OpenFile("input")
    computer := utils.NewComputer(&intcode)
    go computer.Run()
    ok := true
    var out int
    var grid []Tile
    var tile Tile
    count,ans := 4,0
    for ok {
        out, ok = <-computer.GetOutput()
        switch count % 3 {
        case 0:
	        switch out {
	        case 0:
	        	tile.id = out
	        case 1:
	        	tile.id = out
	        case 2:
	        	ans++
	        	tile.id = out
	        case 3: 
	        	tile.id = out
	        case 4:
	        	tile.id = out
	        }
	        grid = append(grid, tile)
	    case 1:
	    	tile.x = out
	    case 2:
	    	tile.y = out
        }
        count++
    }
    fmt.Println("Silver:",ans)
}

func gold() {
	intcode := utils.OpenFile("input")
    computer := utils.NewComputer(&intcode)
    memory := computer.GetMemory()
    (*memory)[0] = 2
    go computer.Run()
    ok := true
    var out int
    var grid []Tile
    var tile Tile
    var score bool = false
    count,ans := 4,0
    input := computer.GetInput()
    for ok {
        out, ok = <-computer.GetOutput()
        switch count % 3 {
        case 0:
        	if score {
        		fmt.Println("Score:",out)
        		score = false
        		continue
        	}
	        switch out {
	        case 0:
	        	tile.id = out
	        case 1:
	        	tile.id = out
	        case 2:
	        	ans++
	        	tile.id = out
	        case 3: 
	        	tile.id = out
	        case 4:
	        	tile.id = out
	        }
	        grid = append(grid, tile)
	    case 1:
	    	switch out {
	    	case -1:
	    		score = true
	    	default:
	    		tile.x = out
	    	}
	    case 2:
        	if score {
        		continue
        	}
	    	tile.y = out
        }
        count++
        input <- 0
    }
}

func main () {
	silver()
	gold()
}