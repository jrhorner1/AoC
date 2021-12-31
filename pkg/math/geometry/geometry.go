package geometry

import (
	"github.com/jrhorner1/AoC/pkg/math"
)

type Point Point2D

type Point2D struct {
	X, Y int
}

type Point3D struct {
	X, Y, Z int
}

func (p *Point) ManhattanDistance(q Point) int {
	return math.Abs(q.X-p.X).(int) + math.Abs(q.Y-p.Y).(int)
}

type Line struct {
	A Point
	B Point
}

func (l *Line) Start() Point {
	if l.IsHorizontal() {
		if l.A.X < l.B.X {
			return l.A
		} else {
			return l.B
		}
	} else if l.IsVertical() {
		if l.A.Y < l.B.Y {
			return l.A
		} else {
			return l.B
		}
	} else if l.IsDiagonal() {
		if l.A.X < l.B.X {
			return l.A
		} else {
			return l.B
		}
	}
	return Point{}
}

func (l *Line) End() Point {
	if l.IsHorizontal() {
		if l.A.X < l.B.X {
			return l.B
		} else {
			return l.A
		}
	} else if l.IsVertical() {
		if l.A.Y < l.B.Y {
			return l.B
		} else {
			return l.A
		}
	} else if l.IsDiagonal() {
		if l.A.X < l.B.X {
			return l.B
		} else {
			return l.A
		}
	}
	return Point{}
}

func (l *Line) Length() int {
	if l.IsDiagonal() {
		return math.Abs(l.A.X - l.B.X).(int)
	} else if l.IsHorizontal() {
		return math.Abs(l.A.X - l.B.X).(int)
	} else if l.IsVertical() {
		return math.Abs(l.A.Y - l.B.Y).(int)
	}
	return 0
}

func (l *Line) IsHorizontal() bool {
	if l.A.Y == l.B.Y {
		return true
	}
	return false
}

func (l *Line) IsVertical() bool {
	if l.A.X == l.B.X {
		return true
	}
	return false
}

func (l *Line) IsDiagonal() bool {
	if math.Abs(l.A.X-l.B.X).(int) == math.Abs(l.A.Y-l.B.Y).(int) {
		return true
	}
	return false
}
