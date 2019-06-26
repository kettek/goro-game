package entity

import (
	"github.com/kettek/goro"
)

// Entity is a type for representing an active creature on a Tile.
type Entity struct {
	X, Y  int
	Rune  rune
	Style goro.Style
}

// Move moves the entity by a given amount.
func (b *Entity) Move(x, y int) {
	b.X += x
	b.Y += y
}

// NewEntity returns a pointer to a new populated Entity.
func NewEntity(x int, y int, r rune, s goro.Style) *Entity {
	return &Entity{
		X:     x,
		Y:     y,
		Rune:  r,
		Style: s,
	}
}
