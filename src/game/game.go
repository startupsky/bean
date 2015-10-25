package game

import (
	"geo"
	"fmt"
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
	Row			int
	Column		int
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
	
	distanceX := 2/11000.0 // 2m
	distanceY := distanceX
	
	beans := []*geo.Point{}
	
	this.Row = (int) ((stopY-startY)/distanceY)
	this.Column = (int) ((stopX - startX)/distanceX)
	
	for i := 0; i < this.Row; i++{
		for j:=0;j<this.Column;j++{
			pointX := startX + 0.5*distanceX + float64(i)*distanceX
			pointY := startY + 0.5*distanceY + float64(j)*distanceY
			point := &geo.Point{RowIndex: i, ColumnIndex: j, X : pointX, Y : pointY, Role:1}
			beans = append(beans, point)	
		}
	}
	
	for _,player:=range this.Players{
		beansIndex :=0
		if player.X > startX && player.X < stopX && player.Y > startY && player.Y < stopY{
			playerRow := int( (player.X - startX)/distanceX)
			playerColumn := int( (player.Y - startY)/distanceY)
			beansIndex = playerRow*this.Column + playerColumn
		}
		for beans[beansIndex].Role == -1{
			beansIndex = (beansIndex+1)%(this.Row*this.Column)
		}
		player.X = beans[beansIndex].X
		player.Y = beans[beansIndex].Y
		beans[beansIndex].Role = -1
	}
	this.Beans = beans
}

func (this *Game) UpdateMap(X float64, Y float64) (score int) {
	distanceX := 0.5/11000.0 // 0.5m
	distanceY := distanceX	
	score = 0
	for _,bean:=range this.Beans{
		if bean.Role == 1{
			startX := bean.X - distanceX
			stopX := bean.X + distanceX
			startY := bean.Y - distanceY
			stopY := bean.Y + distanceY
			
			if X > startX && X < stopX && Y > startY && Y < stopY{
				log.Debug(fmt.Sprintf("UpdateMap: eat [%f]:[%f]", bean.X, bean.Y))
				bean.Role = -1
				score = score+1
			}			
		}
	}
	return score
}