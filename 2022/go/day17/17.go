package day17

import (
	"image"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	formations = [5]int{
		(0b0011110 << 24) | (0b0000000 << 16) | (0b0000000 << 8) | (0b0000000), // horizontal line
		(0b0001000 << 24) | (0b0011100 << 16) | (0b0001000 << 8) | (0b0000000), // plus sign
		(0b0011100 << 24) | (0b0000100 << 16) | (0b0000100 << 8) | (0b0000000), // backwards L
		(0b0010000 << 24) | (0b0010000 << 16) | (0b0010000 << 8) | (0b0010000), // vertical line
		(0b0011000 << 24) | (0b0011000 << 16) | (0b0000000 << 8) | (0b0000000), // square
	}
)

func Puzzle(input *[]byte, total int) int {
	logrus.SetLevel(logrus.DebugLevel)
	jets := strings.TrimSpace(string(*input))
	chamber := &Chamber{}
	floor := [7]image.Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}}
	limit := total
	count := 0
	rock := Rock{formations[count%5], 3}
loop:
	for {
		for _, jet := range jets {
			rock.push(jet, chamber)
			if !rock.fall(chamber) {
				count++
				if count >= limit {
					break loop
				}
				rock = Rock{formations[count%5], 3 + chamber.maxY()}
			}
		}
		chamber.floor(&floor)
		logrus.Debugf("Total Rocks: %d|%v", count, floor)
	}
	chamber.floor(&floor)
	logrus.Debugf("Total Rocks: %d|%v", count, floor)
	return chamber.maxY()
}

type Chamber []byte

func (c *Chamber) maxY() int { return len(*c) }
func (c *Chamber) floor(floor *[7]image.Point) {
	for x := 0; x < 7; x++ {
		xbit := byte(1 << x)
		for i := c.maxY() - 1; i >= 0; i-- {
			if (*c)[i]&xbit != 0 && i > floor[x].Y {
				floor[x].Y = i + 1
				break
			}
		}
	}
}

type Rock [2]int

func (r *Rock) row(i int) byte { return byte(((*r)[0] >> ((3 - i) * 8)) & 0xff) }
func (r *Rock) extent() byte   { return byte(r.row(0) | r.row(1) | r.row(2) | r.row(3)) }

func (r *Rock) push(direction rune, c *Chamber) {
	switch direction {
	case '>':
		if r.extent()&1 == 0 {
			(*r)[0] >>= 1
			if r.overlap(c) {
				(*r)[0] <<= 1
			}
		}
	case '<':
		if r.extent()&0b1000000 == 0 {
			(*r)[0] <<= 1
			if r.overlap(c) {
				(*r)[0] >>= 1
			}
		}
	}
}

func (r *Rock) fall(c *Chamber) bool {
	(*r)[1]--               // move rock down 1
	if (*r)[1] > c.maxY() { // if rock Y higher than max Y continue
		return true
	} else if (*r)[1] < 0 { // stop at the bottom of the chamber
		r.stop(c)
		return false
	} else if r.overlap(c) { // check if rock overlaps with anything in chamber
		r.stop(c)
		return false
	} // if no overlap, movement succeeds
	return true
}

func (r *Rock) stop(c *Chamber) {
	(*r)[1]++
	for i := 0; i < 4; i++ {
		y := (*r)[1] + i
		if r.row(i) != 0 {
			if y < c.maxY() {
				(*c)[y] |= r.row(i)
			} else {
				*c = append(*c, r.row(i))
			}
		}
	}
}

func (r *Rock) overlap(c *Chamber) bool {
	// only need to check the rocks 4 Y values
	for i := 0; i < 4; i++ {
		y := (*r)[1] + i
		if r.row(i) == 0 || y >= c.maxY() {
			continue
		} else if r.row(i)&(*c)[y] != 0 {
			return true
		}
	}
	return false
}
