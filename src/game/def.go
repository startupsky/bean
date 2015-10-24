package game

import (
	"fmt"
)

const (
	ErrorReply = int16(-1)
	LoginReply = int16(0)
	CreategameReply = int16(1)
	ListgameReply = int16(2)
	JoingameReply = int16(3)
	ShowplayersReply = int16(4)
	LeavegameReply = int16(5)
	StartgameReply = int16(6)
	QuerygameReply = int16(7)
	QuerymapReply = int16(8)
	ReportReply = int16(9)
	StopgameReply = int16(10)
)

const (
	gameWaiting = iota
	gameStarted
	gameStopped
)

const (
	CMDLOGIN = "login"
	CREATEGAME = "creategame"
	LISTGAME = "listgame"
	JOINGAME = "joingame"
	LOGOUT = "logout"
	SHOWPLAYERS = "showplayers"
	LEAVEGAME = "leavegame"
	STARTGAME = "startgame"
	QUERYGAME = "querygame"
	QUERYMAP = "querymap"
	REPORT = "report"
	STOPGAME = "stopgame"
)

type BeanError struct {
	ErrNo  int
	ErrMsg string
}

func (this *BeanError) Error() string {
	return fmt.Sprint(this.ErrNo, this.ErrMsg)
}

var GamePlayersFullError = &BeanError{1, "Game Players Full"}
var GameNotFoundError = &BeanError{2, "Game Not Found"}
