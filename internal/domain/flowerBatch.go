package domain

import "time"

type FlowerBatch struct {
	FlowerName  string    `json:"flower_name"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	ArrivalDate time.Time `json:"arrival_date"`
}
