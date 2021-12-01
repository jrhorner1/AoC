package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
    "strings"
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

    var intcode_str []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	intcode_str = strings.Split(scanner.Text(),",")
    }

    // Print the length of the intcode_str array
    // fmt.Println(len(intcode_str))

    // Convert each element of the intcode_str array to an int and assign the value to the appropriate position in the intcode array
    var intcode [145]int
    for i := 0; i < len(intcode_str); i++ {
        var err error
        intcode[i], err = strconv.Atoi(intcode_str[i])
        // fmt.Println(intcode[i])
        check(err)
    }

    // Set position 1 and 2 to the values at the time of the error code
    intcode[1] = 12
    intcode[2] = 2

    // Process each opcode
    var break_code bool = false
    var opcode [4]int
    for intcode_pos := 0; break_code == false ; intcode_pos += 4{
        for i := 0; i < 4; i++ {
            intcode_tmp := intcode_pos + i
            opcode[i] = intcode[intcode_tmp]
        }
        switch opcode[0] {
        case 1: 
            intcode[opcode[3]] = intcode[opcode[1]] + intcode[opcode[2]]
        case 2:
            intcode[opcode[3]] = intcode[opcode[1]] * intcode[opcode[2]]
        case 99:
            break_code = true
        }
    }

    // Print each element of the intcode array after processing
    fmt.Println("Program memory: ")
    for i := 0; i < len(intcode); i++ {
        if i == ( len(intcode) - 1 ) {
            fmt.Printf("%d\n\n", intcode[i])
            break
        }
        fmt.Printf("%d,", intcode[i])
    }

    // print only the final value in position 0
    fmt.Println("1202 program alarm: ",intcode[0])

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}