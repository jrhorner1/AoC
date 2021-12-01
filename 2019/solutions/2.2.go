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

    var intcode [145]int
    var noun, verb int
    for noun = 0; noun < 100; noun++ {
        for verb = 0; verb < 100; verb++ {
            // Convert each element of the intcode_str array to an int and assign the value to the appropriate position in the intcode array
            for i := 0; i < len(intcode_str); i++ {
                var err error
                intcode[i], err = strconv.Atoi(intcode_str[i])
                // fmt.Println(intcode[i])
                check(err)
            }

            // set the noun and verb
            intcode[1] = noun
            intcode[2] = verb

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
            // check if intcode[0] == 19690720
            if intcode[0] == 19690720 {
                break
            }
        }
        // check if intcode[0] == 19690720
        if intcode[0] == 19690720 {
            break
        }
    }

    // Print each element of the intcode array
    // for i := 0; i < len(intcode); i++ {
    //     fmt.Printf("%d,", intcode[i])
    // }
    // fmt.Println()
    fmt.Println("Noun/Verb pair for 19690720 output: ",100 * noun + verb)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}