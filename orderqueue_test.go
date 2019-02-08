package orderbook

import (
	"strconv"
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewOrderList(t *testing.T) {
	orderList := NewOrderQueue(testPrice)

	if !(orderList.length == 0) {
		t.Errorf("Orderlist Length incorrect, got: %d, want: %d.", orderList.length, 0)
	}

	if !(orderList.price.Equal(testPrice)) {
		t.Errorf("Orderlist price incorrect, got: %d, want: %d.", orderList.length, 0)
	}

	if !(orderList.volume.Equal(decimal.Zero)) {
		t.Errorf("Orderlist Length incorrect, got: %d, want: %d.", orderList.length, 0)
	}
}

func TestOrderList(t *testing.T) {
	orderList := NewOrderQueue(testPrice)

	var emptyList OrderQueue
	dummyOrder := make(map[string]string)
	dummyOrder["timestamp"] = testTimestamp.String()
	dummyOrder["quantity"] = testQuanity.String()
	dummyOrder["price"] = testPrice.String()
	dummyOrder["order_id"] = strconv.Itoa(testOrderId)
	dummyOrder["trade_id"] = strconv.Itoa(testTradeId)

	order := NewOrderFromMap(dummyOrder, &emptyList)

	orderList.Append(order)

	if !(orderList.Length() == 1) {
		t.Errorf("Orderlist Length incorrect, got: %d, want: %d.", orderList.length, 0)
	}

	if !(orderList.price.Equal(testPrice)) {
		t.Errorf("Orderlist price incorrect, got: %d, want: %d.", orderList.length, 0)
	}

	if !(orderList.volume.Equal(order.quantity)) {
		t.Errorf("Orderlist Length incorrect, got: %d, want: %d.", orderList.length, 0)
	}

	if !(orderList.volume.Equal(order.quantity)) {
		t.Errorf("Orderlist Length incorrect, got: %d, want: %d.", orderList.length, 0)
	}

	dummyOrder1 := make(map[string]string)
	dummyOrder1["timestamp"] = testTimestamp1.String()
	dummyOrder1["quantity"] = testQuanity1.String()
	dummyOrder1["price"] = testPrice1.String()
	dummyOrder1["order_id"] = strconv.Itoa(testOrderId1)
	dummyOrder1["trade_id"] = strconv.Itoa(testTradeId1)

	order1 := NewOrderFromMap(dummyOrder1, &emptyList)

	orderList.Append(order1)

	if !(orderList.Length() == 2) {
		t.Errorf("Orderlist Length incorrect, got: %d, want: %d.", orderList.length, 0)
	}

	if !(orderList.volume.Equal(order.quantity.Add(order1.quantity))) {
		t.Errorf("Orderlist Length incorrect, got: %d, want: %d.", orderList.length, 0)
	}

	headOrder := orderList.Head()
	if !(headOrder.orderID == "1") {
		t.Errorf("headorder id incorrect, got: %s, want: %d.", headOrder.orderID, 0)
	}

	nextOrder := headOrder.Next()

	if !(nextOrder.orderID == "2") {
		t.Errorf("Next headorder id incorrect, got: %s, want: %d.", headOrder.Next().orderID, 2)
	}
}
