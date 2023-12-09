package day8

import (
	"regexp"
	"strings"

	math "github.com/jrhorner1/AoC/pkg/math"
	log "github.com/sirupsen/logrus"
)

type node struct {
	id          string
	left, right *node
	stop        bool
}
type nodeSlice []*node

func (ns *nodeSlice) New(id string) *node {
	n := &node{id, nil, nil, false}
	*ns = append(*ns, n)
	return n
}

func (ns *nodeSlice) Search(id string) *node {
	for _, n := range *ns {
		if id == n.id {
			return n
		}
	}
	return ns.New(id)
}

func (ns *nodeSlice) At(id string) bool {
	for _, n := range *ns {
		if id == n.id {
			return true
		}
	}
	return false
}

func Puzzle(input *[]byte, part2 bool) int64 {
	in := strings.Split(strings.TrimSpace(string(*input)), "\n\n")
	directions, mapInput, nodes := in[0], in[1], nodeSlice{}
	nodes.New("AAA")
	nodes.New("ZZZ")
	var start, stop nodeSlice
	if !part2 {
		start = append(start, nodes[0])
		stop = append(stop, nodes[1])
	}
	for _, line := range strings.Split(strings.TrimSpace(mapInput), "\n") {
		re := regexp.MustCompile("[0-9A-Z]{3}")
		ids := re.FindAllString(line, -1)
		current := nodes.Search(ids[0])
		current.left = nodes.Search(ids[1])
		current.right = nodes.Search(ids[2])
		if part2 {
			for _, id := range ids {
				re = regexp.MustCompile("[0-9A-Z]{2}[A]")
				if re.MatchString(id) {
					start = append(start, nodes.Search(id))
				}
				re = regexp.MustCompile("[0-9A-Z]{2}[Z]")
				if re.MatchString(id) {
					stop = append(stop, nodes.Search(id))
				}
			}
		}
	}
	log.Debug("Total nodes: ", len(nodes))
	log.Debug("Total starts: ", len(start))
	log.Debug("Total stops: ", len(stop))
	current, steps := start, []int64{}
	allStopped := false
	for !allStopped {
		for _, next := range directions {
			for i := 0; i < len(current); i++ {
				if len(steps) < i+1 {
					steps = append(steps, 0)
				}
				if current[i].stop {
					continue
				}
				switch next {
				case 'L':
					current[i] = current[i].left
				case 'R':
					current[i] = current[i].right
				}
				steps[i]++
				current[i].stop = stop.At(current[i].id)
			}
			allStopped = true
			for _, c := range current {
				if !c.stop {
					allStopped = false
				}
			}
		}
	}
	for _, c := range current {
		log.Debug(*c)
	}
	log.Debug("Total Steps: ", steps)
	if part2 {
		lcm := int64(1)
		for _, s := range steps {
			lcm = math.LCM(lcm, s)
		}
		log.Debug("LCM: ", lcm)
		return lcm
	}
	return steps[0]
}
