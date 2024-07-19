package model

import (
	"fmt"
	"github.com/game-demo/game-demo/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./image/ship1.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		Image:  img,
		Width:  width,
		Height: height,
		X:      float64(screenWidth-width) / 2,
		Y:      float64(screenHeight - height),
	}

	fmt.Println(ship.X)
	return ship
}

type Ship struct {
	Image  *ebiten.Image
	Width  int
	Height int
	X      float64 // x坐标
	Y      float64 // y坐标
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *config.Config) {
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(float64(cfg.ScreenWidth-ship.Width)/2, float64(cfg.ScreenHeight-ship.Height))
	op.GeoM.Translate(ship.X, ship.Y)
	screen.DrawImage(ship.Image, op)
}
