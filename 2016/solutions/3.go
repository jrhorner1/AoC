package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type triangle struct {
	a int
	b int
	c int
}

func main() {
	input, _ := ioutil.ReadFile("2016/input/3")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	triangles := []triangle{}
	for _, s := range in {
		raw := strings.Fields(s)
		a, _ := strconv.Atoi(raw[0])
		b, _ := strconv.Atoi(raw[1])
		c, _ := strconv.Atoi(raw[2])
		t := triangle{a, b, c}
		triangles = append(triangles, t)
	}
	validTriangles := testTrianges(&triangles)
	fmt.Println("Part 1:", validTriangles)

	ts := &triangles
	validTriangles = 0
	for i := 0; i <= len(*ts)-3; i += 3 {
		t := []triangle{}
		t = append(t, triangle{(*ts)[i].a, (*ts)[i+1].a, (*ts)[i+2].a})
		t = append(t, triangle{(*ts)[i].b, (*ts)[i+1].b, (*ts)[i+2].b})
		t = append(t, triangle{(*ts)[i].c, (*ts)[i+1].c, (*ts)[i+2].c})
		validTriangles += testTrianges(&t)
	}
	fmt.Println("Part 2:", validTriangles)
	fmt.Println("Happy Holidays 2021!")
}

func testTrianges(triangles *[]triangle) int {
	valid := 0
	for _, t := range *triangles {
		if (t.a+t.b) > t.c && (t.a+t.c) > t.b && (t.b+t.c) > t.a {
			valid++
		}
	}
	return valid
}
