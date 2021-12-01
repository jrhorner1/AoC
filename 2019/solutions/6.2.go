package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
type planet struct {
    name string
    lChild string
    rChild string
    direct int
    indirect int
}

func recurseMap(orbitMap *[]planet, child string, indirect *int) {
    for i := range *orbitMap {
        if (*orbitMap)[i].name == child {
            *indirect++ // moving 1 orbit away from COM
            temp := *indirect // store the current count temporarily
            (*orbitMap)[i].indirect = *indirect 
            fmt.Println((*orbitMap)[i])
            recurseMap(orbitMap, (*orbitMap)[i].lChild, indirect)
            if (*orbitMap)[i].rChild != "" { // if there is a second orbit, 
                *indirect = temp // reset the count to the original value before recurse
                (*orbitMap)[i].indirect = *indirect * 2 // double count back to COM since there are 2 orbits
                recurseMap(orbitMap, (*orbitMap)[i].rChild, indirect)
            }
        }
    }
}

func printMap(orbitMap *[]planet) {
    var indirect int = 0
    for i := range *orbitMap {
        if (*orbitMap)[i].name == "COM" {
            fmt.Println((*orbitMap)[i])
            recurseMap(orbitMap, (*orbitMap)[i].lChild, &indirect)
            if (*orbitMap)[i].rChild != "" {
                recurseMap(orbitMap, (*orbitMap)[i].rChild, &indirect)
            }
        }
    }
}

func recurseHeir( orbitMap *[]planet, oMap *[]planet, value string) {
    for i := range *orbitMap {
        if (*orbitMap)[i].lChild == value || (*orbitMap)[i].rChild == value {
            *oMap = append(*oMap, (*orbitMap)[i])
            recurseHeir(orbitMap,oMap,(*orbitMap)[i].name)
        }
    }
}

func main() {
	file, _ := os.Open("input")
    defer file.Close()

    var orbitMap []planet
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	orbit := strings.Split(scanner.Text(),")")
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


    // print orbit map in order starting at COM
    printMap(&orbitMap)

    // count direct and indirect 
    var direct,indirect int = 0,0
    for i := range orbitMap {
        direct += orbitMap[i].direct
        indirect += orbitMap[i].indirect
    }
    fmt.Println("Part 1:",direct + indirect)

    // map the heirarchy COM > YOU and COM > SAN
    var youMap,sanMap []planet
    recurseHeir(&orbitMap,&youMap,"YOU")
    recurseHeir(&orbitMap,&sanMap,"SAN")

    // find the lowest common ancestor in the heirarchy and count
    // the orbits needed to reach it from both YOU and SAN
    var lca planet
    var youCount,sanCount int = 0,0
    var breakCode bool
    for i := range youMap {
        youCount++
        sanCount = 0 // reset the count since we are reiterating over the SAN map
        for j := range sanMap {
            sanCount++
            if youMap[i].name == sanMap[j].name {
                lca = youMap[i]
                fmt.Println("Lowest Common Ancestor:",lca)
                breakCode = true
                break
            }
        }
        if breakCode {
            break
        }
    }
    // subtract 2 from the answer since it should not include the orbits of YOU and SAN
    fmt.Println("Part 2:",youCount + sanCount - 2)
}