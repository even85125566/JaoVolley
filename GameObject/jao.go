package gameobject

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Jao struct {
	gameObject
	OriginalY float64
	speedx    float64
	speedy    float64
	jumpspeed float64
	Size      float64
	IsJumping bool
}

func NewJao(screenWidth, screenHeight float64) Jao {
	var j Jao
	j.Size = 64
	j.SetX(screenWidth / 2)
	j.SetY(screenHeight - j.Size)
	j.speedx = 4
	j.speedy = 4
	j.jumpspeed = 8
	j.IsJumping = false
	return j
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
		if jao.y > screenHeight-jao.Size {
			jao.y = screenHeight - jao.Size
			jao.IsJumping = false
		}

	}

}
func (jao *Jao) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(jao.x-jao.Size/2), float32(jao.y-jao.Size/2), float32(jao.Size), float32(jao.Size), color.RGBA{255, 215, 0, 255}, true)

}
