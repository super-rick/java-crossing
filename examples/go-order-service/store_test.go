// TDD 第一步:并发扣减正确性测试(先红后绿)
package main

import (
	"sync"
	"testing"
)

func TestCreateAndPay(t *testing.T) {
	s := NewStore()
	o, err := s.CreateOrder("coffee", 2)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.PayOrder(o.ID); err != nil {
		t.Fatal(err)
	}
	got, _ := s.GetOrder(o.ID)
	if !got.Paid {
		t.Error("order should be paid")
	}
	if s.Inventory("coffee") != 98 {
		t.Errorf("stock = %d, want 98", s.Inventory("coffee"))
	}
}

// 并发幂等:100 个并发支付同一订单,只扣一次库存,其余幂等成功
func TestConcurrentPayIdempotent(t *testing.T) {
	s := NewStore()
	o, _ := s.CreateOrder("coffee", 5)

	var wg sync.WaitGroup
	errs := make([]error, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errs[i] = s.PayOrder(o.ID)
		}(i)
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			t.Fatalf("pay #%d failed: %v", i, err)
		}
	}
	if got := s.Inventory("coffee"); got != 95 {
		t.Errorf("stock = %d, want 95(只扣一次)", got)
	}
}

func TestInsufficientStock(t *testing.T) {
	s := NewStore()
	o, _ := s.CreateOrder("coffee", 1000)
	if err := s.PayOrder(o.ID); err == nil {
		t.Error("expected insufficient stock error")
	}
}

func TestPayNotFound(t *testing.T) {
	s := NewStore()
	if err := s.PayOrder(999); err == nil {
		t.Error("expected not found error")
	}
}
