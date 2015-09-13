package game

import (
	"bufio"
	"fmt"
	"net"
	"protocol"
	"strconv"
	"strings"
	"geo"
)

const (
	loginStat = iota
)

type Player struct {
	id          uint64
	passwd      string
	displayName string
	conn        net.Conn
	state       int
	eventQ      chan *protocol.Event
	parent      *PlayerManager
}

func NewPlayer(id uint64, passwd, displayName string, conn net.Conn, parent *PlayerManager) *Player {
	this := new(Player)
	this.id = id
	this.displayName = displayName
	this.passwd = passwd
	this.conn = conn
	this.parent = parent
	this.eventQ = make(chan *protocol.Event, 1024)
	return this
}

func (this *Player) DoWork(gameMgr *GameManager) {
	reader := bufio.NewReader(this.conn)
	proto := this.parent.parent.parent.proto
	for {
		if cmd, err := proto.ReadCommand(reader); err == nil {
			this.handleCommand(cmd, gameMgr)
		} else if err == protocol.WrongFmtError {
			resp := proto.CreateResponse()
			resp.ErrNo = ErrFormat
			if _, err := this.conn.Write(resp.Serialize()); err != nil {
				fmt.Println(6)
				break
			}
		} else {
			fmt.Println(5)
			break
		}

	}
}

func (this *Player) PostEvent(event *protocol.Event) {
	this.eventQ <- event
}

func (this *Player) handleCommand(cmd *protocol.Command, gameMgr *GameManager) {
	proto := gameMgr.parent.parent.proto
	resp := proto.CreateResponse()
	if cmd.CommandID == CREATEGAME && len(cmd.Arguments) == 5 {
		maxPlayer, _ := strconv.Atoi(cmd.Arguments[1])
		cityID, _ := strconv.Atoi(cmd.Arguments[2])
		topLeft := strings.Split(cmd.Arguments[3], ":")
		minX, _ := strconv.ParseInt(topLeft[0], 10, 64)
		minY, _ := strconv.ParseInt(topLeft[1], 10, 64)
		
		bottomRight := strings.Split(cmd.Arguments[4], ":")
		maxX, _ := strconv.ParseInt(bottomRight[0], 10, 64)
		maxY, _ := strconv.ParseInt(bottomRight[1], 10, 64)
		
		rect := &geo.Rectangle{MinX : minX, MinY : minY, MaxX : maxX, MaxY : maxY}
		game := gameMgr.CreateGame(this, cmd.Arguments[0], maxPlayer, cityID, *rect)
		
		if game == nil{
			resp.ErrNo = ErrFormat
			resp.Data = []string{"0"}
		} else {
			resp.ErrNo = ErrOK
			resp.Data = []string{"1"}
		}
		this.conn.Write(resp.Serialize())
	}
	if cmd.CommandID == LISTGAME && len(cmd.Arguments) == 1 {
		cityID,_ := strconv.Atoi(cmd.Arguments[0])
		games := gameMgr.ListGame(cityID)
		
		resp.ErrNo = ErrOK
		data := []string{}
		
		for i:=0;i<len(games);i++{
			gamestr := fmt.Sprintf("%d %s %d:%d %d:%d %d %d", games[i].Id, games[i].Name, games[i].Rect.MinX, games[i].Rect.MinY, games[i].Rect.MaxX, games[i].Rect.MaxY, len(games[i].Players)+1, games[i].MaxPlayers)
			data = append(data, gamestr)
		}
		
		resp.Data = data
		this.conn.Write(resp.Serialize())
	}
	if cmd.CommandID == JOINGAME && len(cmd.Arguments) == 1{
		gameId,_ := strconv.ParseUint(cmd.Arguments[0], 10, 64)
	 	err := gameMgr.JoinGame(this, gameId)
		if err == nil{
			resp.ErrNo = ErrOK
			resp.Data = []string{"0"}
		} else {
			resp.ErrNo = ErrFormat
			resp.Data = []string{err.Error()}
		}
		this.conn.Write(resp.Serialize())
	}
}
