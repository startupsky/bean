package model

type Player struct {
	Id        uint64 `json:"id"`
	Longitude int64  `json:"longitude"`
	Latitude  int64  `json:"latitude"`
}
