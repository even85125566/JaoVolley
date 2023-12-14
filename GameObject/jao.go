package gameobject

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Jao struct {
	GameObject
	jumpspeed float64
	IsJumping bool
}

func NewJao(screenWidth, screenHeight float64) Jao {
	img, _, err := ebitenutil.NewImageFromFile("Images/stickJaoleft.png")
	if err != nil {
		log.Fatal(err)
	}
	var j Jao
	j.image = img
	j.width = img.Bounds().Dx()
	j.height = img.Bounds().Dy()
	j.Reset(screenWidth, screenHeight)
	return j
}
func (jao *Jao) Reset(screenWidth, screenHeight float64) {
	jao.SetX(screenWidth / 2)
	jao.SetY(screenHeight - float64(jao.Height()))
	jao.SetSpeed(4, 4)
	jao.SetGravity(0.5)
	jao.jumpspeed = 8
	jao.IsJumping = false
}
func (jao *Jao) Update(screenHeight float64) {
	// 控制饒的移動
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		jao.x -= jao.speedx
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		jao.x += jao.speedx
	}
	//按下方向鍵上 啟動跳躍
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && !jao.IsJumping {
		jao.IsJumping = true
		jao.speedy = -jao.jumpspeed
	}
	//跳躍判斷
	if jao.IsJumping {
		jao.y += jao.speedy
		//給予掉落效果 模擬重力
		jao.speedy += jao.gravity
		if jao.y > screenHeight-float64(jao.height) {
			jao.y = screenHeight - float64(jao.height)
			jao.IsJumping = false
		}

	}

}
func (jao *Jao) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(jao.x, jao.y)
	screen.DrawImage(jao.image, op)
}
