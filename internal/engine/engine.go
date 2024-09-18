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
	if order.Type == book.MarketSell{
		e.matchMarketOrder(order, book.Sell)
	}else if order.Type == book.MarketSell{
		e.matchMarketOrder(order, book.Buy)
	}else{
		e.OrderBook.AddOrder(order)
		matcher.MatchOrders(e.OrderBook)
	}
}

func (e *TradingEngine) matchMarketOrder(order book.Order, matchType book.OrderType) {
	// Ensure market buy matches with sell and market sell matches with buy
	matcher.MatchMarketOrder(e.OrderBook, order, matchType)
}

// GetOrderBook provides access to the current order book
func (e *TradingEngine) GetOrderBook() *book.Book {
	return e.OrderBook
}

func(e *TradingEngine) CancelOrder(orderID string) bool{
	return e.OrderBook.CancelOrderById(orderID)
}
