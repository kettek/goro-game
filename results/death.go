package results

import (
	"github.com/kettek/goro-game/interfaces"
)

type Death struct {
	Killer interfaces.Entity
	Killed interfaces.Entity
}
