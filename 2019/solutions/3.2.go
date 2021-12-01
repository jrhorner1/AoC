package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "math"
)

type carte struct {
    x int
    y int
    s int
}

func get_points(wire string) []carte {
    instr := strings.Split(wire,",")
    var point = carte{ x: 0, y: 0, s: 0 }
    var points []carte
    points = append(points, point)
    for i := 0; i < len(instr); i++ {
        cmd := instr[i]
        dir := cmd[0]
        length, _ := strconv.Atoi(cmd[1:])
        switch dir {
        case 82: // ascii "R"
            for j := 0; j < length; j++ {
                point.x++
                point.s++
                points = append(points, point)
            }
        case 76: // ascii "L"
            for j := 0; j < length; j++ {
                point.x--
                point.s++
                points = append(points, point)
            }
        case 85: // ascii "U"
            for j := 0; j < length; j++ {
                point.y++
                point.s++
                points = append(points, point)
            }
        case 68: // ascii "D"
            for j := 0; j < length; j++ {
                point.y--
                point.s++
                points = append(points, point)
            }
        }

    }
    return points
}

func print_points(points []carte) {
    for i := range points {
        fmt.Println(points[i])
    }
} 

func main() {
    file, _ := os.Open("input")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    points_a := get_points(scanner.Text())
    scanner.Scan()
    points_b := get_points(scanner.Text())

    // print_points(points_a)
    // fmt.Println()
    // print_points(points_b)

    var isecs []carte
    for i := range points_a {
        for j := range points_b {
            if points_a[i].x == points_b[j].x && points_a[i].y == points_b[j].y {
                isecs = append(isecs, points_a[i])
            }
        }
    }

    // print_points(isecs)

    var dist []int
    for i := range isecs {
        raw_dist := int(math.Abs(float64(isecs[i].x)) + math.Abs(float64(isecs[i].y)))
        dist = append(dist, raw_dist)
    }

    // fmt.Println(dist)

    var min_dist int = dist[1]
    for i := 2; i < len(dist); i++ {
        min_dist = int(math.Min(float64(min_dist), float64(dist[i])))
    }

    fmt.Println("Minimum distance:",min_dist)

    var steps []int
    var steps_a, steps_b int 
    for i := 1; i < len(isecs); i++ {
        for j := range points_a {
            if isecs[i].x == points_a[j].x && isecs[i].y == points_a[j].y {
                steps_a = points_a[j].s
            }
        }
        for j := range points_b {
            if isecs[i].x == points_b[j].x && isecs[i].y == points_b[j].y {
                steps_b = points_b[j].s
            }
        }
        steps = append(steps, steps_a + steps_b)
    }
    // fmt.Println(steps)

    min_steps := steps[0]
    for i := range steps {
        min_steps = int(math.Min(float64(min_steps), float64(steps[i])))
    }
    fmt.Println("Minimum steps:",min_steps)
}