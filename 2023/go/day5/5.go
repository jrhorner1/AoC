package day5

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Category struct {
	id                   string
	dest, src, len, diff []int
	prev, next           *Category
}

func (c Category) GetLocation(n int) int {
	o := n
	for i := range c.src {
		if n >= c.src[i] && n < c.src[i]+c.len[i] {
			o += c.diff[i]
			break
		}
	}
	log.Debug("GetLocation | Category: ", c.id, " Input: ", n, " Output: ", o)
	if c.next != nil {
		return c.next.GetLocation(o)
	}
	return o
}

func (c Category) GetSeed(n int) int {
	o := n
	for i := range c.dest {
		if n >= c.dest[i] && n < c.dest[i]+c.len[i] {
			o -= c.diff[i]
			break
		}
	}
	if c.prev != nil {
		return c.prev.GetSeed(o)
	}
	return o
}

func Puzzle(input *[]byte, part2 bool) int {
	//log.SetLevel(log.DebugLevel)
	categories := strings.Split(strings.TrimSpace(string(*input)), "\n\n")
	var seeds []int
	almanac := []Category{}
	for i, category := range categories {
		if i == 0 {
			parsed := strings.Split(category, " ")[1:]
			for _, elem := range parsed {
				seed, err := strconv.Atoi(elem)
				if err != nil {
					log.Error(err)
				}
				seeds = append(seeds, seed)
			}
			continue
		}
		newCategory := Category{}
		for j, line := range strings.Split(strings.TrimSpace(category), "\n") {
			if j == 0 {
				newCategory.id = strings.Split(line, " ")[0]
				continue
			}
			p := []int{}
			for _, elem := range strings.Split(line, " ") {
				value, err := strconv.Atoi(elem)
				if err != nil {
					log.Error(err)
				}
				p = append(p, value)
			}
			newCategory.dest = append(newCategory.dest, p[0])
			newCategory.src = append(newCategory.src, p[1])
			newCategory.len = append(newCategory.len, p[2])
			newCategory.diff = append(newCategory.diff, p[0]-p[1])
		}
		log.Debug("New category: ", i-1, newCategory)
		almanac = append(almanac, newCategory)
	}
	for i := range almanac {
		if i == 0 {
			continue
		}
		almanac[i].prev = &almanac[i-1]
		almanac[i-1].next = &almanac[i]
	}
	log.Debug(almanac)
	if part2 {
		for location := 0; location < 100000000; location++ {
			locationSeed := almanac[len(almanac)-1].GetSeed(location)
			for i := 0; i < len(seeds)-1; i += 2 {
				seed, len := seeds[i], seeds[i+1]
				if locationSeed >= seed && locationSeed < seed+len {
					return location
				}
			}
		}
	}
	lowest := int(^uint(0) >> 1)
	for _, seed := range seeds {
		if location := almanac[0].GetLocation(seed); location < lowest {
			lowest = location
		}
	}
	return lowest
}
