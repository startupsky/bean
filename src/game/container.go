package game

import ()

type Container struct {
	parent    *Server
	playerMgr *PlayerManager
	gameMgr   *GameManager
}

func NewContainer(parent *Server) *Container {
	this := new(Container)
	this.parent = parent
	this.playerMgr = NewPlayerManager(this)
	this.gameMgr = NewGameManager(this)
	return this
}

func (this *Container) GetPlayerMgr() *PlayerManager {
	return this.playerMgr
}

func (this *Container) GetGameMgr() *GameManager {
	return this.gameMgr
}
