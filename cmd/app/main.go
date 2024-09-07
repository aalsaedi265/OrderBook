package main

import (
	"fmt"
	"orderbook_tradingEngine/internal/book"
	"orderbook_tradingEngine/internal/engine"
)


func main() {
	// Initialize the trading engine
	tradingEngine := engine.NewTradingEngine()

	// Example: Add a buy order through the engine
	order1 := book.NewLimitOrder("1", book.Buy, 100.0, 5)
	tradingEngine.ProcessOrder(order1)

	// Example: Add a sell order through the engine
	order2 := book.NewLimitOrder("2", book.Sell, 105.0, 3)
	tradingEngine.ProcessOrder(order2)

	// Add a higher buy order to see if matching happens
	order3 := book.NewLimitOrder("3", book.Buy, 120.0, 5)
	tradingEngine.ProcessOrder(order3)

	// Display the current order book
	fmt.Println("Buy Orders:", tradingEngine.GetOrderBook().BuyOrders())
	fmt.Println("Sell Orders:", tradingEngine.GetOrderBook().SellOrders())
}