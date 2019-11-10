package entity

import (
	"time"
	"github.com/kettek/goro-game/interfaces"
)

type Message struct {
	when time.Time
	from interfaces.Entity
	to   interfaces.Entity
  text string
}
