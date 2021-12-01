package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "../../utils"
)

func OpenFile(filename string) []int {

    file, _ := os.Open(filename)
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
    return intcode
}

func example1() {
    filename := "example1"
    intcode := OpenFile(filename)
    computer := utils.NewComputer(&intcode)
    go computer.Run()

    ok := true
    var out int
    var example1 []int
    for ok {
        out, ok = <-computer.GetOutput()
        example1 = append(example1, out)
    }

    correct := true
    if len(example1) == len(intcode) {
        for i := range example1 {
            if example1[i] != intcode[i] {
                correct = false
            }
        }
    }
    if correct {
        fmt.Println("Example 1: YEET!")
    }
}

func example2() {
    filename := "example2"
    intcode := OpenFile(filename)
    computer := utils.NewComputer(&intcode)
    go computer.Run()
    var example2 int
    example2 = <-computer.GetOutput()
    correct := true
    if len(strconv.Itoa(example2)) != 16 {
        correct = false
    }
    if correct {
        fmt.Println("Example 2: YEET!")
    }
}

func example3() {
    filename := "example3"
    intcode := OpenFile(filename)
    computer := utils.NewComputer(&intcode)
    go computer.Run()
    var example3 int
    example3 = <-computer.GetOutput()
    correct := true
    if example3 != intcode[1] {
        correct = false
    }
    if correct {
        fmt.Println("Example 3: YEET!")
    }
}

func silver() {
    filename := "input"
    intcode := OpenFile(filename)
    computer := utils.NewComputer(&intcode)
    go computer.Run()
    input := computer.GetInput()
    input <- 1
    ok := true
    var out int
    for ok {
        out, ok = <-computer.GetOutput()
        if ok {
            fmt.Println("Silver:",out)
        }
    }
}

func gold() {
    filename := "input"
    intcode := OpenFile(filename)
    computer := utils.NewComputer(&intcode)
    go computer.Run()
    input := computer.GetInput()
    input <- 2
    ok := true
    var out int
    for ok {
        out, ok = <-computer.GetOutput()
        if ok {
            fmt.Println("Gold:",out)
        }
    }
}

func main() {
    example1()
    example2()
    example3()
    silver()
    gold()
}
