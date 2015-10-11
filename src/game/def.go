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
)

const (
	gameWaiting = iota
	gameStarted
)

const (
	CMDLOGIN = "login"
	CREATEGAME = "creategame"
	LISTGAME = "listgame"
	JOINGAME = "joingame"
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
