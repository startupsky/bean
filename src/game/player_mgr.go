package game

import (
	"bufio"
	"net"
	"strconv"
	"util"
)

type PlayerManager struct {
	parent        *Container
	onlinePlayers util.BeanSlice
}

func NewPlayerManager(parent *Container) *PlayerManager {
	this := new(PlayerManager)
	this.onlinePlayers = util.BeanSlice{}
	this.parent = parent

	return this
}

func (this *PlayerManager) Login(conn net.Conn) (player *Player, err error) {
	proto := this.parent.parent.proto
	if cmd, err := proto.ReadCommand(bufio.NewReader(conn)); err == nil {
		if cmd.CommandID == CMDLOGIN && len(cmd.Arguments) == 2 {
			id, _ := strconv.ParseUint(cmd.Arguments[0], 10, 64)
			passwd := cmd.Arguments[1]
			//valid logic here

			resp := proto.CreateResponse()
			resp.ErrNo = ErrOK
			resp.Data = []string{"1"}
			if _, err := conn.Write(resp.Serialize()); err == nil {
				player = NewPlayer(id, passwd, "dislay name", conn, this)
				this.onlinePlayers = append(this.onlinePlayers, player)
				return player, nil
			}
		}
	}
	return nil, nil
}

func (this *PlayerManager) Logout(player *Player) {
	this.onlinePlayers = this.onlinePlayers.Remove(player)
}
