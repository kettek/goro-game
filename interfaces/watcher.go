package interfaces

import (
	"github.com/kettek/goro/fov"
)

type Watcher interface {
	Owner() Entity
	SetOwner(Entity)
	CreateFromMap(gameMap GameMap)
	Fov() fov.Map
	Recompute()
	ShouldRecompute(bool)
	Radius() int
	SetRadius(int)
}
