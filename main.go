package main

import (
	"log"

	"github.com/kettek/goro"
	"github.com/kettek/goro-game/entity"
	"github.com/kettek/goro-game/mapping"
)

func main() {
	// Initialize goro!
	if err := goro.InitTCell(); err != nil {
		log.Fatal(err)
	}

	goro.Run(func(screen *goro.Screen) {
		// Screen configuration.
		screen.SetTitle("goRo-game")

		// Randomize our seed so the map is randomized per run.
		goro.SetSeed(goro.RandomSeed())

		// Our initial variables.
		var entities []*entity.Entity
		mapWidth, mapHeight := 80, 45
		maxRooms, roomMinSize, roomMaxSize := 30, 10, 6

		colors := map[string]goro.Color{
			"darkWall":   goro.Color{R: 0, G: 0, B: 100, A: 255},
			"darkGround": goro.Color{R: 100, G: 100, B: 100, A: 255},
		}

		player := entity.NewEntity(screen.Columns/2, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorWhite})
		npc := entity.NewEntity(screen.Columns/2-5, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorYellow})

		entities = append(entities, player, npc)

		gameMap := mapping.GameMap{
			Width:  mapWidth,
			Height: mapHeight,
		}
		gameMap.Initialize()
		gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, mapWidth, mapHeight, player)

		for {
			// Draw screen.
			DrawAll(screen, entities, gameMap, colors)

			ClearAll(screen, entities)

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					if !gameMap.IsBlocked(player.X+action.X, player.Y+action.Y) {
						player.Move(action.X, action.Y)
					}
				case ActionQuit:
					goro.Quit()
				}
			case goro.EventQuit:
				return
			}
		}
	})
}
