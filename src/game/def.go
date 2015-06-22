package game

import (
	"fmt"
)

const (
	ErrOK         = int16(0)
	ErrInternal   = int16(1)
	ErrUnExpected = int16(2)
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
