package book

import "testing"

func TestAddOrder(t *testing.T){
	orderBook := NewBook()
	order := Order{
		ID:     "1",
		Type:   Buy,
		Price:  100.0,
		Amount: 5,
	}	
	initialLength := len(orderBook.ListOrders(Buy))
	orderBook.AddOrder(order)
	newLength := len(orderBook.ListOrders(Buy))

	// Check if the length has increased by 1
	if newLength != initialLength+1 {
		t.Fatalf("expected buy orders length to increase by 1, but it didn't (initial: %d, new: %d)", initialLength, newLength)
	}

	buyOrders := orderBook.ListOrders(Buy)
	lastIndex := len(buyOrders) - 1

	addedOrder := buyOrders[lastIndex]
	if addedOrder.ID != order.ID {
		t.Errorf("expected order ID '%s', got '%s'", order.ID, addedOrder.ID)
	}
	if addedOrder.Price != order.Price {
		t.Errorf("expected order price %f, got %f", order.Price, addedOrder.Price)
	}
	if addedOrder.Amount != order.Amount {
		t.Errorf("expected order amount %d, got %d", order.Amount, addedOrder.Amount)
	}
}

func TestListOrders(t *testing.T){
	orderBook := NewBook()

	order1 := Order{
		ID:     "1",
		Type:   Buy,
		Price:  100.0,
		Amount: 5,
	}
	order2 := Order{
		ID:     "2",
		Type:   Sell,
		Price:  105.0,
		Amount: 3,
	}

	initialBuyLength := len(orderBook.ListOrders(Buy))
	orderBook.AddOrder(order1)
	newBuyLength := len(orderBook.ListOrders(Buy))
	if newBuyLength != initialBuyLength+1{
		t.Fatalf("expected buy orders length to increase by 1, but it didn't (initial: %d, new: %d)", initialBuyLength, newBuyLength)
	}

	initialSellLength := len(orderBook.ListOrders(Sell))
	orderBook.AddOrder(order2)
	newSellLength := len(orderBook.ListOrders(Sell))
	if newSellLength != initialSellLength+1{
		t.Fatalf("expected sell orders length to increase by 1, but it didn't (initial: %d, new: %d)", initialSellLength, newSellLength)
	}

	// Test Buy Orders - check the last order added
	buyOrders := orderBook.ListOrders(Buy)
	lastBuyOrder := buyOrders[len(buyOrders)-1]
	if lastBuyOrder.ID != order1.ID {
		t.Errorf("expected buy order ID '%s', got '%s'", order1.ID, lastBuyOrder.ID)
	}

	// Test Sell Orders - check the last order added
	sellOrders := orderBook.ListOrders(Sell)
	lastSellOrder := sellOrders[len(sellOrders)-1]
	if lastSellOrder.ID != order2.ID {
		t.Errorf("expected sell order ID '%s', got '%s'", order2.ID, lastSellOrder.ID)
	}
}