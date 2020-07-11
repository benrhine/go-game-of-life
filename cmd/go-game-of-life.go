package main

import (
	"benrhine.com/go-game-of-life/internal/app/go-game-of-life/models"
	"benrhine.com/go-game-of-life/internal/app/go-game-of-life/services"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type CytoGrid = models.CytoGrid
type Life = models.Life

// NewCytoGrid returns an empty CytoGrid of the specified width and height.
func NewCytoGrid(w, h int) *CytoGrid {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &CytoGrid{S: s, W: w, H: h}
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(w, h int) *Life {
	a := NewCytoGrid(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		A: a, B: NewCytoGrid(w, h),
		W: w, H: h,
	}
}

func main() {
	// Flag info from here https://gobyexample.com/command-line-flags
	width := flag.Int("width", 10, "Width of CytoGrid")
	height := flag.Int("height", 10, "Height of CytoGrid")

	flag.Parse()  // without this line it wont set the new grid values

	fmt.Printf("Setting incoming CytoGrid to %d by %d\n", *width, *height)

	// What do the * do? The stars are de-referencing the variable so you get the actual value and not the memory address
	l := NewLife(*width, *height)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print("\x0c", l) // Clear screen and print CytoGrid.
		time.Sleep(time.Second / 30)
	}
}
