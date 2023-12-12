package day10

import (
	"image"
	"slices"
	"strings"

	log "github.com/sirupsen/logrus"
)

var deltas = map[string]image.Point{
	"left":  {X: -1, Y: 0},
	"right": {X: 1, Y: 0},
	"up":    {X: 0, Y: -1},
	"down":  {X: 0, Y: 1},
}

type pipe struct {
	shape    rune
	distance int
}

func Puzzle(input *[]byte, part2 bool) int {
	// parse the sketch of pipes
	pipes := map[image.Point]pipe{}
	start := image.Point{}
	for y, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		for x, r := range line {
			switch r {
			case '.': // not a pipe, do nothing
			case 'S': // starting point, can't determine shape yet
				start = image.Point{x, y}
			default:
				pipes[image.Point{x, y}] = pipe{r, 0}
			}
		}
	}
	pipes[start] = pipe{shapeOf(start, &pipes), 0}
	log.Debugf("Starting pipe shape: %c", pipes[start].shape)
	// determine distances from start
	a, b := paths(start, &pipes)
	ap, bp := start, start
	for dist := 1; a != b; dist++ {
		next(&a, &ap, dist, &pipes)
		log.Debugf("AP: %s, {%c, %d}", ap, pipes[ap].shape, pipes[ap].distance)
		next(&b, &bp, dist, &pipes)
		log.Debugf("BP: %s, {%c, %d}", bp, pipes[bp].shape, pipes[bp].distance)
	}
	next(&a, &ap, pipes[ap].distance+1, &pipes)
	log.Debugf("A: %s, {%c, %d}", ap, pipes[a].shape, pipes[a].distance)
	next(&b, &bp, pipes[bp].distance+1, &pipes)
	log.Debugf("B: %s, {%c, %d}", bp, pipes[b].shape, pipes[b].distance)
	// determine the farthest pipe
	farthest := 0
	for _, pipe := range pipes {
		if pipe.distance > farthest {
			farthest = pipe.distance
		}
	}
	if part2 {
		return 42
	}
	return farthest
}

func shapeOf(start image.Point, pipes *map[image.Point]pipe) rune {
	// determine shape of starting pipe
	potentialShapes := []rune{'-', '7', 'F', '|', 'J', 'L'}
	for direction, delta := range deltas {
		adj := start.Add(delta)
		log.Debugf("Checking %s delta: %s", direction, delta)
		if p, ok := (*pipes)[adj]; ok {
			log.Debugf("There's a pipe here! %c", p.shape)
			switch direction {
			case "left":
				// []rune{'-', '7', 'J'}
				if p.shape == '-' || p.shape == 'F' || p.shape == 'L' {
					prune(&potentialShapes, []rune{'|', 'F', 'L'})
				}
			case "right":
				// []rune{'-', 'F', 'L'}
				if p.shape == '-' || p.shape == 'J' || p.shape == '7' {
					prune(&potentialShapes, []rune{'|', 'J', '7'})
				}
			case "up":
				// []rune{'|', 'J', 'L'}
				if p.shape == '|' || p.shape == 'F' || p.shape == '7' {
					prune(&potentialShapes, []rune{'-', 'F', '7'})
				}
			case "down":
				// []rune{'|', '7', 'F'}
				if p.shape == '|' || p.shape == 'J' || p.shape == 'L' {
					prune(&potentialShapes, []rune{'-', 'J', 'L'})
				}
			}
		}
	}
	if len(potentialShapes) != 1 {
		log.Fatal("Too many potential shapes! ", potentialShapes)
	}
	return potentialShapes[0]
}

func prune(ps *[]rune, s []rune) {
	for _, r := range s {
		i := slices.Index(*ps, r)
		if i != -1 {
			log.Debugf("Deleting index %d: %c", i, (*ps)[i])
			*ps = slices.Delete(*ps, i, i+1)
		}
	}
}

func paths(pos image.Point, pipes *map[image.Point]pipe) (a, b image.Point) {
	pipe := (*pipes)[pos]
	switch pipe.shape {
	case '-':
		a = pos.Add(deltas["left"])
		b = pos.Add(deltas["right"])
	case '|':
		a = pos.Add(deltas["up"])
		b = pos.Add(deltas["down"])
	case 'F':
		a = pos.Add(deltas["down"])
		b = pos.Add(deltas["right"])
	case '7':
		a = pos.Add(deltas["left"])
		b = pos.Add(deltas["down"])
	case 'L':
		a = pos.Add(deltas["up"])
		b = pos.Add(deltas["right"])
	case 'J':
		a = pos.Add(deltas["left"])
		b = pos.Add(deltas["up"])
	}
	return a, b
}

func next(pos, prev *image.Point, dist int, pipes *map[image.Point]pipe) {
	pipe := (*pipes)[*pos]
	pipe.distance = dist
	log.Debug("pipe", pipe)
	(*pipes)[*pos] = pipe
	switch pipe.shape {
	case '-':
		checkPath(pos, prev, "left", "right")
	case '|':
		checkPath(pos, prev, "up", "down")
	case 'F':
		checkPath(pos, prev, "down", "right")
	case '7':
		checkPath(pos, prev, "left", "down")
	case 'L':
		checkPath(pos, prev, "up", "right")
	case 'J':
		checkPath(pos, prev, "left", "up")
	}
}

func checkPath(pos, prev *image.Point, n, p string) {
	next := pos.Add(deltas[n])
	if next != *prev {
		*prev = *pos
		*pos = next
		return
	}
	*prev = *pos
	*pos = pos.Add(deltas[p])
}
