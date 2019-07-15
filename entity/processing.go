package entity

import (
  "github.com/kettek/goro-game/interfaces"
)

// FindEntityAtLocation finds and returns the first entity at x and y matching the provided flags. If none exists, it returns nil.
func FindEntityAtLocation(entities []interfaces.Entity, x, y int, checkMask int, matchFlags int) interfaces.Entity {
	for _, e := range entities {
		if (e.Flags()&checkMask) == matchFlags && e.X() == x && e.Y() == y {
			return e
		}
	}
	return nil
}
