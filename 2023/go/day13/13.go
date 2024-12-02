package day13

import (
	"image"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Pattern struct {
	X, Y  int
	Rocks map[image.Point]any
}

func Puzzle(input *[]byte, part2 bool) int {
	log.Debug("Insert Code Here")
	notes := strings.Split(strings.TrimSpace(string(*input)), "\n\n")

	patterns := []Pattern{}
	for _, note := range notes {
		rows := strings.Split(strings.TrimSpace(note), "\n")
		pattern := Pattern{0, len(rows), make(map[image.Point]any)}
		for y, row := range rows {
			pattern.X = len(row)
			for x, column := range row {
				if column == '#' { // its a rock
					rock := image.Point{x, y}
					pattern.Rocks[rock] = nil
				}
			}
		}
		patterns = append(patterns, pattern)
	}
	log.Debug(patterns)
	reflections := []int{}
	for _, pattern := range patterns {
		if y := pattern.findMirror("Y"); y != 0 {
			log.Debugf("Mirrored on row: %d", y)
			reflections = append(reflections, y*100)
			continue
		}
		if x := pattern.findMirror("X"); x != 0 {
			log.Debugf("Mirrored on column: %d", x)
			reflections = append(reflections, x)
		}
	}
	log.Debug("Reflections: ", reflections)
	sum := 0
	for _, v := range reflections {
		sum += v
	}
	if part2 {
		return 42
	}
	return sum
}

func (p *Pattern) findMirror(XY string) int {
checkPlane:
	for x := 1; x < p.X; x++ {
		for rock := range p.Rocks {
			if rock.X == x {
				inverse := rock
				inverse.X -= 1
				if _, ok := p.Rocks[inverse]; !ok {
					log.Debugf("Column: %d Rock: %v %v", x, rock, inverse)
					continue checkPlane
				}
			}
			if rock.X > x {
				inverse := rock
				inverse.X = x - (rock.X - x)
				if inverse.X < 0 {
					continue
				}
				if _, ok := p.Rocks[inverse]; !ok {
					log.Debugf("Column: %d Rock: %v %v", x, rock, inverse)
					continue checkPlane
				}
			}
			if rock.X < x {
				inverse := rock
				inverse.X = x + (x - rock.X)
				if x == 1 {
					inverse.X = 1
				}
				if inverse.X >= p.X {
					continue
				}
				if _, ok := p.Rocks[inverse]; !ok {
					log.Debugf("Column: %d Rock: %v %v", x, rock, inverse)
					continue checkPlane
				}
			}
		}
		return x
	}
	return 0
}
