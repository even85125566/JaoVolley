package gameobject

type VolleyNet struct {
	GameObject
}


func NewVolleyNet(screenWidth, screenHeight float64) VolleyNet {

	var v VolleyNet

	v.width = 10
	v.height = 170
	v.SetX(float64(screenWidth-float64(v.width)) / 2)
	v.SetY(screenHeight - float64(v.height))
	return v
}
