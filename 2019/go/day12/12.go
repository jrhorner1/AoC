package day12

import (
	"strconv"
	"strings"
)

type v3d struct {
	x, y, z int
}

type Moon struct {
	pos, vel v3d
}

func Puzzle(input *[]byte, part2 bool) int {
	positions := []v3d{}
	for _, i := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		positionString := strings.Split(i[1:len(i)-1], ", ")
		position := v3d{}
		position.x, _ = strconv.Atoi(positionString[0][2:])
		position.y, _ = strconv.Atoi(positionString[1][2:])
		position.z, _ = strconv.Atoi(positionString[2][2:])
		positions = append(positions, position)
	}
	var moons []Moon
	initialVelocity := v3d{x: 0, y: 0, z: 0}
	for i := range positions {
		moons = append(moons, Moon{pos: positions[i], vel: initialVelocity})
	}
	if part2 {
		steps, found := v3d{x: 0, y: 0, z: 0}, 0
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
					found++
					steps.z = i * 2
				}
			}
		}
		return LCM(LCM(steps.x, steps.y), steps.z)
	}
	for i := 0; i < 1000; i++ {
		CalculateVelocity(&moons)
		for j := range moons {
			moons[j].Move()
		}
	}
	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += moon.Energy()
	}
	return totalEnergy
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

func GCD(a, b int) int { // greatest common denominator
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int { // lowest common multiple
	return a * b / GCD(a, b)
}
