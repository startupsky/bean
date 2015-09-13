package game

import (
	"fmt"
)

const (
	ErrOK     = int16(0)
	ErrFormat = int16(1)
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
