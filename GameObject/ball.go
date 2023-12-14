package gameobject

import (
	_ "image/png"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	GameObject
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
	b.SetSpeed(6, 6)

	return b
}

func (ball *Ball) Reset(screenWidth, screenHeight float64) {
	ball.SetX(0)
	ball.SetY(10)
	ball.SetSpeed(3, 3)
	ball.SetGravity(0.05)
}
func (ball *Ball) BeCollided() bool {
	return ball.canBeCollided
}

func (ball *Ball) Update(screenWidth, screenHeight int, jao []Jao) {

	// 移動球
	ball.x += ball.speedx
	ball.y += ball.speedy
	//處理重力
	ball.speedy += ball.gravity
	//檢查球的碰撞時間
	if !ball.canBeCollided {
		if time.Now().After(ball.collidedTime.Add(time.Millisecond * 1 * 500)) {
			ball.canBeCollided = true
		}
	}
	// 檢查球是否碰到畫面邊緣
	if ball.x < 0 || ball.x+float64(ball.width) > float64(screenWidth) {
		ball.speedx = -ball.speedx
	}
	if ball.y < 0 {
		ball.speedy = math.Abs(ball.speedy)
	}
	// 檢查球是否碰到饒
	//TODO:身體下半部處理

	for i := 0; i < len(jao); i++ {
		
		if ball.canBeCollided && ball.y+float64(ball.height) > jao[i].y {

			switch {
			// 左半邊
			case IsOverlap(jao[i].LeftSide(), ball.GameObject):
				ball.speedx = -math.Abs(ball.speedx)
				ball.speedy = -ball.speedy
				ball.canBeCollided = false
				ball.collidedTime = time.Now()
	
				//右半邊
			case IsOverlap(jao[i].RightSide(), ball.GameObject):
				ball.speedx = math.Abs(ball.speedx)
				ball.speedy = -ball.speedy
				ball.canBeCollided = false
				ball.collidedTime = time.Now()
	
			default:
	
			}
	}
	

	}

}
func (ball *Ball) Draw(screen *ebiten.Image, ScreenWidth, ScreenHeight int) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(ball.x, ball.y)
	screen.DrawImage(ball.image, op)

}
