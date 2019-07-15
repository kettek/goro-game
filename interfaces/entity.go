package interfaces

import (
	"github.com/kettek/goro"
)

type Entity interface {
	X() int
	SetX(int)
	Y() int
	SetY(int)
	Rune() rune
	SetRune(rune)
	Name() string
	SetName(string)
	Flags() int
	SetFlags(int)
	Style() goro.Style
	SetStyle(goro.Style)
	//
	Fighter() Fighter
	SetFighter(Fighter)
	Actor() Actor
	SetActor(Actor)
	//
	Move(x, y int)
}
