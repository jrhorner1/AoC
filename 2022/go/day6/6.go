package day6

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type runeSlice []rune

func (c *runeSlice) push(r rune) { *c = append(*c, r) }
func (c *runeSlice) pull() rune  { r := (*c)[0]; *c = (*c)[1:]; return r }
func (c *runeSlice) count(r rune) int {
	n := 0
	for i := range *c {
		if (*c)[i] == r {
			n++
		}
	}
	return n
}

func (c *runeSlice) unique() bool {
	for _, r := range *c {
		if (*c).count(r) > 1 {
			return false
		}
	}
	return true
}

func Puzzle(input *[]byte, markerLen int) int {
	//logrus.SetLevel(logrus.DebugLevel)
	marker := runeSlice{}
	for i, char := range strings.TrimSpace(string(*input)) {
		marker.push(char)
		if i < markerLen-1 { // account for zero index
			continue
		}
		if marker.unique() {
			logrus.Debug(string(marker), i+1)
			return i + 1 // account for zero index
		}
		marker.pull()
	}
	return -1
}
