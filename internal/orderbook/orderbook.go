package orderbook

import (
	"sync"
	"time"
)

type Orderbook struct{
	buyOrders  []Order
	sellOrders []Order
	mu         sync.Mutex
}

func NewOrderBook() *Orderbook {
	return &Orderbook{
		buyOrders:  []Order{},
		sellOrders: []Order{},
	}
}

func (ob *Orderbook) AddOrder(order Order) {
	ob.mu.Lock()
	defer ob.mu.Unlock() //executed at the end of the AddOrder function

	order.Timestamp = time.Now()
	
}