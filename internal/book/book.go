
package book

import (
	"sync"
	"time"
)

type Book struct {
	buyOrders  []Order
	sellOrders []Order
	mu         sync.Mutex
}

func NewBook() *Book {
	return &Book{
		buyOrders:  []Order{},
		sellOrders: []Order{},
	}
}

func (b *Book) AddOrder(order Order) {
	b.mu.Lock()
	defer b.mu.Unlock()

	order.Timestamp = time.Now()

	if order.Type == Buy {
		b.buyOrders = append(b.buyOrders, order)
	} else {
		b.sellOrders = append(b.sellOrders, order)
	}
}

func (b *Book) ListOrders(orderType OrderType) []Order {
	b.mu.Lock()
	defer b.mu.Unlock()

	if orderType == Buy {
		return b.buyOrders
	} else if orderType == Sell {
		return b.sellOrders
	} else {
		return append(b.buyOrders, b.sellOrders...)
	}
}
