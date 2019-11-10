package interfaces

type GameMap interface {
	Width() int
	Height() int
	Explored(x, y int) bool
	MakeMap(maxRooms, roomMinSize, roomMaxSize int, entities *[]Entity, maxMonsters int)
	SetExplored(x, y int, explored bool)
	IsBlocked(x, y int) bool
	IsOpaque(x, y int) bool
	InBounds(x, y int) bool
	CostAt(x, y int) uint32
}
