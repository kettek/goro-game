package main

import (
	"fmt"
	"log"

	"github.com/kettek/goro"
	"github.com/kettek/goro-game/entity"
	"github.com/kettek/goro-game/interfaces"
	"github.com/kettek/goro-game/mapping"
	"github.com/kettek/goro-game/results"
	"github.com/kettek/goro/fov"
)

func main() {
	// Initialize goro!
	if err := goro.InitEbiten(); err != nil {
		//if err := goro.InitTCell(); err != nil {
		log.Fatal(err)
	}

	goro.Setup(func(screen *goro.Screen) {
		// Screen configuration.
		screen.SetTitle("goRo-game")
		screen.SetSize(80, 24)

		screen.SetDefaultForeground(goro.ColorWhite)
		screen.SetDefaultBackground(goro.ColorBlack)
		// Randomize our seed so the map is randomized per run.
		goro.SetSeed(goro.RandomSeed())
	})

	goro.Run(func(screen *goro.Screen) {
		// Our initial variables.
		mapWidth, mapHeight := screen.Size()
		mapHeight = mapHeight - 1 // For our results line
		maxRooms, roomMinSize, roomMaxSize := 30, 6, 10
		maxMonstersPerRoom := 3
		gameState := PlayerTurnState
		lastGameState := gameState

		fovRadius := 10

		colors := map[string]goro.Color{
			"darkWall":    goro.Color{R: 25, G: 25, B: 25, A: 255},
			"darkGround":  goro.Color{R: 100, G: 100, B: 100, A: 255},
			"lightWall":   goro.Color{R: 50, G: 50, B: 50, A: 255},
			"lightGround": goro.Color{R: 150, G: 150, B: 150, A: 255},
		}

		player := entity.NewPlayerEntity()

		entities := []interfaces.Entity{
			player,
		}

		gameMap := mapping.NewGameMap(mapWidth, mapHeight)
		gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, &entities, maxMonstersPerRoom)

		fovRecompute := true

		fovMap := InitializeFoV(gameMap)

		var turnResults []interfaces.Result
		maxResults := 1

		for {
			if fovRecompute {
				RecomputeFoV(fovMap, player.X(), player.Y(), fovRadius, fov.Light{})
			}

			// Draw screen.
			DrawAll(screen, entities, gameMap, fovMap, fovRecompute, colors)

			DrawResults(screen, turnResults, maxResults)

			screen.Flush()

			fovRecompute = false

			ClearAll(screen, entities, fovMap)

			ClearResults(screen, turnResults, maxResults)

			// Adjust our game state to allow for processing of results that we cannot see yet due to screen limitations.
			if gameState == ProcessResultsState {
				gameState = lastGameState
			}

			if len(turnResults) > maxResults {
				if gameState != ProcessResultsState {
					lastGameState = gameState
					gameState = ProcessResultsState
				}
			} else {
				turnResults = nil
			}

			// Clear out all Seen results
			i := 0
			for _, r := range turnResults {
				switch result := r.(type) {
				case *results.Message:
					if !result.Seen {
						turnResults[i] = r
						i++
					}
				default:
					turnResults[i] = r
					i++
				}
			}
			turnResults = turnResults[:i]

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					if gameState == PlayerTurnState {
						x := player.X() + action.X
						y := player.Y() + action.Y
						if !gameMap.IsBlocked(x, y) {
							otherEntity := entity.FindEntityAtLocation(entities, x, y, entity.BlockMovement, entity.BlockMovement)
							if otherEntity != nil {
								turnResults = append(turnResults, &results.Message{
									Text: fmt.Sprintf("You lick the %s in the shins, much to its enjoyment!\n", otherEntity.Name()),
								})
							} else {
								player.Move(action.X, action.Y)
								fovRecompute = true
							}
						}
						gameState = NPCTurnState
					}
				case ActionQuit:
					goro.Quit()
				}
			case goro.EventQuit:
				return
			}

			// Handle entity updates.
			if gameState == NPCTurnState {
				for _, e := range entities {
					if e.Actor() != nil {
						// Target the player and take a turn.
						e.Actor().SetTarget(entities[0])
						turnResults = append(turnResults, e.Actor().TakeTurn(fovMap, gameMap, entities)...)
					}
				}
				gameState = PlayerTurnState
			}

		}
	})
}
