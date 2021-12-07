package day4

import (
	"strconv"
	"strings"
)

func Puzzle(input *[]byte, part2 bool) int {
	in := strings.Split(strings.TrimSpace(string(*input)), "-")
	pwd_l, _ := strconv.Atoi(in[0])
	pwd_h, _ := strconv.Atoi(in[1])
	var ans []int
	for pwd := pwd_l; pwd <= pwd_h; pwd++ {
		var doub, decl bool = false, false
		pwd_s := strconv.Itoa(pwd)
		for i := 0; i < len(pwd_s)-1; i++ {
			if part2 {
				// check for non-repeating doubled ints
				switch i {
				case 0:
					if pwd_s[i] == pwd_s[i+1] && pwd_s[i+1] != pwd_s[i+2] {
						doub = true
					}
				case len(pwd_s) - 2:
					if pwd_s[i-1] != pwd_s[i] && pwd_s[i] == pwd_s[i+1] {
						doub = true
					}
				default:
					if pwd_s[i-1] != pwd_s[i] && pwd_s[i] == pwd_s[i+1] && pwd_s[i+1] != pwd_s[i+2] {
						doub = true
					}
				}
				// check for declining ints
				if pwd_s[i] > pwd_s[i+1] {
					decl = true
				}
			} else {
				// check for doubled ints
				if pwd_s[i] == pwd_s[i+1] {
					doub = true
				}
				// check for declining ints
				if pwd_s[i] > pwd_s[i+1] {
					decl = true
				}
			}
		}
		if doub == true && decl == false {
			ans = append(ans, pwd)
		}
	}
	return len(ans)
}
