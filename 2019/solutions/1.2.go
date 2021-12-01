package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func fuelRequired(mass int) int {
    fuel_req := ( mass / 3 ) - 2
    if fuel_req < 0 {
        return 0
    }
    return fuel_req
}

func main() {
	// change to request the input file from user
    file, err := os.Open("input")
    check(err)
    defer file.Close()

    var fuel int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	module_mass, err := strconv.Atoi(scanner.Text())
    	check(err)

        fuel_req := fuelRequired(module_mass)
        for fuel_req > 0 {
            fuel += fuel_req
            fuel_req = fuelRequired(fuel_req)
        }
    }

    fmt.Println(fuel)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}