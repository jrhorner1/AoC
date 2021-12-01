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

func main() {
	// change to request the input file from user
    file, err := os.Open("input")
    check(err)
    defer file.Close()

    var fuel int64

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
    	check(err)
    	fuel_req := ( mass / 3 ) - 2
    	fuel += fuel_req
    }

    fmt.Println(fuel)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}