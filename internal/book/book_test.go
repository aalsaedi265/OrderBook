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