package adapter

import "testing"

func TestMusic(t *testing.T) {
	var player Player = MusicPLayer{Src: "晴天"}
	player.PlayMusic()
}

func TestMusic1(t *testing.T) {
	var gamePlayerAdapter Player = GamePlayerAdapter{GamePlayer{Src: "未闻花名"}}
	gamePlayerAdapter.PlayMusic()
}
