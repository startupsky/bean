package game

import (
	"bufio"
	"net"
)

type PlayerManager struct {
	parent  *Container
	players []*Player
}

func NewPlayerManager(parent *Container) *PlayerManager {
	this := new(PlayerManager)
	this.players = []*Player{}
	this.parent = parent

	return this
}

func (this *PlayerManager) Login(conn net.Conn) (player *Player, err error) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	if cmd, err := this.parent.parent.proto.ReadCommand(rw); err == nil {
		if cmd.CommandID == CMDLOGIN && len(cmd.Arguments) == 2 {
			name := cmd.Arguments[0]
			passwd := cmd.Arguments[1]
		}
	}
	return nil, nil
}

func (this *PlayerManager) Logout(player *Player) {

}
