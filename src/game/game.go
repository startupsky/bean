package game

import (
	"geo"
)

type Game struct {
	Id         int64
	Name       string
	City       int
	Rect       geo.Rectangle
	MaxPlayers int
	Players    []*Player
	State      int
}
