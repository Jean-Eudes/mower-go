package mower

type Rectangle struct {
	Height int
	Weight int
}

func (r Rectangle) IsInPosition(x int, y int) bool {
	return (x >= 0 && y >= 0) && (x < r.Height && y < r.Weight)
}
