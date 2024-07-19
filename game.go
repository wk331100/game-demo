package main

import (
	"fmt"
	"github.com/game-demo/game-demo/config"
	"github.com/game-demo/game-demo/model"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	input *Input
	ship  *model.Ship
	cfg   *config.Config
}

func NewGame() *Game {
	cfg := config.LoadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	return &Game{
		input: &Input{},
		ship:  model.NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		cfg:   cfg,
	}
}

func (g *Game) Update() error {
	g.input.Update(g.ship, g.cfg)
	fmt.Println("update")
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen, g.cfg)
	fmt.Println("1")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}
