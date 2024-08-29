package orderbook

import (
	"fmt"
	"orderbook_tradingEngine/internal/orderbook"
)

func main() {
	ob := orderbook.NewOrderBook()

	order := orderbook.Order{
		ID:     "1",
		Type:   orderbook.Buy,
		Price:  100.0,
		Amount: 5,
	}

	ob.AddOrder(order)

	fmt.Println("Order added: ", order)

}