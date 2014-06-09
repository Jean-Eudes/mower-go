package mower

import "fmt"

type Orientation int

const (
	N Orientation = iota
	E
	S
	W
)

type Mower struct {
	Area        Area
	Orientation Orientation
	x           int
	y           int
}

func (m *Mower) Advance() {
	switch m.Orientation {
	case N:
		if m.Area.IsInPosition(m.x, m.y+1) {
			m.y++
		}
	case S:
		if m.Area.IsInPosition(m.x, m.y-1) {
			m.y--
		}
	case W:
		if m.Area.IsInPosition(m.x-1, m.y) {
			m.x--
		}
	case E:
		if m.Area.IsInPosition(m.x+1, m.y) {
			m.x++
		}
	default:
	}
}

func (m *Mower) TurnRight() {
	m.Orientation = (m.Orientation + 1) % 4
}

func (m *Mower) TurnLeft() {
	m.Orientation = (m.Orientation + 3) % 4
}

func (m *Mower) Print() {
	fmt.Println("tondeuse : (", m.x, ",", m.y, "), orientation :", m.Orientation)
}

func NewMower(x int, y int, area Area, orientation Orientation) *Mower {
	var m = Mower{area, orientation, x, y}
	return &m
}
