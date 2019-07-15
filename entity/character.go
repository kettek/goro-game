package entity

import (
	"github.com/kettek/goro"
	"github.com/kettek/goro-game/interfaces"
)

type Character struct {
	x, y    int
	rune    rune
	style   goro.Style
	name    string
	flags   Flags
	fighter interfaces.Fighter
	actor   interfaces.Actor
}

// Move moves the entity by a given amount.
func (e *Character) Move(x, y int) {
	e.x += x
	e.y += y
}

// X returns the entity's x.
func (e *Character) X() int {
	return e.x
}

// SetX sets the entity's x.
func (e *Character) SetX(x int) {
	e.x = x
}

// Y returns the entity's y.
func (e *Character) Y() int {
	return e.y
}

// SetY sets the entity's y.
func (e *Character) SetY(y int) {
	e.y = y
}

// Rune returns the entity's rune.
func (e *Character) Rune() rune {
	return e.rune
}

// SetRune sets the entity's rune.
func (e *Character) SetRune(r rune) {
	e.rune = r
}

// Style returns the entity's style.
func (e *Character) Style() goro.Style {
	return e.style
}

// SetStyle sets the entity's style.
func (e *Character) SetStyle(style goro.Style) {
	e.style = style
}

// Name gets the entity's name.
func (e *Character) Name() string {
	return e.name
}

// SetName sets the entity's name.
func (e *Character) SetName(n string) {
	e.name = n
}

// Flags gets the entity's flags.
func (e *Character) Flags() Flags {
	return e.flags
}

// SetFlags sets the entity's flags.
func (e *Character) SetFlags(f Flags) {
	e.flags = f
}

func (e *Character) Fighter() interfaces.Fighter {
	return e.fighter
}
func (e *Character) SetFighter(f interfaces.Fighter) {
	e.fighter = f
}
func (e *Character) Actor() interfaces.Actor {
	return e.actor
}
func (e *Character) SetActor(a interfaces.Actor) {
	e.actor = a
	e.actor.SetOwner(e)
}

// Newinterfaces.Entity returns a pointer to a new populated interfaces.Entity.
func NewCharacter(x int, y int, r rune, style goro.Style, name string, flags Flags) interfaces.Entity {
	return &Character{
		x:     x,
		y:     y,
		rune:  r,
		style: style,
		name:  name,
		flags: flags,
	}
}
