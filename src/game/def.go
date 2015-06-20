package game

import (
	"fmt"
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

var ErrOK = &BeanError{0, ""}
var ErrLogin = &BeanError{1, "Login Error"}
