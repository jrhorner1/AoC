package utils

import (
	"os"
	"bufio"
	"strings"
	"strconv"
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