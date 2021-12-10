package day4

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Puzzle(input *[]byte, index int) int {
	salt := strings.TrimSpace(string(*input))
	pepper := 1
	var zeros string
	for i := 0; i < index; i++ {
		zeros += "0"
	}
	for {
		coin := fmt.Sprintf("%x", md5.Sum([]byte(salt+fmt.Sprint(pepper))))
		if coin[:index] == zeros {
			return pepper
		}
		pepper++
	}
}
