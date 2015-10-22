package game

import (
	"geo"
)

type Game struct {
	Id         uint64
	Name       string
	City       string
	Rect       geo.Rectangle
	MaxPlayers int
	Players    map[string]*Player
	HostPlayer *Player
	State      int
	GameType   string
	Beans	[]*geo.Point
}

func (this *Game) getXRange() (float64,float64){
	if this.Rect.MinX > this.Rect.MaxX{
		return this.Rect.MaxX, this.Rect.MinX
	}
	return this.Rect.MinX, this.Rect.MaxX
}

func (this *Game) getYRange() (float64,float64){
	if this.Rect.MinY > this.Rect.MaxY{
		return this.Rect.MaxY, this.Rect.MinY
	}
	return this.Rect.MinY, this.Rect.MaxY
}

func (this *Game) SetupMap() {
	
	startX, stopX := this.getXRange()
	startY, stopY := this.getYRange()
	
	distanceX := 1/11000.0 // 1m
	distanceY := distanceX
	
	beans := []*geo.Point{}
	
	if distanceX > 0 && distanceY > 0{
		for i := startX; i < stopX; i += distanceX {
			for j:= startY; j < stopY; j += distanceY {
				point := &geo.Point{X : i, Y : j}
				beans = append(beans, point)			
			}
		}
	}
	this.Beans = beans
}