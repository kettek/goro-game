package entity

import (
	"github.com/kettek/goro-game/interfaces"
	"github.com/kettek/goro/fov"
)

type MonsterActor struct {
	owner  interfaces.Entity
	target interfaces.Entity
}

func (a *MonsterActor) Owner() interfaces.Entity {
	return a.owner
}

func (a *MonsterActor) SetOwner(e interfaces.Entity) {
	a.owner = e
}

func (a *MonsterActor) Target() interfaces.Entity {
	return a.target
}

func (a *MonsterActor) SetTarget(e interfaces.Entity) {
	a.target = e
}

func (a *MonsterActor) TakeTurn(fovMap fov.Map, gameMap interfaces.GameMap, entities []interfaces.Entity) {
	if fovMap.Visible(a.owner.X(), a.owner.Y()) {
		if false { // move within 2 squares of player
		} else if a.target.Fighter() != nil && a.target.Fighter().Hp() > 0 {
			//
		}
		// TODO: A* move logic
	}
}
