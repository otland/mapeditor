package ot

type House struct {
	id      uint32
	doorPos Position
	tiles   []Tile
}

type Town struct {
	id        uint32
	name      string
	templePos Position
}

type Map struct {
	width  uint16
	height uint16

	description string
	houseFile   string
	spawnFile   string

	tiles     map[Position]Tile
	houses    []House
	towns     []Town
	waypoints map[Position]string
}
