package gameobject

import "github.com/hajimehoshi/ebiten/v2"

type gameObject struct {
	width     int
	height    int
	x         float64
	y         float64
	speed     float64
	direction float64
	image     *ebiten.Image
}

func (gameObject *gameObject) Width() int {
	return gameObject.width
}
func (gameObject *gameObject) Height() int {
	return gameObject.height
}
func (gameObject *gameObject) X() float64 {
	return gameObject.x
}
func (gameObject *gameObject) Y() float64 {
	return gameObject.y
}

func (gameObject *gameObject) SetX(x float64) {
	gameObject.x = x
}
func (gameObject *gameObject) SetY(y float64) {
	gameObject.y = y
}
