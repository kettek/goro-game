package interfaces

type GameMap interface {
	Explored(x, y int) bool
	SetExplored(x, y int, explored bool)
	IsBlocked(x, y int) bool
	InBounds(x, y int) bool
}
