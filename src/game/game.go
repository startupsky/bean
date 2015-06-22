package game

import (
	"geo"
)

type Game struct {
	Id         uint64
	Name       string
	City       int
	Rect       geo.Rectangle
	MaxPlayers int
	Players    map[uint64]*Player
	HostPlayer *Player
	State      int
}
