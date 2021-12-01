package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    // "time"
)

func GetArgs(instr int, params []int, ip int, memory []int) []int {
    var param_len int
    var args []int
    switch instr {
    case 1,2,7,8: param_len = len(params)
    case 5,6: param_len = len(params) - 1
    }
    for i := 0; i < param_len; i++ {
        temp := memory[ip + i + 1]
        if i == 2 {
            args = append(args, temp)
            continue
        }
        switch params[i] {
        case 1:
            args = append(args, temp)
        case 0:
            args = append(args, memory[temp])
        }
    }
    return args
}

func GetParams(opcode int) (int, []int) {
    opcode_str := strconv.Itoa(opcode)
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

func Run(intcode *[]int, phase int, input <-chan int, output chan<- int, GR int) {
    memory := make([]int, len(*intcode))
    copy(memory, *intcode)
    var bc bool
    var ip_adv, inc int = 0, 0
    // fmt.Println("GR",GR,"Memory:",memory)

    for ip := 0; bc == false; ip += ip_adv {

        instr, params := GetParams(memory[ip])
        args := GetArgs(instr, params, ip, memory)
        if instr == 3 || instr == 4 {
            args = append(args, memory[ip + 1])
        }

        // fmt.Println("GR",GR,"Position:",ip,"Instruction:",instr,"Parameters:",params,"Arguments:",args)
        
        switch instr {
        case 1: // addition
            ip_adv = 4
            memory[args[2]] = args[0] + args[1]
        case 2: // multiplication
            ip_adv = 4
            memory[args[2]] = args[0] * args[1]
        case 3: // input
            ip_adv = 2 
            switch inc {
            case 0:
                memory[args[0]] = phase
                // fmt.Println("GR",GR,"Phase Input:",memory[args[0]])
            default:
                memory[args[0]] = <-input
                // fmt.Println("GR",GR,"Signal Input:",memory[args[0]])
            }
            inc++
        case 4: // output
            ip_adv = 2
            output <- memory[args[0]]
            // fmt.Println("GR",GR,"Signal Output:",memory[args[0]])
        case 5: // jump-if-true
            if args[0] != 0 {
                ip_adv = args[1] - ip
            } else {
                ip_adv = 3
            }
        case 6: // jump-if-false
            if args[0] == 0 {
                ip_adv = args[1] - ip
            } else {
                ip_adv = 3
            }
        case 7: // less than
            ip_adv = 4
            if args[0] < args[1] {
                memory[args[2]] = 1
            } else {
                memory[args[2]] = 0
            }
        case 8: // equals
            ip_adv = 4
            if args[0] == args[1] {
                memory[args[2]] = 1
            } else {
                memory[args[2]] = 0
            }
        case 99: // end program
            bc = true
            if GR == 4 {
                close(output)
                // fmt.Println("GR", GR, "Closing output channel")
            }
            // fmt.Println("GR", GR, "Exiting")
        }
        // fmt.Println("GR",GR,"Memory:",memory)
    }
}

func permutations(arr []int)[][]int{
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
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

    var intcode []int
    for i := range intcode_str {
        temp, _ := strconv.Atoi(intcode_str[i])
        intcode = append(intcode, temp)
    }
    // fmt.Println(intcode)

    // generate list of permutations
    var phasesList [][]int
    // phaseVals := []int{0, 1, 2, 3, 4}
    phaseVals := []int{5, 6, 7, 8, 9}
    for _, perm := range(permutations(phaseVals)){
       phasesList = append(phasesList, perm)
    }
    // fmt.Println(len(phasesList),phasesList)

    // loop through list using each as the values for the intcode
    var ans int = 0
    var output []chan int
    for i := range phasesList {
        phases := phasesList[i]
        // fmt.Println("Beginning phase pattern:", phases)
        output = append(output, make(chan int))
        var input [5]chan int
        for j := range input {
            input[j] = make(chan int, 1)
        }
        input[0] <- 0
        for k := range phases {
            phase := phases[k]
            // fmt.Println("Creating goroutine",k)
            go Run(&intcode, phase, input[k], output[i], k)
        }
        var out int 
        ok := true
        for ok {
            for l := range phases {
                out, ok = <- output[i]
                if ans < out {
                    ans = out
                }
                if ok {
                    if l != len(phases) - 1 {
                        input[l+1] <- out
                    } else {
                        input[0] <- out
                    }
                }
            }
        }        
    }
    fmt.Println("Solution:",ans)
}