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
	return math.IntAbs(q.X-p.X) + math.IntAbs(q.Y-p.Y)
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
		return math.IntAbs(l.A.X - l.B.X)
	} else if l.IsHorizontal() {
		return math.IntAbs(l.A.X - l.B.X)
	} else if l.IsVertical() {
		return math.IntAbs(l.A.Y - l.B.Y)
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
	if math.IntAbs(l.A.X-l.B.X) == math.IntAbs(l.A.Y-l.B.Y) {
		return true
	}
	return false
}
