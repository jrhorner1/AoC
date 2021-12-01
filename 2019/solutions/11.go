package main

import (
	"fmt"
	"sort"
	"../../utils"
)

type Panel struct {
	x,y int
	dir int
	color int
}

func paint(start int) *[]Panel {
	intcode := utils.OpenFile("input")
    computer := utils.NewComputer(&intcode)
    go computer.Run()
    input := computer.GetInput()
    // panel starts out at 0,0 (x,y) and all panels are black except the starting panel
    panels := make([]Panel, 0)
    current := Panel{x: 0, y: 0, dir: 0, color: start}
    panels = append(panels, current)
	input <- current.color
    ok := true
    var out int
    count := 0
    for ok {
        out, ok = <-computer.GetOutput()
        if ok {
    	Switch: 
        	switch count % 2 {
	        case 0: // first output
	            switch out { // paint the panel
	            case 0: // black
	            	current.color = 0
	            case 1: // white
	            	current.color = 1
	            }
	            for i, panel := range panels {
	            	if current.x == panel.x && current.y == panel.y { 
	            		panels[i].color = current.color
	            		break
	            	}
	            }
	        case 1: // second output
	            switch out { // direction to turn
	            case 0: // left
	            	switch current.dir {
	            	case 0: // north to west
	            		current.dir = 3
	            	default: 
	            		current.dir--
	            	}
	            case 1: // right
	            	switch current.dir {
	            	case 3: // west to north
	            		current.dir = 0
	            	default: 
	            		current.dir++
	            	}
	            }
	            switch current.dir {
	            case 0: // north
	            	current.y++
	            case 1: // east 
	            	current.x++
	            case 2: // south
	            	current.y--
	            case 3: // west
	            	current.x--
	            }
	            current.color = 0 // reset the panel color to black
	            for _, panel := range panels {
	            	if current.x == panel.x && current.y == panel.y { 
	            		current.color = panel.color // set the panel color if we have already painted it
    					input <- current.color
	            		break Switch
	            	}
	            }
	            panels = append(panels, current)
    			input <- current.color
	        }
        }
        count++
    }
    return &panels
}

func silver() {
	startingPanelColor := 0
	panels := paint(startingPanelColor)
	fmt.Println("Silver:",len(*panels))
}

func gold() {
	startingPanelColor := 1
	panels := paint(startingPanelColor)
	sort.Slice(*panels, func(i, j int) bool {
		if (*panels)[i].y > (*panels)[j].y {
			return true
		}
		if (*panels)[i].y < (*panels)[j].y {
			return false
		}
		return (*panels)[i].x < (*panels)[j].x
	})
	fmt.Println("Gold:")
	newline := 0
	for _, panel := range *panels {
		if panel.y < newline {
			newline--
			fmt.Print("\n")
		}
		if panel.x == 0 {
			continue
		}
		switch panel.color {
		case 0:
			fmt.Print(" ")
		case 1:
			fmt.Print("#")
		}
	}
}

func main() {
	silver()
	gold()
}