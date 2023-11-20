package gameobject
type Ball struct {
	gameObject
	speedx float64
	speedy float64
	Size   float64
}

func (ball *Ball) Update(screenWidth,screenHeight int,jao *Jao) {
	//debug cfg
	ball.Size=16
	// 移動球
	ball.x += ball.speedx
	ball.y += ball.speedy

	// 檢查球是否碰到畫面邊緣
	if ball.x-ball.Size/2 < 0 || ball.x+ball.Size/2 > float64(screenWidth) {
		ball.speedx = -ball.speedx
	}
	if ball.y-ball.Size/2 < 0 {
		ball.speedy = -ball.speedy
	}

	// 檢查球是否碰到皮卡丘
	if g.ballY+ballSize/2 > g.pikachuY && g.ballX > g.pikachuX-pikachuSize/2 && g.ballX < g.pikachuX+pikachuSize/2 {
		g.ballSpeedY = -g.ballSpeedY
	}

	// 檢查球是否碰到底部，如果碰到，遊戲結束
	if g.ballY+ballSize/2 > screenHeight {
		g.gameOver = true
	}
}
func (ball *Ball) Draw() {

}
