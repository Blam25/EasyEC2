package EasyEC2

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var Systems []func()
var DrawSystems []func(screen *ebiten.Image)

type Game struct{}

func (g *Game) Update() error {

	for _, s := range Systems {
		s()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, s := range DrawSystems {
		s(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
