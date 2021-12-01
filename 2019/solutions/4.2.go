package main

import (
	"fmt"
	"strconv"
)

func main() {
	pwd_l := 240298
	pwd_h := 784956
	var ans []int
	for pwd := pwd_l; pwd <= pwd_h; pwd++ {
		var doub, decl bool = false, false
		pwd_s := strconv.Itoa(pwd)
		for i := 0; i < len(pwd_s)-1; i++ {
			// check for non-repeating doubled ints
			switch i {
			case 0: 
				if pwd_s[i] == pwd_s[i+1] && pwd_s[i+1] != pwd_s[i+2] {
					doub = true
				} 
			case len(pwd_s)-2: 
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
		}
		if doub == true && decl == false {
			ans = append(ans, pwd)
		}
	}
	// fmt.Println(ans)
	fmt.Println("Possible passwords:",len(ans))
}