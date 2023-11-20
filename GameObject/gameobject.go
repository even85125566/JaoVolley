package gameobject

type gameObject struct {
	width  int
	height int
	x      float64
	y      float64
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
