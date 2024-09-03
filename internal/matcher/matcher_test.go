package matcher

import (
	"testing"
	"orderbook_tradingEngine/internal/book"
)

func TestMatchOrders_FullMatch(t *testing.T) {
	orderBook := book.NewBook()

	// Add a buy order and a matching sell order
	orderBook.AddOrder(book.Order{
		ID:     "1",
		Type:   book.Buy,
		Price:  100.0,
		Amount: 10,
	})
	orderBook.AddOrder(book.Order{
		ID:     "2",
		Type:   book.Sell,
		Price:  100.0,
		Amount: 10,
	})

	// Perform matching
	MatchOrders(orderBook)

	// Check the final state of the order book
	if len(orderBook.BuyOrders()) != 0 {
		t.Errorf("expected no buy orders, got %d", len(orderBook.BuyOrders()))
	}

	if len(orderBook.SellOrders()) != 0 {
		t.Errorf("expected no sell orders, got %d", len(orderBook.SellOrders()))
	}
}

func TestMatchOrders_PartialMatch(t *testing.T) {
	orderBook := book.NewBook()

	// Add a buy order and a partially matching sell order
	orderBook.AddOrder(book.Order{
		ID:     "1",
		Type:   book.Buy,
		Price:  100.0,
		Amount: 10,
	})
	orderBook.AddOrder(book.Order{
		ID:     "2",
		Type:   book.Sell,
		Price:  100.0,
		Amount: 5,
	})

	// Perform matching
	MatchOrders(orderBook)

	// Check the final state of the order book
	if len(orderBook.BuyOrders()) != 1 {
		t.Fatalf("expected 1 buy order, got %d", len(orderBook.BuyOrders()))
	}

	remainingBuyOrder := orderBook.BuyOrders()[0]
	if remainingBuyOrder.Amount != 5 {
		t.Errorf("expected 5 remaining on buy order, got %d", remainingBuyOrder.Amount)
	}

	if len(orderBook.SellOrders()) != 0 {
		t.Errorf("expected no sell orders, got %d", len(orderBook.SellOrders()))
	}
}

func TestMatchOrders_NoMatch(t *testing.T) {
	orderBook := book.NewBook()

	// Add a buy order and a sell order with non-matching prices
	orderBook.AddOrder(book.Order{
		ID:     "1",
		Type:   book.Buy,
		Price:  90.0,
		Amount: 10,
	})
	orderBook.AddOrder(book.Order{
		ID:     "2",
		Type:   book.Sell,
		Price:  100.0,
		Amount: 10,
	})

	// Perform matching
	MatchOrders(orderBook)

	// Check the final state of the order book
	if len(orderBook.BuyOrders()) != 1 {
		t.Errorf("expected 1 buy order, got %d", len(orderBook.BuyOrders()))
	}
	if len(orderBook.SellOrders()) != 1 {
		t.Errorf("expected 1 sell order, got %d", len(orderBook.SellOrders()))
	}
	if orderBook.BuyOrders()[0].Amount != 10 {
		t.Errorf("expected 10 remaining on buy order, got %d", orderBook.BuyOrders()[0].Amount)
	}
	if orderBook.SellOrders()[0].Amount != 10 {
		t.Errorf("expected 10 remaining on sell order, got %d", orderBook.SellOrders()[0].Amount)
	}
}
