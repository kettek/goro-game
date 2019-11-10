package interfaces

import (
	"time"
)

type Message interface {
	Text() string
	When() time.Time
	From() Entity
	To() Entity
}
