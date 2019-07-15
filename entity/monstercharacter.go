package entity

import (
	"github.com/kettek/goro"
	"github.com/kettek/goro-game/interfaces"
)

func NewMonsterCharacter(x, y int, r rune, style goro.Style, name string) interfaces.Entity {
	c := &Character{
		x:     x,
		y:     y,
		rune:  r,
		name:  name,
		style: style,
		flags: BlockMovement,
		fighter: &BasicFighter{
			hp:      10,
			maxHp:   10,
			defense: 2,
			power:   2,
		},
	}
	c.SetActor(&MonsterActor{
		owner: c,
	})
	return c
}
