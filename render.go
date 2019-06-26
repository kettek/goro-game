package main

import (
	"github.com/kettek/goro"
	"github.com/kettek/goro-game/entity"
	"github.com/kettek/goro-game/mapping"
)

// RenderAll does something.
func RenderAll(screen *goro.Screen, entities []*entity.Entity, gameMap mapping.GameMap, colors map[string]goro.Color) {
	// Draw all the tiles in the game map.
	for x, column := range gameMap.Tiles {
		for y, tile := range column {
			if tile.BlockSight {
				screen.SetBackground(x, y, colors["darkWall"])
			} else {
				screen.SetBackground(x, y, colors["darkGround"])
			}
		}
	}

	// Draw all the entities in the game map.
	for _, entity := range entities {
		DrawEntity(screen, entity)
	}
	screen.Flush()
}

func ClearAll(screen *goro.Screen, entities []*entity.Entity) {
	for _, entity := range entities {
		ClearEntity(screen, entity)
	}
}

func DrawEntity(screen *goro.Screen, e *entity.Entity) {
	screen.DrawRune(e.X, e.Y, e.Rune, e.Style)
}

func ClearEntity(screen *goro.Screen, e *entity.Entity) {
	screen.DrawRune(e.X, e.Y, ' ', goro.Style{})
}
