package orderbook

import "time"

type OrderType int

const (
	Buy OrderType = iota
	sell
)

type Order struct {
	ID        string
	Type      OrderType
	Price     float64
	Amount    int
	Timestamp time.Time
}