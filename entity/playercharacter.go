package entity

import (
	"github.com/kettek/goro"
	"github.com/kettek/goro-game/interfaces"
)

func NewPlayerCharacter() interfaces.Entity {
	pc := &Character{
		rune: '@',
		name: "Player",
		style: goro.Style{
			Foreground: goro.ColorWhite,
		},
		flags: BlockMovement,
		fighter: &BasicFighter{
			hp:      10,
			maxHp:   10,
			defense: 2,
			power:   2,
		},
	}
	return pc
}
