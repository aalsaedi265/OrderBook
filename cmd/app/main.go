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

	// Add market buy order (will execute immediately at the best available sell price)
	marketBuyOrder := book.Order{
		ID:     "3",
		Type:   book.MarketBuy,
		Amount: 2,
	}
	tradingEngine.ProcessOrder(marketBuyOrder)

	// Add market sell order (will execute immediately at the best available buy price)
	marketSellOrder := book.Order{
		ID:     "4",
		Type:   book.MarketSell,
		Amount: 3,
	}
	tradingEngine.ProcessOrder(marketSellOrder)


	// Display the current order book
	fmt.Println("Buy Orders:", tradingEngine.GetOrderBook().BuyOrders())
	fmt.Println("Sell Orders:", tradingEngine.GetOrderBook().SellOrders())
}