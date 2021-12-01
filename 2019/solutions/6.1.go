package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func recurseMap(orbitMap *map[string][]string, indirect *int, key string) {
	for newkey, value := range *orbitMap {
		for i := range value {
			if value[i] == key {
				*indirect++
				recurseMap(orbitMap,indirect,newkey)
			}
		}
	}
}

func main() {
	file, _ := os.Open("input")
    defer file.Close()

    var orbitMap = make(map[string][]string)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	orbit := strings.Split(scanner.Text(),")")
		orbitMap[orbit[0]] = append(orbitMap[orbit[0]], orbit[1])
    }

    // print orbit map
    for key, value := range orbitMap {
    	fmt.Print(key," ")
    	for i := range value {
    		fmt.Print(value[i]," ")
    	}
    	fmt.Println()
    }

    var direct,indirect int = 0, 0
    for key, value := range orbitMap {
    	for _ = range value {
    		direct++
			recurseMap(&orbitMap,&indirect,key)
    	}
    }
    fmt.Println(direct + indirect)
}