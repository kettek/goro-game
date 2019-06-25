package main

type GameMap struct {
	Width, Height int
	Tiles         [][]Tile
}

func (g *GameMap) Initialize() {
	g.Tiles = make([][]Tile, g.Width)
	for x := range g.Tiles {
		g.Tiles[x] = make([]Tile, g.Height)
	}

	g.Tiles[30][22].BlockMovement = true
	g.Tiles[30][22].BlockSight = true
	g.Tiles[31][22].BlockMovement = true
	g.Tiles[31][22].BlockSight = true
	g.Tiles[32][22].BlockMovement = true
	g.Tiles[32][22].BlockSight = true
}

// IsBlocked returns if the given coordinates are blocking.
func (g *GameMap) IsBlocked(x, y int) bool {
	return g.Tiles[x][y].BlockMovement
}
