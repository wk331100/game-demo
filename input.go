package main

import (
	"fmt"
	"github.com/game-demo/game-demo/config"
	"github.com/game-demo/game-demo/model"
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	msg string
}

func (i *Input) Update(ship *model.Ship, cfg *config.Config) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ship.X -= cfg.ShipSpeedFactor
		fmt.Println("<<<<<<")
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ship.X += cfg.ShipSpeedFactor
		fmt.Println(">>>>>")
	}
}
