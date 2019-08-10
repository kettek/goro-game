package entity

import (
	"github.com/kettek/goro-game/interfaces"
	"github.com/kettek/goro/fov"
)

type BasicWatcher struct {
	owner interfaces.Entity
	fov fov.Map
	radius int
	recompute bool
}

func (w *BasicWatcher) CreateFromMap(gameMap interfaces.GameMap) {
	w.fov = fov.NewMap(gameMap.Width(), gameMap.Height(), fov.AlgorithmBBQ)

	for x := 0; x < gameMap.Width(); x++ {
		for y := 0; y < gameMap.Height(); y++ {
			w.fov.SetBlocksMovement(x, y, gameMap.IsBlocked(x, y))
			w.fov.SetBlocksLight(x, y, gameMap.IsOpaque(x, y))
		}
	}
}

func (w *BasicWatcher) Fov() fov.Map {
	return w.fov
}

func (w *BasicWatcher) Recompute() {
	w.fov.Recompute(w.owner.X(), w.owner.Y(), w.radius, fov.Light{})
	w.recompute = false
}

func (w *BasicWatcher) ShouldRecompute(should bool) {
	w.recompute = should
}

func (w *BasicWatcher) Radius() int {
	return w.radius
}

func (w *BasicWatcher) SetRadius(r int) {
	w.radius = r
}

func (w *BasicWatcher) Owner() interfaces.Entity {
	return w.owner
}

func (w *BasicWatcher) SetOwner(e interfaces.Entity) {
	w.owner = e
}