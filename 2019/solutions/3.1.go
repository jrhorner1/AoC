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
}

func get_points(wire string) []carte {
    instr := strings.Split(wire,",")
    var point = carte{ x: 0, y: 0 }
    var points []carte
    points = append(points, point)
    for i := 0; i < len(instr); i++ {
        cmd := instr[i]
        dir := cmd[0]
        len, _ := strconv.Atoi(cmd[1:])
        switch dir {
        case 82: // ascii "R"
            for j := 0; j < len; j++ {
                point.x++
                points = append(points, point)
            }
        case 76: // ascii "L"
            for j := 0; j < len; j++ {
                point.x--
                points = append(points, point)
            }
        case 85: // ascii "U"
            for j := 0; j < len; j++ {
                point.y++
                points = append(points, point)
            }
        case 68: // ascii "D"
            for j := 0; j < len; j++ {
                point.y--
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
            if points_a[i] == points_b[j] {
                isecs = append(isecs, points_a[i])
            }
        }
    }

    // print_points(isecs)

    var dist []int
    for i := range isecs {
        if isecs[i].x < 0 {
            isecs[i].x = isecs[i].x * -1
        }
        if isecs[i].y < 0 {
            isecs[i].y = isecs[i].y * -1
        }
        raw_dist := isecs[i].x + isecs[i].y
        dist = append(dist, raw_dist)
    }

    // fmt.Println(dist)

    var min_dist int = dist[1]
    for i := 2; i < len(dist); i++ {
        min_dist = int(math.Min(float64(min_dist), float64(dist[i])))
    }

    fmt.Println("Minimum distance:",min_dist)
}