
package book

import (
	"time"
)



func NewBook() *Book {
	return &Book{
		buyOrders:  []Order{},
		sellOrders: []Order{},
	}
}

func (b *Book) AddOrder(order Order) {
	b.Mu.Lock()
	defer b.Mu.Unlock()

	order.Timestamp = time.Now()

	if order.Type == Buy {
		b.buyOrders = append(b.buyOrders, order)
	} else {
		b.sellOrders = append(b.sellOrders, order)
	}
}


func (b *Book) ListOrders(orderType OrderType) []Order {
	b.Mu.Lock()
	defer b.Mu.Unlock()

	if orderType == Buy {
		return b.buyOrders
	} else if orderType == Sell {
		return b.sellOrders
	} else {
		return append(b.buyOrders, b.sellOrders...)
	}
}

// removes orders that have been already exectuted fully. Optimizes system performance by reducing the size of the order book, which speeds up the matching process and avoids unnecessary data processing.

//this works because go slice makes refrences to the underlying array.
func (b *Book) RemoveBuyExecutedOrders(index int) {
	//remove by skipping elements
	b.buyOrders = append(b.buyOrders[:index], b.buyOrders[index+1:]...)
}
func (b *Book) RemoveSellOrders(index int) {
     // Replace the order at `index` with the last order in the list
    b.sellOrders[index] = b.sellOrders[len(b.sellOrders)-1]

    // Shorten the slice to remove the last element
    b.sellOrders = b.sellOrders[:len(b.sellOrders)-1]
}