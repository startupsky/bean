package game

import (
	"geo"
)

type GameManager struct {
	parent      *Container
	onlineGames map[uint64]*Game
	maxId       uint64
}

func NewGameManager(parent *Container) *GameManager {
	this := new(GameManager)
	this.onlineGames = map[uint64]*Game{}
	this.parent = parent
	this.maxId = 1
	return this
}

func (this *GameManager) getNewId() uint64 {
	id := this.maxId
	this.maxId += 1
	return id
}

func (this *GameManager) CreateGame(host *Player, name string, maxPlayers int, city string, rect geo.Rectangle, gametype string) (game *Game) {
	g := new(Game)
	g.Id = this.getNewId()
	g.Name = name
	g.MaxPlayers = maxPlayers
	g.City = city
	g.Rect = rect
	g.State = gameWaiting

	g.Players = map[string]*Player{}
	g.Players[host.id] = host
	g.HostPlayer = host
	g.GameType = gametype

	this.onlineGames[g.Id] = g
	return g
}

func (this *GameManager) ListGame(city string) []*Game {
	cityGames := []*Game{}
	for _, v := range this.onlineGames {
		if city == "-1" || v.City == city {
			cityGames = append(cityGames, v)
		}
	}

	return cityGames
}

func (this *GameManager) JoinGame(player *Player, gameId uint64) error {
	for _, game := range this.onlineGames {
		if game.Id == gameId {
			if game.MaxPlayers <= len(game.Players) {
				return GamePlayersFullError
			}
			game.Players[player.id] = player
			log.Debug("Player=%v Join the Game=%v", player, game)
			return nil
		}
	}
	return GameNotFoundError
}