package entity

import (
	"github.com/kettek/goro"
	"github.com/kettek/goro-game/interfaces"
)

func NewPlayerEntity() interfaces.Entity {
	pc := &Entity{
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
	pc.SetWatcher(&BasicWatcher{
		owner: pc,
		radius: 10,
		recompute: true,
	})
	return pc
}
