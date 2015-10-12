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
	id          string
	passwd      string
	displayName string
	conn        net.Conn
	state       int
	eventQ      chan *protocol.Event
	parent      *PlayerManager
}

func NewPlayer(id string, passwd, displayName string, conn net.Conn, parent *PlayerManager) *Player {
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
			logout := this.handleCommand(cmd, gameMgr)
			if logout{
				return
			}
		} else if err == protocol.WrongFmtError {
			resp := proto.CreateResponse()
			resp.ReplyNo = ErrorReply
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

func (this *Player) handleCommand(cmd *protocol.Command, gameMgr *GameManager) (logout bool) {
	proto := gameMgr.parent.parent.proto
	resp := proto.CreateResponse()
	
	switch cmd.CommandID{
	case CREATEGAME:
		resp.ReplyNo = CreategameReply
		if len(cmd.Arguments) == 5{
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
				resp.Data = []string{"0"}
			} else {
				resp.Data = []string{"1"}
			}
		}else{
			resp.Data = []string{"1"}
		}
		this.conn.Write(resp.Serialize())
	
	case LISTGAME:
		resp.ReplyNo = ListgameReply
		if len(cmd.Arguments) == 1{
			cityID,_ := strconv.Atoi(cmd.Arguments[0])
			games := gameMgr.ListGame(cityID)
	
			data := []string{}
			
			for i:=0;i<len(games);i++{
				gamestr := fmt.Sprintf("%d %d %s %d:%d %d:%d %d %d", games[i].Id, games[i].City, games[i].Name, games[i].Rect.MinX, games[i].Rect.MinY, games[i].Rect.MaxX, games[i].Rect.MaxY, len(games[i].Players), games[i].MaxPlayers)
				data = append(data, gamestr)
			}
			resp.Data = data
		}else{
			resp.Data = []string{}
		}
		this.conn.Write(resp.Serialize())
	
	case JOINGAME:
		resp.ReplyNo = JoingameReply
		if len(cmd.Arguments) == 1{
			gameId,_ := strconv.ParseUint(cmd.Arguments[0], 10, 64)
		 	err := gameMgr.JoinGame(this, gameId)
			
			if err == nil{
				resp.Data = []string{"0"}
			} else {
				resp.Data = []string{err.Error()}
			}
		}else{
			resp.Data = []string{"1"}
		}

		this.conn.Write(resp.Serialize())
		
	case LOGOUT:
		return true
	default:
		resp.ReplyNo = ErrorReply
		resp.Data = []string{"UnknownCMD"}
		this.conn.Write(resp.Serialize())
	}

	return false
}
