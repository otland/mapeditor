package ot

type Position struct {
	x uint16
	y uint16
	z uint8
}

func (pos Position) cmp(target Position) bool {
	return pos.x == target.x && pos.y == target.y && pos.z == target.z
}
