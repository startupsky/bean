package game

import (
	"fmt"
	"github.com/nporsche/np-golang-logging"
	"net"
	"protocol"
)

var log = logging.MustGetLogger("game")

type Server struct {
	proto *protocol.Protocol
	port  int
}

func NewServer(port int, proto *protocol.Protocol) *Server {
	this := new(Server)
	this.proto = proto
	this.port = port

	return this
}

func (this *Server) Start() {
	container := NewContainer(this)
	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", this.port))
	if err != nil {
		log.Critical("Listen Error=[%s]", err.Error())
		return
	}
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go this.handleConnection(conn, container)
	}
}

func (this *Server) handleConnection(conn net.Conn, container *Container) {
	defer func() {
		if r := recover(); r != nil {
			log.Critical("UnExpected Fatal %v", r)
		}
		conn.Close()
		log.Debug("connection leaves")
	}()
	log.Debug("A new connection comes")
	playerMgr := container.GetPlayerMgr()
	if player, err := playerMgr.Login(conn); err != nil {
		log.Warning("Login failed from=[%v], err=[%s]", conn.RemoteAddr().String(), err.Error())
	} else {
		player.DoWork()
		playerMgr.Logout(player)
	}
}
