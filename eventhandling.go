package main

import (
	"github.com/kettek/goro"
)

// handleKeyEvent converts a KeyEvent into a corresponding Action.
func handleKeyEvent(ev goro.EventKey) Action {
	switch ev.Key {
	case goro.KeyUp, goro.KeyK:
		return ActionMove{Y: -1}
	case goro.KeyDown, goro.KeyJ:
		return ActionMove{Y: 1}
	case goro.KeyRight, goro.KeyL:
		return ActionMove{X: 1}
	case goro.KeyLeft, goro.KeyH:
		return ActionMove{X: -1}
	case goro.KeyEscape:
		return ActionQuit{}
	}
	return nil
}
