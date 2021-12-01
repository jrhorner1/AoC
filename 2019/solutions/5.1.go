package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func get_opcode(opcode_str string) (int, []int) {
    var instr int
        var params []int
        switch len(opcode_str) {
        case 1,2:
            instr, _ = strconv.Atoi(opcode_str)
            for i := 2; i >= 0; i-- {
                param := 0
                params = append(params, param)
            }
        case 3:
            instr, _ = strconv.Atoi(opcode_str[1:])
            for i := 0; i >= -2; i-- {
                if i == 0 {
                    param := int(opcode_str[i] - '0')
                    params = append(params, param)
                } else {
                    param := 0
                    params = append(params, param)
                }
            }
        case 4:
            instr, _ = strconv.Atoi(opcode_str[2:])
            for i := 1; i >= -1; i-- {
                if i == 0 || i == 1 {
                    param := int(opcode_str[i] - '0')
                    params = append(params, param)
                } else {
                    param := 0
                    params = append(params, param)
                }
            }
        case 5:
            instr, _ = strconv.Atoi(opcode_str[3:])
            for i := 2; i >= 0; i-- {
                param := int(opcode_str[i] - '0')
                params = append(params, param)
            }
        }
    return instr, params
}

func main() {
	// change to request the input file from user
    file, _ := os.Open("input")
    defer file.Close()

    var intcode_str []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	intcode_str = strings.Split(scanner.Text(),",")
    }

    // var intcode []int
    // // Convert each element of the intcode_str array to an int and append the value to the intcode array
    // for i := 0; i < len(intcode_str); i++ {
    //     temp, _ := strconv.Atoi(intcode_str[i])
    //     intcode = append(intcode, temp)
    // }
    // fmt.Println(intcode_str)

    // Process each opcode
    var break_code bool
    var intcode_adv int = 0
    for intcode_pos := 0; break_code == false; intcode_pos += intcode_adv {
        //opcode_str := intcode_str[intcode_pos]
        instr, params := get_opcode(intcode_str[intcode_pos])
        // process opcode
        switch instr {
        case 1: // addition
            intcode_adv = 4
            //intcode[opcode[3]] = intcode[opcode[1]] + intcode[opcode[2]]
            for i := 0; i < len(params); i++ {
                switch params[i] {
                case 1:
                    params[i], _ = strconv.Atoi(intcode_str[intcode_pos + i + 1])
                case 0:
                    temp, _ := strconv.Atoi(intcode_str[intcode_pos + i + 1])
                    params[i], _ = strconv.Atoi(intcode_str[temp])
                }
            }
            params[2] = params[0] + params[1]
            temp, _ := strconv.Atoi(intcode_str[intcode_pos + 3])
            intcode_str[temp] = strconv.Itoa(params[2])
        case 2: // multiplication
            intcode_adv = 4
            //intcode[opcode[3]] = intcode[opcode[1]] * intcode[opcode[2]]
            for i := 0; i < len(params); i++ {
                switch params[i] {
                case 1:
                    params[i], _ = strconv.Atoi(intcode_str[intcode_pos + i + 1])
                case 0:
                    temp, _ := strconv.Atoi(intcode_str[intcode_pos + i + 1])
                    params[i], _ = strconv.Atoi(intcode_str[temp])
                }
            }
            params[2] = params[0] * params[1]
            temp, _ := strconv.Atoi(intcode_str[intcode_pos + 3])
            intcode_str[temp] = strconv.Itoa(params[2])
        case 3: // input
            intcode_adv = 2 
            input := bufio.NewScanner(os.Stdin)
            fmt.Print("ID: ")
            input.Scan()
            intcode_str[intcode_pos+1] = scanner.Text()
        case 4: // output
            intcode_adv = 2
            if params[0] == 0 {
                temp, _ := strconv.Atoi(intcode_str[intcode_pos+1])
                fmt.Println(intcode_str[temp])
                continue
            }
            fmt.Println(intcode_str[intcode_pos+1])
        case 99: // end program
            break_code = true
        }
        // fmt.Println(instr,params)
    }

    // Print each element of the intcode array
    // for i := 0; i < len(intcode); i++ {
    //     fmt.Printf("%d,", intcode[i])
    // }
    // fmt.Println(intcode_str)
}