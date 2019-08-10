package entity

import (
	"math"

	"github.com/kettek/goro"
	"github.com/kettek/goro-game/interfaces"
)

type Entity struct {
	x, y    int
	rune    rune
	style   goro.Style
	name    string
	flags   Flags
	fighter interfaces.Fighter
	watcher interfaces.Watcher // FoV?
	actor   interfaces.Actor
}

// Move moves the entity by a given amount.
func (e *Entity) Move(x, y int) {
	e.x += x
	e.y += y
}

func (e *Entity) DistanceTo(other interfaces.Entity) float64 {
	destX := float64(e.x - other.X())
	destY := float64(e.y - other.Y())

	return math.Sqrt(math.Pow(destX, 2) + math.Pow(destY, 2))
}

// X returns the entity's x.
func (e *Entity) X() int {
	return e.x
}

// SetX sets the entity's x.
func (e *Entity) SetX(x int) {
	e.x = x
}

// Y returns the entity's y.
func (e *Entity) Y() int {
	return e.y
}

// SetY sets the entity's y.
func (e *Entity) SetY(y int) {
	e.y = y
}

// Rune returns the entity's rune.
func (e *Entity) Rune() rune {
	return e.rune
}

// SetRune sets the entity's rune.
func (e *Entity) SetRune(r rune) {
	e.rune = r
}

// Style returns the entity's style.
func (e *Entity) Style() goro.Style {
	return e.style
}

// SetStyle sets the entity's style.
func (e *Entity) SetStyle(style goro.Style) {
	e.style = style
}

// Name gets the entity's name.
func (e *Entity) Name() string {
	return e.name
}

// SetName sets the entity's name.
func (e *Entity) SetName(n string) {
	e.name = n
}

// Flags gets the entity's flags.
func (e *Entity) Flags() Flags {
	return e.flags
}

// SetFlags sets the entity's flags.
func (e *Entity) SetFlags(f Flags) {
	e.flags = f
}

func (e *Entity) IsBlocking() bool {
	return e.flags & BlockMovement != 0
}

func (e *Entity) Fighter() interfaces.Fighter {
	return e.fighter
}
func (e *Entity) SetFighter(f interfaces.Fighter) {
	e.fighter = f
}
func (e *Entity) Actor() interfaces.Actor {
	return e.actor
}
func (e *Entity) SetActor(a interfaces.Actor) {
	e.actor = a
	e.actor.SetOwner(e)
}
func (e *Entity) Watcher() interfaces.Watcher {
	return e.watcher
}
func (e *Entity) SetWatcher(w interfaces.Watcher) {
	e.watcher = w
}

// Newinterfaces.Entity returns a pointer to a new populated interfaces.Entity.
func NewEntity(x int, y int, r rune, style goro.Style, name string, flags Flags) interfaces.Entity {
	return &Entity{
		x:     x,
		y:     y,
		rune:  r,
		style: style,
		name:  name,
		flags: flags,
	}
}
