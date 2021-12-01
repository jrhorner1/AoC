package main

import (
	"fmt"
)

type v3d struct {
	x,y,z int
}

type Moon struct {
	pos,vel v3d
}

func ApplyGravity(pos1, pos2 int) int {
	if pos1 > pos2 { 
		return -1
	} else if pos1 < pos2 {
		return 1
	}
	return 0
}

func CalculateVelocity(moons *[]Moon) {
	for i, m1 := range *moons {
		for _, m2 := range *moons {
			if m1 != m2 {
				m1.vel.x += ApplyGravity(m1.pos.x, m2.pos.x)
				m1.vel.y += ApplyGravity(m1.pos.y, m2.pos.y)
				m1.vel.z += ApplyGravity(m1.pos.z, m2.pos.z)
			}
		}
		(*moons)[i].vel = m1.vel
	}
}

func (moon *Moon) Move() {
	moon.pos.x += moon.vel.x
	moon.pos.y += moon.vel.y
	moon.pos.z += moon.vel.z
}

func Abs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}

func (moon *Moon) Energy() int {
	potential := Abs(moon.pos.x) + Abs(moon.pos.y) + Abs(moon.pos.z)
	kinetic := Abs(moon.vel.x) + Abs(moon.vel.y) + Abs(moon.vel.z)
	return potential * kinetic
}

func silver() {
	var moons []Moon
	initialVelocity := v3d{ x: 0, y: 0, z: 0 } 
	positions := []v3d{
		v3d{ x: 14, y: 4,  z: 5   },
		v3d{ x: 12, y: 10, z: 8   },
		v3d{ x: 1,  y: 7,  z: -10 },
		v3d{ x: 16, y: -5, z: 3   }}
	for i := range positions {
		moons = append(moons, Moon{ pos: positions[i], vel: initialVelocity })
	}
	for i := 0; i < 1000; i++ {
		CalculateVelocity(&moons)
		for j := range moons {
			moons[j].Move()
		}
	}
	totalEnergy := 0
	for _, moon := range moons {
		fmt.Println(moon, moon.Energy())
		totalEnergy += moon.Energy()
	}
	fmt.Println("Silver:",totalEnergy)
}

func GCD(a, b int) int { // greatest common denominator
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func LCM(a, b int) int { // lowest common multiple
	return a * b / GCD(a, b)
}

func gold() {
	var moons []Moon
	initialVelocity := v3d{ x: 0, y: 0, z: 0 } 
	positions := []v3d{
		v3d{ x: 14, y: 4,  z: 5   },
		v3d{ x: 12, y: 10, z: 8   },
		v3d{ x: 1,  y: 7,  z: -10 },
		v3d{ x: 16, y: -5, z: 3   }}
	for i := range positions {
		moons = append(moons, Moon{ pos: positions[i], vel: initialVelocity })
	}
	fmt.Println(moons)
	steps, found := v3d{ x: 0, y: 0, z: 0 }, 0
	for i := 1; found < 3; i++ {
		CalculateVelocity(&moons)
		for j := range moons {
			moons[j].Move()
		}
		if steps.x == 0 {
			valid := true
			for k := range moons {
				if moons[k].vel.x != 0 {
					valid = false
				}
			}
			if valid {
				fmt.Println("X:",moons)
				found++
				steps.x = i * 2
			}
		}
		if steps.y == 0 {
			valid := true
			for k := range moons {
				if moons[k].vel.y != 0 {
					valid = false
				}
			}
			if valid {
				fmt.Println("Y:",moons)
				found++
				steps.y = i * 2
			}
		}
		if steps.z == 0 {
			valid := true
			for k := range moons {
				if moons[k].vel.z != 0 {
					valid = false
				}
			}
			if valid {
				fmt.Println("Z:",moons)
				found++
				steps.z = i * 2
			}
		}
	}
	fmt.Println(steps)
	fmt.Println("Gold:",LCM(LCM(steps.x, steps.y), steps.z))
}

func main() {
	silver()
	gold()
}