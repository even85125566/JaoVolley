package gameobject

import "github.com/hajimehoshi/ebiten/v2"

type Jao struct {
	gameObject
	speed float64
}

func (jao *Jao) Update() {
	// 控制皮卡丘的移動
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		jao.x -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		jao.x += 4
	}

}
func (jao *Jao) Draw() {

}
