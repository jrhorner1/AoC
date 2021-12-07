package day6

import (
	"io/ioutil"
	"strings"
)

var filename = "2019/input/6"

func Part1() int {
	return puzzle(input(filename), false)
}

func Part2() int {
	return puzzle(input(filename), true)
}

func input(file string) []string {
	input, _ := ioutil.ReadFile(file)
	output := strings.Split(strings.TrimSpace(string(input)), "\n")
	return output
}

type planet struct {
	name     string
	lChild   string
	rChild   string
	direct   int
	indirect int
}

func puzzle(input []string, part2 bool) int {
	var orbitMap []planet
	for _, in := range input {
		orbit := strings.Split(in, ")")
		inMap := false
		for i := range orbitMap {
			if orbitMap[i].name == orbit[0] {
				inMap = true
				orbitMap[i].direct++
				if orbitMap[i].lChild == "" {
					orbitMap[i].lChild = orbit[1]
				} else {
					orbitMap[i].rChild = orbit[1]
				}
				break
			}
		}
		if !inMap {
			var cPlanet = planet{name: orbit[0], lChild: orbit[1], rChild: "", direct: 1, indirect: 0}
			orbitMap = append(orbitMap, cPlanet)
		}
	}
	var indirect int = 0
	for i := range orbitMap {
		if (orbitMap)[i].name == "COM" {
			recurseMap(&orbitMap, (orbitMap)[i].lChild, &indirect)
			if (orbitMap)[i].rChild != "" {
				recurseMap(&orbitMap, (orbitMap)[i].rChild, &indirect)
			}
		}
	}
	if part2 {
		// map the heirarchy COM > YOU and COM > SAN
		var youMap, sanMap []planet
		recurseHeir(&orbitMap, &youMap, "YOU")
		recurseHeir(&orbitMap, &sanMap, "SAN")

		// find the lowest common ancestor in the heirarchy and count
		// the orbits needed to reach it from both YOU and SAN
		var youCount, sanCount int = 0, 0
	loop:
		for i := range youMap {
			youCount++
			sanCount = 0 // reset the count since we are reiterating over the SAN map
			for j := range sanMap {
				sanCount++
				if youMap[i].name == sanMap[j].name {
					break loop
				}
			}
		}
		// subtract 2 from the answer since it should not include the orbits of YOU and SAN
		return youCount + sanCount - 2
	}
	// count direct and indirect
	orbits := 0
	for i := range orbitMap {
		orbits += orbitMap[i].direct + orbitMap[i].indirect
	}
	return orbits
}

func recurseMap(orbitMap *[]planet, child string, indirect *int) {
	for i := range *orbitMap {
		if (*orbitMap)[i].name == child {
			*indirect++       // moving 1 orbit away from COM
			temp := *indirect // store the current count temporarily
			(*orbitMap)[i].indirect = *indirect
			recurseMap(orbitMap, (*orbitMap)[i].lChild, indirect)
			if (*orbitMap)[i].rChild != "" { // if there is a second orbit,
				*indirect = temp                        // reset the count to the original value before recurse
				(*orbitMap)[i].indirect = *indirect * 2 // double count back to COM since there are 2 orbits
				recurseMap(orbitMap, (*orbitMap)[i].rChild, indirect)
			}
		}
	}
}

func recurseHeir(orbitMap *[]planet, oMap *[]planet, value string) {
	for i := range *orbitMap {
		if (*orbitMap)[i].lChild == value || (*orbitMap)[i].rChild == value {
			*oMap = append(*oMap, (*orbitMap)[i])
			recurseHeir(orbitMap, oMap, (*orbitMap)[i].name)
		}
	}
}
