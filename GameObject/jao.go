package gameobject

import (
	"bytes"
	resources "jaovolleyball/Resources"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Jao struct {
	GameObject
	Type      Type
	jumpspeed float64
	IsJumping bool
}
type Type int

const (
	Left  Type = 0
	Right Type = 1
)

func NewJao(jaoType Type, screenWidth, screenHeight float64) Jao {

	var j Jao
	switch jaoType {
	case Left:
		img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.StickJaoleft))
		if err != nil {
			log.Fatal(err)
		}
		j.image = img

	case Right:
		img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.StickJaoright))
		if err != nil {
			log.Fatal(err)
		}
		j.image = img

	default:
		j.image = nil
	}
	j.Type = jaoType
	j.width = j.image.Bounds().Dx()
	j.height = j.image.Bounds().Dy()
	j.Reset(screenWidth, screenHeight)
	return j
}
func (jao *Jao) Reset(screenWidth, screenHeight float64) {
	switch jao.Type {
	case Left:
		jao.SetX(0)
	case Right:
		jao.SetX(screenWidth - float64(jao.width))
	default:
		jao.SetX(screenWidth / 2)

	}
	log.Printf("reset饒方向%v", jao.Type)

	jao.SetY(screenHeight - float64(jao.Height()))
	jao.SetSpeed(4, 4)
	jao.SetGravity(0.5)
	jao.jumpspeed = 8
	jao.IsJumping = false
}
func (jao *Jao) Update(screenHeight float64) {
	switch jao.Type {
	case Left:
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

	case Right:
		// 控制饒的移動
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			jao.x -= jao.speedx
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			jao.x += jao.speedx
		}
		//按下方向鍵上 啟動跳躍
		if ebiten.IsKeyPressed(ebiten.KeyW) && !jao.IsJumping {
			jao.IsJumping = true
			jao.speedy = -jao.jumpspeed
		}
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
