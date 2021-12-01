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
			// check for doubled ints
			if pwd_s[i] == pwd_s[i+1] {
				doub = true
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
	//fmt.Println(ans)
	fmt.Println(len(ans))
}