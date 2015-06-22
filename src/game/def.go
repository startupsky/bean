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
)

type BeanError struct {
	ErrNo  int
	ErrMsg string
}

func (this *BeanError) Error() string {
	return fmt.Sprint(this.ErrNo, this.ErrMsg)
}
