package gameobject

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	gameObject
	speedx float64
	speedy float64
}

func NewBall(screenWidth, screenHeight float64) Ball {
	img, _, err := ebitenutil.NewImageFromFile("Images/volleyball64.png")
	if err != nil {
		log.Fatal(err)
	}

	var b Ball
	b.image = img
	b.width = img.Bounds().Dx()
	b.height = img.Bounds().Dy()
	b.SetX(float64(screenWidth-float64(b.width))/2)
	b.SetY(screenHeight-float64(b.height))
	b.SetSpeed(3, -3)
	
	return b
}

func (ball *Ball) Reset(screenWidth, screenHeight float64) {
	ball.SetX(screenWidth / 2)
	ball.SetY(screenHeight / 2)
	ball.SetSpeed(3, -3)
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
	if ball.x< 0 || ball.x+float64(ball.width) > float64(screenWidth) {
		ball.speedx = -ball.speedx
	}
	if ball.y < 0 {
		ball.speedy = -ball.speedy
	}

	// 檢查球是否碰到饒
	// if ball.y+float64(ball.height) > jao.y && ball.x > jao.x-float64(jao.width)/2 && ball.x < jao.x+float64(jao.width)/2 {
	// 	ball.speedy = -ball.speedy
	// }
	// 檢查球是否碰到饒
	if ball.y+float64(ball.height) > jao.y  {
		ball.speedy = -ball.speedy
	}

}
func (ball *Ball) Draw(screen *ebiten.Image, ScreenWidth, ScreenHeight int) {
	op := &ebiten.DrawImageOptions{}
	
	op.GeoM.Translate(ball.x, ball.y)
	screen.DrawImage(ball.image, op)

}
