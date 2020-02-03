package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Position : Positon on the board
type Position struct {
	X float64
	Y float64
}

// SQUARE_SIZE
const SQUARE_SIZE = 16

var square *ebiten.Image
var pos Position

func debug(screen *ebiten.Image, value string) {
	ebitenutil.DebugPrint(screen, value)

}

func update(screen *ebiten.Image) error {
	/** Game state updates */

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	/** Rendering */

	// Fill the screen with #FF0000
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})

	if square == nil {
		// Create an 16x16 image
		square, _ = ebiten.NewImage(SQUARE_SIZE, SQUARE_SIZE, ebiten.FilterNearest)
	}

	// Fill the square with the white color
	square.Fill(color.White)

	opts := &ebiten.DrawImageOptions{}

	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(pos.X, pos.Y)

	screen.DrawImage(square, opts)

	// Keyboard event
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		pos.Y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		pos.Y++
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		pos.X++
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		pos.X--
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// Move the square to the cursor (aka drag n drop)
		pos.X = float64(x)
		pos.Y = float64(y)

	}

	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
