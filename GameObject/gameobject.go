package gameobject

import "github.com/hajimehoshi/ebiten/v2"

type GameObject struct {
	width   int
	height  int
	x       float64
	y       float64
	speedx  float64
	speedy  float64
	gravity float64
	image   *ebiten.Image
}

func (gameObject *GameObject) Width() int {
	return gameObject.width
}
func (gameObject *GameObject) Height() int {
	return gameObject.height
}
func (gameObject *GameObject) X() float64 {
	return gameObject.x
}
func (gameObject *GameObject) Y() float64 {
	return gameObject.y
}

func (gameObject *GameObject) SetX(x float64) {
	gameObject.x = x
}
func (gameObject *GameObject) SetY(y float64) {
	gameObject.y = y
}
func (gameObject *GameObject) SetSpeed(speedx, speedy float64) {
	gameObject.speedx = speedx
	gameObject.speedy = speedy
}
func (gameObject *GameObject) SetGravity(gravity float64) {
	gameObject.gravity = gravity
}

//RightSide 供Overlap函數使用
func (gameObject *GameObject) RightSide() GameObject {
	var newObject GameObject
	newObject = *gameObject
	newObject.x = gameObject.x + float64(gameObject.width/2)
	newObject.width = gameObject.width / 2
	return newObject

}

//LeftSide 供Overlap函數使用
func (gameObject *GameObject) LeftSide() GameObject {
	var newObject GameObject
	newObject = *gameObject
	newObject.width = gameObject.width / 2
	return newObject

}

//TopSideY 取得物件上半部的最底端Y
func (gameObject *GameObject) TopSideY() float64 {

	return gameObject.y + float64(gameObject.height*2/3)

}

//TopSideY 取得物件上半部的最底端Y
func (gameObject *GameObject) BottomSideY() float64 {

	return gameObject.y + float64(gameObject.height)

}
