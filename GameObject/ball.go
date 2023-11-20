package gameobject

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	gameObject
	speedx float64
	speedy float64
	Size   float64
}

func NewBall(screenWidth, screenHeight float64) Ball {
	var b Ball
	b.SetX(screenWidth / 2)
	b.SetY(screenHeight / 2)
	b.SetSpeed(1, -1)
	b.Size = 16
	return b
}
func (ball *Ball) SetSpeed(x, y float64) {
	ball.speedx = x
	ball.speedy = y
}

func (ball *Ball) Update(screenWidth, screenHeight int, jao *Jao) {

	// 移動球
	ball.x += ball.speedx
	ball.y += ball.speedy

	// 檢查球是否碰到畫面邊緣
	if ball.x-ball.Size/2 < 0 || ball.x+ball.Size/2 > float64(screenWidth) {
		ball.speedx = -ball.speedx
	}
	if ball.y-ball.Size/2 < 0 {
		ball.speedy = -ball.speedy
	}

	// 檢查球是否碰到饒
	//
	if ball.y+ball.Size/2 > jao.y && ball.x > jao.x-jao.Size/2 && ball.x < jao.x+jao.Size/2 {
		ball.speedy = -ball.speedy
	}

}
func (ball *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(ball.x-ball.Size/2), float32(ball.y-ball.Size/2), float32(ball.Size), float32(ball.Size), color.RGBA{255, 0, 0, 255}, true)

}
