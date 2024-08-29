package main

import (
	"fmt"
	"orderbook_tradingEngine/internal/book"
)


func main() {
	// Initialize the order book
	orderBook := book.NewBook()

	// Example: Add a few orders
	order1 := book.Order{
		ID:     "1",
		Type:   book.Buy,
		Price:  100.0,
		Amount: 5,
	}
	orderBook.AddOrder(order1)

	order2 := book.Order{
		ID:     "2",
		Type:   book.Sell,
		Price:  105.0,
		Amount: 3,
	}
	orderBook.AddOrder(order2)

	// List all buy orders
	fmt.Println("Buy Orders:", orderBook.ListOrders(book.Buy))

	// List all sell orders
	fmt.Println("Sell Orders:", orderBook.ListOrders(book.Sell))
}