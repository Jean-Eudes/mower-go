package mower

type Area interface {
	IsInPosition(x int, y int) bool
}
