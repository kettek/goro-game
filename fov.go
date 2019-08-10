package main

import (
	"github.com/kettek/goro-game/interfaces"
	"github.com/kettek/goro/fov"
)

func InitializeFoV(g interfaces.GameMap) fov.Map {
	fovMap := fov.NewMap(g.Width(), g.Height(), fov.AlgorithmBBQ)

	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			fovMap.SetBlocksMovement(x, y, g.IsBlocked(x, y))
			fovMap.SetBlocksLight(x, y, g.IsOpaque(x, y))
		}
	}

	return fovMap
}

func RecomputeFoV(fovMap fov.Map, centerX, centerY int, radius int, light fov.Light) {
	fovMap.Recompute(centerX, centerY, radius, light)
}
