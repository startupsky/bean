package game

import (
	"geo"
)

type Game struct {
	Id         uint64
	Name       string
	City       string
	Rect       geo.Rectangle
	MaxPlayers int
	Players    map[string]*Player
	HostPlayer *Player
	State      int
	GameType   string
	Beans	[]*geo.Point
}
