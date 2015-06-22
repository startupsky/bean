package game

import (
	"geo"
)

type GameManager struct {
	parent *Container
	games  map[int64]*Game
	maxId  int64
}

func NewGameManager(parent *Container) *GameManager {
	this := new(GameManager)
	this.games = map[int64]*Game{}
	this.parent = parent
	this.maxId = 1
	return this
}

func (this *GameManager) getNewId() int64 {
	id := this.maxId
	this.maxId += 1
	return id
}

func (this *GameManager) CreateGame(maxPlayers int, cityId int, rect geo.Rectangle) (game *Game, err error) {
	g := new(Game)
	g.Id = this.getNewId()
	g.MaxPlayers = maxPlayers
	g.City = cityId
	g.Rect = rect
	g.State = gameWaiting
	this.games[g.Id] = g
	return g, nil
}

func (this *GameManager) JoinGame(player *Player, game *Game) error {
	return nil
}
