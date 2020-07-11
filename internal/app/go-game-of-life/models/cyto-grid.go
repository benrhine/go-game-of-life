package models

// CytoGrid represents a two-dimensional CytoGrid of cells.
type CytoGrid struct {
	S    [][]bool
	W, H int
}

// Set sets the state of the specified cell to the given value.
func (f *CytoGrid) Set(x, y int, b bool) {
	f.S[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the x or y coordinates are outside the CytoGrid boundaries they are wrapped
// toroidally. For instance, an x value of -1 is treated as width-1.
func (f *CytoGrid) Alive(x, y int) bool {
	x += f.W
	x %= f.W
	y += f.H
	y %= f.H
	return f.S[y][x]
}

// Next returns the state of the specified cell at the next time step.
func (f *CytoGrid) Next(x, y int) bool {
	// Count the adjacent cells that are alive.
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	// Return next state according to the game rules:
	//   exactly 3 neighbors: on,
	//   exactly 2 neighbors: maintain current state,
	//   otherwise: off.
	return alive == 3 || alive == 2 && f.Alive(x, y)
}