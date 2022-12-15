package day15

import (
	"fmt"
	"math"
	"strings"

	geom "github.com/jrhorner1/AoC/pkg/math/geometry"
	"github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool, target int) int {
	sensors, beacons := make(Sensors), make(Beacons)
	leftBound, rightBound := math.MaxInt, math.MinInt
	for _, line := range strings.Split(strings.TrimSpace(string(*input)), "\n") {
		sensor, beacon := geom.Point{}, geom.Point{}
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)
		if err != nil {
			logrus.Error(err)
		}
		beacons[beacon] = nil
		sensorRange := sensor.ManhattanDistance(beacon)
		sensors[sensor] = sensorRange
		if sensor.X+sensorRange > rightBound {
			rightBound = sensor.X + sensorRange
		}
		if sensor.X-sensorRange < leftBound {
			leftBound = sensor.X - sensorRange
		}
	}
	if part2 {
		maxXY := target * 2
		for sensor, radius := range sensors {
			for delta := 0; delta <= radius+1; delta++ {
				for i := 0; i < 4; i++ {
					var x, y int
					switch i {
					case 0:
						x = sensor.X - radius - delta + 1
						y = sensor.Y + delta
					case 1:
						x = sensor.X + delta
						y = sensor.Y + radius - delta + 1
					case 2:
						x = sensor.X + radius - delta + 1
						y = sensor.Y - delta
					case 3:
						x = sensor.X - delta
						y = sensor.Y - radius + delta - 1
					}
					current := geom.Point{X: x, Y: y}
					if current.X < 0 || current.Y < 0 || current.X > maxXY || current.Y > maxXY {
						continue
					}
					if sensors.outOfRange(current) {
						return (current.X * 4000000) + current.Y
					}
				}
			}
		}
		return -1
	}
	targetSensors := make(Sensors)
	for sensor, radius := range sensors {
		if (sensor.Y < target && sensor.Y+radius >= target) || (sensor.Y > target && sensor.Y-radius <= target) {
			targetSensors[sensor] = radius
		}
	}
	count := 0
	for x := leftBound; x <= rightBound; x++ {
		current := geom.Point{X: x, Y: target}
		if !sensors.isSensor(current) && !beacons.isBeacon(current) && !targetSensors.outOfRange(current) {
			count++
		}
	}
	return count
}

type Sensors map[geom.Point]int

func (s *Sensors) isSensor(point geom.Point) bool {
	if _, found := (*s)[point]; found {
		return true
	}
	return false
}

func (s *Sensors) outOfRange(point geom.Point) bool {
	for sensor, sensorRange := range *s {
		if dist := sensor.ManhattanDistance(point); dist <= sensorRange {
			return false
		}
	}
	return true
}

type Beacons map[geom.Point]interface{}

func (b *Beacons) isBeacon(point geom.Point) bool {
	if _, found := (*b)[point]; found {
		return true
	}
	return false
}
