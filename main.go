package main

import (
	"image/color"
	gamemanage "jaovolleyball/GameManage"
	gameobject "jaovolleyball/GameObject"
	input "jaovolleyball/Input"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	pikachuSize  = 64
)

type Game struct {
	Input                  input.Input
	Jao                    gameobject.Jao
	Ball                   gameobject.Ball
	Mod                    gamemanage.GameMod
	pikachuX, pikachuY     float64
	ballX, ballY           float64
	ballSpeedX, ballSpeedY float64
	gameOver               bool
}

func (g *Game) Update() error {
	switch g.Mod {
	case gamemanage.GameTitle:

	case gamemanage.Gaming:

	case gamemanage.GameOver:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Restart()
		}
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 繪製背景
	screen.Fill(color.RGBA{135, 206, 250, 255}) // 淡藍色

	vector.DrawFilledRect(screen, float32(g.pikachuX-pikachuSize/2), float32(g.pikachuY-pikachuSize/2), pikachuSize, pikachuSize, color.RGBA{255, 215, 0, 255}, true)
	// 繪製球
	vector.DrawFilledRect(screen, float32(g.ballX-ballSize/2), float32(g.ballY-ballSize/2), ballSize, ballSize, color.RGBA{255, 0, 0, 255}, true)

	// 如果遊戲結束，顯示 Game Over
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over\nPress Enter to restart")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Restart() {
	g.pikachuX = screenWidth / 2
	g.pikachuY = screenHeight - pikachuSize
	g.ballX = screenWidth / 2
	g.ballY = screenHeight / 2
	g.ballSpeedX = 3
	g.ballSpeedY = -3
	g.gameOver = false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{}
	game.Restart()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pikachu Volleyball")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
