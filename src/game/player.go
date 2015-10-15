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
	X			float64
	Y			float64
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
		log.Debug("---creategame----\n")
		resp.ReplyNo = CreategameReply
		if len(cmd.Arguments) == 6{
			maxPlayer, _ := strconv.Atoi(cmd.Arguments[1])
			city := cmd.Arguments[2]
			topLeft := strings.Split(cmd.Arguments[3], ":")
			minX, _ := strconv.ParseFloat(topLeft[0], 64)
			minY, _ := strconv.ParseFloat(topLeft[1], 64)
			
			bottomRight := strings.Split(cmd.Arguments[4], ":")
			maxX, _ := strconv.ParseFloat(bottomRight[0], 64)
			maxY, _ := strconv.ParseFloat(bottomRight[1], 64)
			
			rect := &geo.Rectangle{MinX : minX, MinY : minY, MaxX : maxX, MaxY : maxY}
			
			gametype := cmd.Arguments[5]
			game := gameMgr.CreateGame(this, cmd.Arguments[0], maxPlayer, city, *rect, gametype)
			
			if game == nil{
				log.Debug("create game fail\n")
				resp.Data = []string{"0"}
			} else {
				resp.Data = []string{fmt.Sprintf("%d",game.Id)}
			}
		}else{
			log.Debug("argument wrong\n")
			resp.Data = []string{"1"}
		}
	
	case LISTGAME:
		log.Debug("---listgame----\n")
		resp.ReplyNo = ListgameReply
	
		if len(cmd.Arguments) == 1{
			log.Debug(cmd.Arguments[0])
			log.Debug("\n")
			
			city := cmd.Arguments[0]
			games := gameMgr.ListGame(city)
	
			data := []string{}
			
			for i:=0;i<len(games);i++{
				gamestr := fmt.Sprintf("%d %s %s %f:%f %f:%f %d %d %s", games[i].Id, games[i].City, games[i].Name, games[i].Rect.MinX, games[i].Rect.MinY, games[i].Rect.MaxX, games[i].Rect.MaxY, len(games[i].Players), games[i].MaxPlayers, games[i].GameType)
				data = append(data, gamestr)
			}
			resp.Data = data
		}else{
			log.Debug("argument wrong\n")
			resp.Data = []string{}
		}

	case JOINGAME:
		log.Debug("---joingame----\n")
		resp.ReplyNo = JoingameReply
		if len(cmd.Arguments) == 1{
			gameId,_ := strconv.ParseUint(cmd.Arguments[0], 10, 64)
		 	err := gameMgr.JoinGame(this, gameId)
			
			if err == nil{
				resp.Data = []string{"1"}
			} else {
				resp.Data = []string{err.Error()}
			}
		}else{
			log.Debug("argument wrong\n")
			resp.Data = []string{"0"}
		}
		
	case SHOWPLAYERS:
		log.Debug("---showplayer----\n")
		resp.ReplyNo = ShowplayersReply
		if len(cmd.Arguments) == 1{
			gameId,_ := strconv.ParseUint(cmd.Arguments[0], 10, 64)
		 	game := gameMgr.onlineGames[gameId]
			if game != nil{
				data := []string{}
				for _,player:=range game.Players{
					playerStr := fmt.Sprintf("%s %s %f:%f", player.id, player.displayName, player.X, player.Y)
					data = append(data, playerStr)
				}
				resp.Data = data
			}
		}else{
			log.Debug("argument wrong\n")
			resp.Data = []string{"0"}
		}
		
	case LEAVEGAME:
		log.Debug("---leavegame----\n")
		resp.ReplyNo = LeavegameReply
		if len(cmd.Arguments) == 1{
			gameId,_ := strconv.ParseUint(cmd.Arguments[0], 10, 64)
			delete(gameMgr.onlineGames[gameId].Players, this.id)
			if len(gameMgr.onlineGames[gameId].Players) == 0{
				delete(gameMgr.onlineGames, gameId)
			}else if gameMgr.onlineGames[gameId].HostPlayer == this{
				for _,player:=range gameMgr.onlineGames[gameId].Players{
					gameMgr.onlineGames[gameId].HostPlayer = player
					break
				}
			}
		}else{
			log.Debug("argument wrong\n")
			resp.Data = []string{"0"}
		}
	case LOGOUT:
		return true
	default:
		log.Debug("---unkowncommand----\n")
		resp.ReplyNo = ErrorReply
		resp.Data = []string{"UnknownCMD"}

	}
	
	this.conn.Write(resp.Serialize())
	return false
}
