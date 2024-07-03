package domain

type Furniture struct {
	Name   string  `json:"name"`
	Maker  string  `json:"maker"`
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	Length float64 `json:"length"`
}
