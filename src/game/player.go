package game

import (
	"net"
)

const (
	loginStat = iota
)

type Player struct {
	id          int64
	passwd      string
	displayName string
	conn        net.Conn
	state       int
}

func (this *Player) DoWork() {
	for {

	}
}
