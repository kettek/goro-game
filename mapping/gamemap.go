package mapping

import (
	"github.com/kettek/goro"
	"github.com/kettek/goro/pathing"
	"github.com/kettek/goro-game/entity"
	"github.com/kettek/goro-game/interfaces"
)

// GameMap is our map type for holding our tiles and dimensions.
type GameMap struct {
	width, height int
	Tiles         [][]Tile
}

func NewGameMap(width, height int) interfaces.GameMap {
	g := &GameMap{
		width: width,
		height: height,
	}
	g.Tiles = make([][]Tile, g.width)

	for x := range g.Tiles {
		g.Tiles[x] = make([]Tile, g.height)
		for y := range g.Tiles[x] {
			g.Tiles[x][y] = Tile{
        Flags: BlockSight | BlockMovement,
			}
		}
	}

	return g
}

// MakeMap creates a new randomized map. This is built according to the passed arguments.
func (g *GameMap) MakeMap(maxRooms, roomMinSize, roomMaxSize int, entities *[]interfaces.Entity, maxMonsters int) {
	var rooms []Rect

	for r := 0; r < maxRooms; r++ {
		// Generate a random width and height.
		width := roomMinSize + goro.Random.Intn(roomMaxSize)
		height := roomMinSize + goro.Random.Intn(roomMaxSize)
		// Generate a random position within the map boundaries.
		x := goro.Random.Intn(g.width - width - 1)
		y := goro.Random.Intn(g.height - height - 1)
		// Create a Rect according to our generated sizes.
		room := NewRect(x, y, width, height)

		// Iterate through our existing rooms to check for intersection with our new room.
		intersects := false
		for _, otherRoom := range rooms {
			if room.Intersect(otherRoom) {
				intersects = true
				break
			}
		}
		// Add the room if there is no intersection found.
		if !intersects {
			g.CreateRoom(room)
			roomCenterX, roomCenterY := room.Center()

			// Always place the player in the center of the first room.
			if len(rooms) == 0 {
				(*entities)[0].SetX(roomCenterX)
				(*entities)[0].SetY(roomCenterY)
			} else {
				prevCenterX, prevCenterY := rooms[len(rooms)-1].Center()

				// Flip a coin if we should tunnel horizontally or vertically first.
				if goro.Random.Intn(1) == 1 {
					g.CreateHTunnel(prevCenterX, roomCenterX, prevCenterY)
					g.CreateVTunnel(prevCenterY, roomCenterY, roomCenterX)
				} else {
					g.CreateVTunnel(prevCenterY, roomCenterY, prevCenterX)
					g.CreateHTunnel(prevCenterX, roomCenterX, roomCenterY)
				}
			}
			// Place random monsters in the room.
			g.PlaceEntities(room, entities, maxMonsters)
			// Append our new room to our rooms list.
			rooms = append(rooms, room)
		}
	}
}

// CreateRoom creates a room from a provided rect.
func (g *GameMap) CreateRoom(r Rect) {
	for x := r.X1 + 1; x < r.X2; x++ {
		for y := r.Y1 + 1; y < r.Y2; y++ {
			if g.InBounds(x, y) {
				g.Tiles[x][y] = Tile{}
			}
		}
	}
}

// CreateHTunnel creates a horizontal tunnel from x1 to/from x1 starting at y.
func (g *GameMap) CreateHTunnel(x1, x2, y int) {
	for x := goro.MinInt(x1, x2); x <= goro.MaxInt(x1, x2); x++ {
		if g.InBounds(x, y) {
			g.Tiles[x][y] = Tile{}
		}
	}
}

// CreateVTunnel creates a vertical tunnel from y1 to/from y2 starting at x.
func (g *GameMap) CreateVTunnel(y1, y2, x int) {
	for y := goro.MinInt(y1, y2); y <= goro.MaxInt(y1, y2); y++ {
		if g.InBounds(x, y) {
			g.Tiles[x][y] = Tile{}
		}
	}
}

// PlaceEntities places 0 to maxMonsters monster entities in the provided room.
func (g *GameMap) PlaceEntities(room Rect, entities *[]interfaces.Entity, maxMonsters int) {
	monstersCount := goro.Random.Intn(maxMonsters)

	for i := 0; i < monstersCount; i++ {
		var monster interfaces.Entity
		// Acquire a random location within the room.
		x := (1 + room.X1) + goro.Random.Intn(room.X2-room.X1-1)
		y := (1 + room.Y1) + goro.Random.Intn(room.Y2-room.Y1-1)

		if entity.FindEntityAtLocation(*entities, x, y, 0, 0) == nil {
			// Generate an orc with 80% probability or a troll with 20%.
			if goro.Random.Intn(100) < 80 {
				monster = entity.NewMonsterEntity(x, y, 'o', goro.Style{Foreground: goro.ColorLime}, "Orc")
			} else {
				monster = entity.NewMonsterEntity(x, y, 'T', goro.Style{Foreground: goro.ColorGreen}, "Troll")
			}
			*entities = append(*entities, monster)
		}
	}
}

// Explored returns if the tile at x by y has been explored.
func (g *GameMap) Explored(x, y int) bool {
	if g.InBounds(x, y) {
		return g.Tiles[x][y].Flags&Explored == Explored
	}
	return false
}

// SetExplored sets the explored state of the tile at x and y to the passed explored bool.
func (g *GameMap) SetExplored(x, y int, explored bool) {
	if g.InBounds(x, y) {
    if explored {
      g.Tiles[x][y].Flags = g.Tiles[x][y].Flags | Explored
    } else {
      g.Tiles[x][y].Flags = g.Tiles[x][y].Flags &^ Explored
    }
	}
}

// IsBlocked returns if the given coordinates are blocking.
func (g *GameMap) IsBlocked(x, y int) bool {
  if g.InBounds(x, y) {
	  return g.Tiles[x][y].Flags&BlockMovement == BlockMovement
  }
  return true
}

// IsOpaque returns if the given coordinates block sight.
func (g *GameMap) IsOpaque(x, y int) bool {
  if g.InBounds(x, y) {
		return g.Tiles[x][y].Flags&BlockSight == BlockSight
  }
  return true
}

// Width returns the width of the game map.
func (g *GameMap) Width() int {
	return g.width
}

// Height returns the height of the game map.
func (g *GameMap) Height() int {
	return g.height
}

// InBounds returns if the given coordinates are within the map's boundaries.
func (g *GameMap) InBounds(x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return false
	}
	return true
}

// HEHE
func (g *GameMap) CostAt(x, y int) int {
	if g.IsBlocked(x, y) {
		return pathing.MaximumCost
	}
	return 0
}