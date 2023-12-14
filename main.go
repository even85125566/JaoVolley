package main

import (
	"fmt"
	gamelog "jaovolleyball/GameLog"
	gamemanage "jaovolleyball/GameManage"
	gameobject "jaovolleyball/GameObject"
	input "jaovolleyball/Input"
	socket "jaovolleyball/Socket"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gravity      = 0.5
)

type Game struct {
	Input      input.Input
	Jao        []gameobject.Jao
	Ball       gameobject.Ball
	Mod        gamemanage.GameMod
	Network    *socket.Network
	Background *ebiten.Image
	ErrMsg     []string
}

func (g *Game) Update() error {
	switch g.Mod {
	case gamemanage.GameTitle:

	case gamemanage.Gaming:

		// 檢查球是否碰到底部，如果碰到，遊戲結束
		for i := 0; i < len(g.Jao); i++ {
			g.Jao[i].Update(screenHeight)
		}
		g.Ball.Update(screenWidth, screenHeight, g.Jao)

		if g.Ball.Y()+float64(g.Ball.Height()) > float64(screenHeight) {
			g.Mod = gamemanage.GameOver
			// err := g.Network.SendMessage("god damn")
			// if err != nil {
			// 	g.ErrMsg = append(g.ErrMsg, err.Error())
			// }
		}
	case gamemanage.GameOver:
		//按空白建則繼續遊戲
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Restart()
		}
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 繪製背景
	screen.DrawImage(g.Background, nil)
	// 繪製球
	g.Ball.Draw(screen, screenWidth, screenHeight)
	//繪製饒
	for i := 0; i < len(g.Jao); i++ {
		g.Jao[i].Draw(screen)
	}
	if len(g.ErrMsg) != 0 {
		errMsg := fmt.Sprintf("errmsg:%s", g.ErrMsg[0])
		ebitenutil.DebugPrint(screen, errMsg)
		return
	}
	ballProperty := fmt.Sprintf("ball width:%v,height:%v\n", g.Ball.Width(), g.Ball.Height())
	ballLocation := fmt.Sprintf("ball x:%v,y:%v\n", g.Ball.X(), g.Ball.Y())
	ballBeCollide := fmt.Sprintf("ballCollide:%v\n", g.Ball.BeCollided())

	ebitenutil.DebugPrint(screen, ballProperty+ballLocation+ballBeCollide)

	// 如果遊戲結束，顯示 Game Over
	if g.Mod == gamemanage.GameOver {
		ebitenutil.DebugPrint(screen, "Game Over\nPress Enter to restart")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Restart() {

	g.Ball.Reset(screenWidth, screenHeight)
	for i := 0; i < len(g.Jao); i++ {
		g.Jao[i].Reset(screenWidth, screenHeight)
	}
	g.Mod = gamemanage.Gaming

}

func (g *Game) SetDebugMode() {
	ebiten.SetTPS(10)
}

func main() {
	//遊戲紀錄
	if err := gamelog.InitLogger(); err != nil {
		log.Fatal(err)
	}
	defer gamelog.CloseLogger() // 确保在程序退出前关闭日志文件
	b := gameobject.NewBall(screenWidth, screenHeight)
	j := gameobject.NewJao(gameobject.Left, screenWidth, screenHeight)
	s := []string{}
	//處理背景圖片
	img, _, err := ebitenutil.NewImageFromFile("Images/background.png")
	if err != nil {
		log.Fatal(err)
	}
	//處理網路連線
	conn := socket.Connect()
	defer conn.Close()
	//初始化雙饒
	var jaoList []gameobject.Jao
	jaoList = append(jaoList, j)
	//初始化結構體
	game := &Game{
		Jao:        jaoList,
		Ball:       b,
		ErrMsg:     s,
		Network:    conn,
		Background: img,
	}
	//重啟遊戲達成初始化
	game.Restart()
	// game.SetDebugMode()
	//設置視窗大小及標題
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Jao Volleyball")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
