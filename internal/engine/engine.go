package engine

import (
	"orderbook_tradingEngine/internal/book"
	"orderbook_tradingEngine/internal/matcher"
)

type TradingEngine struct{
	OrderBook *book.Book
}

//new instance of TradingEngin & initializes  OrderBook field with a new Book instance, returns pointer newly created TradingEngin instance
func NewTradingEngine() *TradingEngine{
	return &TradingEngine{
		OrderBook: book.NewBook(),
	}
}

// ProcessOrder adds an order to the book and triggers matching
func (e *TradingEngine) ProcessOrder(order book.Order) {
	e.OrderBook.AddOrder(order)
	matcher.MatchOrders(e.OrderBook)
}

// GetOrderBook provides access to the current order book
func (e *TradingEngine) GetOrderBook() *book.Book {
	return e.OrderBook
}
