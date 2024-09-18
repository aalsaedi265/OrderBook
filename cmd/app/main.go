package main

import (
	"fmt"
	"orderbook_tradingEngine/internal/book"
	"orderbook_tradingEngine/internal/engine"
)


func main() {
	// Initialize the trading engine
	tradingEngine := engine.NewTradingEngine()

	order1 := book.NewLimitOrder("1", book.Buy, 100.0, 5)
	tradingEngine.ProcessOrder(order1)

	order2 := book.NewLimitOrder("2", book.Sell, 105.0, 3)
	tradingEngine.ProcessOrder(order2)


	// marketBuyOrder := book.Order{
	// 	ID:     "3",
	// 	Type:   book.MarketBuy,
	// 	Amount: 2,
	// }
	// tradingEngine.ProcessOrder(marketBuyOrder)

	// marketSellOrder := book.Order{
	// 	ID:     "4",
	// 	Type:   book.MarketSell,
	// 	Amount: 3,
	// }
	// tradingEngine.ProcessOrder(marketSellOrder)

	fmt.Println("Canceling Order 1 (Buy)")
	tradingEngine.CancelOrder("1")

	// Display the current order book
	fmt.Println("Buy Orders:", tradingEngine.GetOrderBook().BuyOrders())
	fmt.Println("Sell Orders:", tradingEngine.GetOrderBook().SellOrders())
}