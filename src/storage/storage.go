package storage

import (
	"model"
)

var Ele *Elements

type Elements struct {
	Players []*model.Player
	Beans   []*model.Bean
}

func init() {
	Ele = new(Elements)
	Ele.Players = make([]*model.Player, 0)
	Ele.Beans = make([]*model.Bean, 0)
}

func (this *Elements) PlayerReport(id uint64, longitude int64, latitude int64) {
	for _, p := range this.Players {
		if p.Id == id {
			p.Longitude = longitude
			p.Latitude = latitude
			return
		}
	}
	player := new(model.Player)
	player.Id = id
	player.Longitude = longitude
	player.Latitude = latitude

	this.Players = append(this.Players, player)
}

func (this *Elements) BeanManipulate(id uint64, state uint8, longitude int64, latitude int64) {
	for _, b := range this.Beans {
		if b.Id == id {
			b.State = state
			b.Longitude = longitude
			b.Latitude = latitude
			return
		}
	}
	bean := new(model.Bean)
	bean.State = state
	bean.Id = id
	bean.Longitude = longitude
	bean.Latitude = latitude

	this.Beans = append(this.Beans, bean)
}

func (this *Elements) CleanPlayers() {
	this.Players = make([]*model.Player, 0)
}

func (this *Elements) CleanBeans() {
	Ele.Beans = make([]*model.Bean, 0)
}
