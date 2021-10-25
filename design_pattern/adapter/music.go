package adapter

import "fmt"

type Player interface {
	PlayMusic()
}

type MusicPLayer struct {
	Src string
}

func (m MusicPLayer) PlayMusic() {
	fmt.Println("音乐播放器正在播放：", m.Src)
}

type GamePlayer struct {
	Src string
}

func (g GamePlayer) PlayGameMusic() {
	fmt.Println("游戏播放器正在播放：", g.Src)
}

//适配器
type GamePlayerAdapter struct {
	GamePlayer
}

//转接了一下
func (g GamePlayerAdapter) PlayMusic() {
	g.GamePlayer.PlayGameMusic()
}
