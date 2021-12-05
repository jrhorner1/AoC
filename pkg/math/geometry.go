package math/geometry

type Point struct {
	X int
	Y int
}

func (p *Point) ManhattanDistance(q Point) int {
	return IntAbs(q.X-p.X) + IntAbs(q.Y-p.Y)
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
	} else {
		return Point{}
	}
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
	} else {
		return Point{}
	}
}

func (l *Line) Length() int {
	if l.IsDiagonal() {
		return l.End().X - l.Start().X
	} else if l.IsHorizontal() {
		return l.End().X - l.Start().X
	} else if l.IsVertical() {
		return l.End().Y - l.Start().Y
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
	if l.Start().X-l.End().X == l.Start().Y-l.End().Y {
		return true
	}
	return false
}
