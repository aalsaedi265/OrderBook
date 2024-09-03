package book

import (
	"sync"
	"time"
)

type OrderType int

const (
	Buy OrderType = iota
	Sell
)

type Order struct {
	ID        string
	Type      OrderType
	Price     float64
	Amount    int
	Timestamp time.Time
}

type Book struct {
    buyOrders  []Order
    sellOrders []Order
    Mu         sync.Mutex
}



func (b *Book) BuyOrders() []Order {
	return b.buyOrders
}

// SellOrders returns a slice of all sell orders in the order book.
func(b *Book)SellOrders() []Order{
	return b.sellOrders
}