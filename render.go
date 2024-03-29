package main

import (
	"github.com/kettek/goro"
	"github.com/kettek/goro/fov"

	"github.com/kettek/goro-game/entity"
	"github.com/kettek/goro-game/mapping"
)

// DrawAll draws all entities and the gameMap to the screen and flushes it.
func DrawAll(screen *goro.Screen, entities []*entity.Entity, gameMap mapping.GameMap, fovMap fov.Map, fovRecompute bool, colors map[string]goro.Color) {
	if fovRecompute {
		// Draw all the tiles in the game map.
		for x, column := range gameMap.Tiles {
			for y, tile := range column {
				visible := fovMap.Visible(x, y)

				if visible {
					if tile.BlockSight {
            screen.SetRune(x, y, '#')
						screen.SetBackground(x, y, colors["lightWall"])
					} else {
            screen.SetRune(x, y, '.')
						screen.SetBackground(x, y, colors["lightGround"])
					}
					gameMap.SetExplored(x, y, true)
				} else if gameMap.Explored(x, y) {
					if tile.BlockSight {
            screen.SetRune(x, y, '#')
						screen.SetBackground(x, y, colors["darkWall"])
					} else {
            screen.SetRune(x, y, '.')
            screen.SetStyle(x, y, goro.Style{
              Background: colors["darkGround"],
              Reverse: true,
              Blink: true,
            })
					}
				}
			}
		}
	}

	// Draw all the entities in the game map.
	for _, entity := range entities {
		DrawEntity(screen, entity, fovMap)
	}
	screen.Flush()
}

// ClearAll clears all entities from the screen.
func ClearAll(screen *goro.Screen, entities []*entity.Entity) {
	for _, entity := range entities {
		ClearEntity(screen, entity)
	}
}

// DrawEntity draws a given entity to the screen.
func DrawEntity(screen *goro.Screen, e *entity.Entity, fovMap fov.Map) {
	if fovMap.Visible(e.X, e.Y) {
		screen.SetRune(e.X, e.Y, e.Rune)
		screen.SetForeground(e.X, e.Y, e.Style.Foreground)
	}
}

// ClearEntity clears a given entity from the screen.
func ClearEntity(screen *goro.Screen, e *entity.Entity) {
	screen.SetRune(e.X, e.Y, ' ')
}
