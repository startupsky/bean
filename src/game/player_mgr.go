package game

import (
	"bufio"
	"net"
	"strconv"
	"util"
	"user"
)

type PlayerManager struct {
	parent        *Container
	onlinePlayers util.BeanSlice
	userManager *user.UserManager
}

func NewPlayerManager(parent *Container) *PlayerManager {
	this := new(PlayerManager)
	this.onlinePlayers = util.BeanSlice{}
	this.parent = parent
	this.userManager = user.NewUserManager()
	return this
}

func (this *PlayerManager) Login(conn net.Conn) (player *Player, err error) {
	proto := this.parent.parent.proto
	if cmd, err := proto.ReadCommand(bufio.NewReader(conn)); err == nil {
		if cmd.CommandID == CMDLOGIN && len(cmd.Arguments) == 2 {
			userid := cmd.Arguments[0]
			id, _ := strconv.ParseUint(cmd.Arguments[0], 10, 64)
			passwd := cmd.Arguments[1]

			resp := proto.CreateResponse()
			
			user := this.userManager.AddUser(userid, passwd)
			if user == nil {
				user = this.userManager.GetUser(userid, passwd)
			}else{
				this.userManager.Save()
			}
			
			if user == nil{
				resp.ErrNo = ErrFormat
				resp.Data = []string{"0"}
			}else{
				resp.ErrNo = ErrOK
				resp.Data = []string{"1"}
				if _, err := conn.Write(resp.Serialize()); err == nil {
					player = NewPlayer(id, passwd, "dislay name", conn, this)
					this.onlinePlayers = append(this.onlinePlayers, player)
					return player, nil
				}
			}
		}
	}
	return nil, nil
}

func (this *PlayerManager) Logout(player *Player) {
	this.onlinePlayers = this.onlinePlayers.Remove(player)
}
