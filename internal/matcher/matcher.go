package matcher

import (
	"fmt"
	"sort"
	"orderbook_tradingEngine/internal/book"
)

// MatchOrders attempts to match buy and sell orders in the order book.
func MatchOrders(b *book.Book) {
	// Lock the order book to ensure thread safety
	b.Mu.Lock()
	defer b.Mu.Unlock()

	// Sort the buy orders by price (highest first)
	sort.SliceStable(b.BuyOrders(), func(i, j int) bool {
		return b.BuyOrders()[i].Price > b.BuyOrders()[j].Price
	})

	// Sort the sell orders by price (lowest first)
	sort.SliceStable(b.SellOrders(), func(i, j int) bool {
		return b.SellOrders()[i].Price < b.SellOrders()[j].Price
	})

	fmt.Println("DEBUG: Start matching orders")
	fmt.Println("DEBUG: Current Buy Orders:", b.BuyOrders())
	fmt.Println("DEBUG: Current Sell Orders:", b.SellOrders())

	// Process matching orders as long as we have both buy and sell orders
	for len(b.BuyOrders()) > 0 && len(b.SellOrders()) > 0 {
		buyOrder := b.BuyOrders()[0]   // Highest buy price
		sellOrder := b.SellOrders()[0] // Lowest sell price

		fmt.Printf("DEBUG: Trying to match Buy Order: %v with Sell Order: %v\n", buyOrder, sellOrder)

		// Check if the buy price is greater than or equal to the sell price
		if buyOrder.Price >= sellOrder.Price {
			// Determine the amount that can be traded
			tradeAmount := min(buyOrder.Amount, sellOrder.Amount)

			// Execute the trade by reducing the amounts
			buyOrder.Amount -= tradeAmount
			sellOrder.Amount -= tradeAmount

			// Debugging info for successful trades
			fmt.Printf("Trade executed: %d units @ $%.2f\n", tradeAmount, buyOrder.Price)

			// Remove fully filled orders or update partially filled ones
			if buyOrder.Amount == 0 {
				b.RemoveBuyExecutedOrders(0) // Remove fully filled buy order
				fmt.Println("Buy order fully filled and removed.")
			} else {
				// Update the order book with partially filled buy order
				b.BuyOrders()[0] = buyOrder
				fmt.Printf("Buy order partially filled, remaining amount: %d\n", buyOrder.Amount)
			}

			if sellOrder.Amount == 0 {
				b.RemoveSellOrders(0) // Remove fully filled sell order
				fmt.Println("Sell order fully filled and removed.")
			} else {
				// Update the order book with partially filled sell order
				b.SellOrders()[0] = sellOrder
				fmt.Printf("Sell order partially filled, remaining amount: %d\n", sellOrder.Amount)
			}

		} else {
			// No match is possible if buy price is lower than sell price, break out of the loop
			fmt.Println("No more matches possible.")
			break
		}
	}
	fmt.Println("DEBUG: End of matching orders")
}

func MatchMarketOrder(b *book.Book, order book.Order, matchType  book.OrderType){
	if matchType  == book.Sell{
		for len(b.SellOrders()) > 0 && order.Amount > 0{
			sellOrder := b.SellOrders()[0]
			if order.Amount > 0{
				tradeAmount := min(order.Amount, sellOrder.Amount)
				order.Amount -= tradeAmount
				sellOrder.Amount -= tradeAmount
				fmt.Printf("Market Buy executed: %d units @ $%.2f\n", tradeAmount, sellOrder.Price)
				if sellOrder.Amount == 0{
					b.RemoveSellOrders(0)
				}
			}
		}
	} else if matchType  == book.Buy{
		for len(b.BuyOrders()) > 0 && order.Amount > 0 {
			buyOrder := b.BuyOrders()[0]
			if order.Amount > 0 {
				tradeAmount := min(order.Amount, buyOrder.Amount)
				order.Amount -= tradeAmount
				buyOrder.Amount -= tradeAmount
				fmt.Printf("Market Sell executed: %d units @ $%.2f\n", tradeAmount, buyOrder.Price)
				if buyOrder.Amount == 0 {
					b.RemoveBuyExecutedOrders(0) // Remove fully filled buy order
				}
			}
		}
	}
}

// Utility function to find the minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
