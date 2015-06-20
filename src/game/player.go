package game

import (
	"net"
)

type Player struct {
	id     int64
	name   string
	passwd string
	conn   net.Conn
}

func (this *Player) DoWork() {

}
