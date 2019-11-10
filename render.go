package main

import (
	"strings"

	"github.com/kettek/goro"
	"github.com/kettek/goro/fov"

	"github.com/kettek/goro-game/interfaces"
	"github.com/kettek/goro-game/results"
)

// DrawAll draws all entities and the gameMap to the screen and flushes it.
func DrawAll(screen *goro.Screen, entities []interfaces.Entity, gameMap interfaces.GameMap, fovMap fov.Map, fovRecompute bool, colors map[string]goro.Color) {
	if fovRecompute {
		// Draw all the tiles in the game map.
		for x := 0; x < gameMap.Width(); x++ {
			for y := 0; y < gameMap.Height(); y++ {
				visible := fovMap.Visible(x, y)
				if visible {
					if gameMap.IsOpaque(x, y) {
						screen.SetRune(x, y+1, '#')
						screen.SetBackground(x, y+1, colors["lightWall"])
					} else {
						screen.SetRune(x, y+1, '.')
						screen.SetBackground(x, y+1, colors["lightGround"])
					}
					gameMap.SetExplored(x, y, true)
				} else if gameMap.Explored(x, y) {
					if gameMap.IsOpaque(x, y) {
						screen.SetRune(x, y+1, '#')
						screen.SetBackground(x, y+1, colors["darkWall"])
					} else {
						screen.SetRune(x, y+1, '.')
						screen.SetBackground(x, y+1, colors["darkGround"])
					}
				}
			}
		}
	}

	// Draw all the entities in the game map.
	for _, entity := range entities {
		DrawEntity(screen, entity, fovMap)
	}

}

// ClearAll clears all entities from the screen.
func ClearAll(screen *goro.Screen, entities []interfaces.Entity, fovMap fov.Map) {
	for _, entity := range entities {
		ClearEntity(screen, entity, fovMap)
	}
}

func DrawResults(screen *goro.Screen, turnResults []interfaces.Result, maxResults int) {
	i := 0
	for _, r := range turnResults {
		switch result := r.(type) {
		case *results.Message:
			if i >= maxResults {
				screen.DrawString(0, i, "--MORE--", goro.Style{})
				continue
			}
			/*if len(result.Text) > screen.Colums {
				// Get remainder of text and set it to result.Text?
			}*/
			screen.DrawString(0, i, result.Text, goro.Style{})
			result.Seen = true
			i = i + 1
		}
	}
}

func ClearResults(screen *goro.Screen, turnResults []interfaces.Result, maxResults int) {
	i := 0
	for _, r := range turnResults {
		switch result := r.(type) {
		case *results.Message:
			if i >= maxResults {
				screen.DrawString(0, i, "        ", goro.Style{})
				continue
			}
			if result.Seen {
				screen.DrawString(0, i, strings.Repeat(" ", len(result.Text)), goro.Style{})
				i = i + 1
			}
		}
	}
}

// DrawEntity draws a given entity to the screen.
func DrawEntity(screen *goro.Screen, e interfaces.Entity, fovMap fov.Map) {
	x := e.X()
	y := e.Y()
	if fovMap.Visible(x, y) {
		screen.SetRune(x, y+1, e.Rune())
		screen.SetForeground(x, y+1, e.Style().Foreground)
	}
}

// ClearEntity clears a given entity from the screen.
func ClearEntity(screen *goro.Screen, e interfaces.Entity, fovMap fov.Map) {
	x := e.X()
	y := e.Y()
	if fovMap.Visible(x, y) {
		screen.SetRune(x, y+1, ' ')
	}
}
