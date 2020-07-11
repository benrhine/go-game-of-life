package models

import (
	"bytes"
)

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	A, B *CytoGrid
	W, H int
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
	// Update the state of the next CytoGrid (b) from the current CytoGrid (a).
	for y := 0; y < l.H; y++ {
		for x := 0; x < l.W; x++ {
			l.B.Set(x, y, l.A.Next(x, y))
		}
	}
	// Swap CytoGrids a and b.
	l.A, l.B = l.B, l.A
}

// String returns the game board as a string.
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.H; y++ {
		for x := 0; x < l.W; x++ {
			b := byte('.')
			if l.A.Alive(x, y) {
				b = '0'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
