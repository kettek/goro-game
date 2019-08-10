package entity

import (
	"fmt"
	"math"

	"github.com/kettek/goro-game/interfaces"
	"github.com/kettek/goro/fov"
	"github.com/kettek/goro/pathing"
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
		if a.owner.DistanceTo(entities[0]) >= 3 {
			a.PathfindTo(entities[0], gameMap, entities)
		} else if a.target != nil && a.target.Fighter() != nil && a.target.Fighter().Hp() > 0 {
			fmt.Printf("The %s eyeballs you ferociously.\n", a.owner.Name())
		}
	}
}

func (a *MonsterActor) MoveTo(x, y int, gameMap interfaces.GameMap, entities []interfaces.Entity) {
	destX := float64(x - a.owner.X())
	destY := float64(y - a.owner.Y())

	distance := math.Sqrt(math.Pow(destX, 2) + math.Pow(destY, 2))

	destX = math.Round(destX / distance)
	destY = math.Round(destY / distance)

	if !gameMap.IsBlocked(a.owner.X()+int(destX), a.owner.Y()+int(destY)) && FindEntityAtLocation(entities, a.owner.X()+int(destX), a.owner.Y()+int(destY), BlockMovement, BlockMovement) == nil {
		a.owner.Move(int(destX), int(destY))
	}
}

func (a *MonsterActor) PathfindTo(target interfaces.Entity, gameMap interfaces.GameMap, entities []interfaces.Entity) {
	path := pathing.NewPathFromFunc(gameMap.Width(), gameMap.Height(), func(x, y int) int {
		if gameMap.IsBlocked(x, y) {
			return pathing.MaximumCost
		}
		other := FindEntityAtLocation(entities, x, y, BlockMovement, BlockMovement)
		if other != nil && other != a.owner && other != target {
			return pathing.MaximumCost
		}
		return 0
	}, pathing.AlgorithmAStar)

	steps := path.Compute(a.owner.X(), a.owner.Y(), target.X(), target.Y())

	if len(steps) > 0 {
		x := steps[0].X()
		y := steps[0].Y()
		a.owner.SetX(x)
		a.owner.SetY(y)
	} else {
		a.MoveTo(target.X(), target.Y(), gameMap, entities)
	}
	
}

