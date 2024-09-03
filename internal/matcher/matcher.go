package matcher

import (
	"orderbook_tradingEngine/internal/book"
	"sort"
	"fmt"
)

func MatchOrders(b *book.Book) {
	// Lock the order book while matching to avoid race conditions
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

	// Process matching orders
	for len(b.BuyOrders()) > 0 && len(b.SellOrders()) > 0 {
		buyOrder := b.BuyOrders()[0]   // Highest buy price
		sellOrder := b.SellOrders()[0] // Lowest sell price

		// Check if a match is possible
		if buyOrder.Price >= sellOrder.Price {
			// Determine the quantity that can be traded
			tradeAmount := min(buyOrder.Amount, sellOrder.Amount)

			// Execute the trade (reduce amounts)
			buyOrder.Amount -= tradeAmount
			sellOrder.Amount -= tradeAmount

			// Debugging information
			fmt.Printf("Trade executed: Buy %d @ %f, Sell %d @ %f\n", tradeAmount, buyOrder.Price, tradeAmount, sellOrder.Price)

			// Remove fully executed orders using the existing methods
			if buyOrder.Amount == 0 {
				b.RemoveBuyExecutedOrders(0) // Remove the first buy order
				fmt.Printf("Buy order fully executed and removed, %d \n", buyOrder.Amount)
			} else {
				b.BuyOrders()[0] = buyOrder // Update the slice with the partially filled order
			}
			if sellOrder.Amount == 0 {
				b.RemoveSellOrders(0) // Remove the first sell order
				fmt.Printf("Sell order fully executed and removed, %d \n", sellOrder.Amount)
			} else {
				b.SellOrders()[0] = sellOrder
			}
		} else {
			// No match possible, exit the loop
			break
		}
	}
}

// Utility function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
