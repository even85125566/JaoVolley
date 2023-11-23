package gameobject

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Jao struct {
	gameObject
	speedx    float64
	speedy    float64
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
	j.SetX(float64(screenWidth-float64(j.width)) / 2)
	j.SetY(screenHeight - float64(j.height))
	j.speedx = 4
	j.speedy = 4
	j.jumpspeed = 8
	j.IsJumping = false
	return j
}
func (jao *Jao) Reset(screenWidth, screenHeight float64) {
	jao.SetX(screenWidth / 2)
	jao.SetY(screenHeight - float64(jao.Height()))
	
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
		jao.speedy += 0.5
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
