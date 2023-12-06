package gameobject

import (
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	gameObject
	canBeCollided bool
	collidedTime  time.Time
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
	b.SetX(float64(screenWidth-float64(b.width)) / 2)
	b.SetY(screenHeight - float64(b.height))
	b.SetSpeed(6)
	b.direction = rand.Float64() * 2 * math.Pi

	return b
}

func (ball *Ball) Reset(screenWidth, screenHeight float64) {
	ball.SetX(screenWidth / 2)
	ball.SetY(screenHeight / 2)
	ball.SetSpeed(6)
	ball.direction = rand.Float64() * 2 * math.Pi
}
func (ball *Ball) SetSpeed(speed float64) {
	ball.speed = speed
}

func (ball *Ball) Update(screenWidth, screenHeight int, jao *Jao) {
	// 移動球
	ball.x += ball.speed * math.Cos(ball.direction)
	ball.y += ball.speed * math.Sin(ball.direction)
	//檢查球的碰撞時間
	if !ball.canBeCollided {
		if time.Now().After(ball.collidedTime.Add(time.Millisecond * 1 * 500)) {
			ball.canBeCollided = true
		}
	}
	// 檢查球是否碰到畫面邊緣
	if ball.x < 0 || ball.x+float64(ball.width) > float64(screenWidth) {
		ball.direction = math.Pi - ball.direction
	}
	if ball.y < 0 {
		ball.direction = -ball.direction
	}
	// 檢查球是否碰到饒
	// 需檢查三面
	switch {
	// 頭
	case ball.y+float64(ball.height) > jao.y && ball.x > jao.x && ball.x < jao.x+float64(jao.width):
		if ball.canBeCollided {
			ball.direction = -ball.direction
			ball.canBeCollided = false
			ball.collidedTime = time.Now()
		}
	}

}
func (ball *Ball) Draw(screen *ebiten.Image, ScreenWidth, ScreenHeight int) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(ball.x, ball.y)
	screen.DrawImage(ball.image, op)

}
