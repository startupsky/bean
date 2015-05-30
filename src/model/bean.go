package model

const (
	BeanStatNew   = 0
	BeanStatEaten = 1
)

type Bean struct {
	Id        uint64 `json:"id"`
	State     uint8  `json:"state"`
	Longitude int64  `json:"longitude"`
	Latitude  int64  `json:"latitude"`
}
