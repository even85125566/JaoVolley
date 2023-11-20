package main

import (
	"image/color"
	gamemanage "jaovolleyball/GameManage"
	gameobject "jaovolleyball/GameObject"
	input "jaovolleyball/Input"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	Input input.Input
	Jao   gameobject.Jao
	Ball  gameobject.Ball
	Mod   gamemanage.GameMod
}

func (g *Game) Update() error {
	switch g.Mod {
	case gamemanage.GameTitle:

	case gamemanage.Gaming:

		// 檢查球是否碰到底部，如果碰到，遊戲結束
		g.Jao.Update(screenHeight)
		g.Ball.Update(screenWidth, screenHeight, &g.Jao)
		
		if g.Ball.Y()+g.Ball.Size/2 > float64(screenHeight) {
			g.Mod = gamemanage.GameOver
		}
	case gamemanage.GameOver:
		//按空白建則繼續遊戲
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Restart()
		}
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 繪製背景
	screen.Fill(color.RGBA{135, 206, 250, 255})

	// 繪製球
	g.Ball.Draw(screen)
	//繪製饒
	g.Jao.Draw(screen)

	// 如果遊戲結束，顯示 Game Over
	if g.Mod == gamemanage.GameOver {
		ebitenutil.DebugPrint(screen, "Game Over\nPress Enter to restart")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Restart() {
	g.Jao.SetX(screenWidth / 2)
	g.Jao.SetY(screenHeight - g.Jao.Size)
	g.Ball.SetX(screenWidth / 2)
	g.Ball.SetY(screenHeight / 2)
	g.Ball.SetSpeed(3, -3)
	g.Mod = gamemanage.Gaming

}

func main() {
	b := gameobject.NewBall(screenWidth, screenHeight)
	j := gameobject.NewJao(screenWidth, screenHeight)

	game := &Game{
		Jao:  j,
		Ball: b,
	}
	game.Restart()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Jao Volleyball")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
