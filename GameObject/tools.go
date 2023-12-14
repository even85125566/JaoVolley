package gameobject


func IsOverlap(o1, o2 GameObject) bool {
	xRange1 := []float64{o1.X(), o1.X() + float64(o1.Width())}
	xRange2 := []float64{o2.X(), o2.X() + float64(o2.Width())}

	// 排序 x 坐标范围
	if xRange1[0] > xRange1[1] {
		xRange1[0], xRange1[1] = xRange1[1], xRange1[0]
	}
	if xRange2[0] > xRange2[1] {
		xRange2[0], xRange2[1] = xRange2[1], xRange2[0]
	}

	// 检查两个 x 坐标范围是否有重叠
	if xRange1[1] >= xRange2[0] && xRange2[1] >= xRange1[0] {
		return true
	}
	return false
}
