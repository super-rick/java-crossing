// 订单存储:内存实现,并发安全(sync.Mutex 保护,≈ ConcurrentHashMap 语义)
package main

import (
	"errors"
	"sync"
)

type Order struct {
	ID       int64  `json:"id"`
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
	Paid     bool   `json:"paid"`
}

type Store struct {
	mu        sync.Mutex
	orders    map[int64]*Order
	inventory map[string]int
	nextID    int64
}

func NewStore() *Store {
	return &Store{
		orders:    make(map[int64]*Order),
		inventory: map[string]int{"coffee": 100},
		nextID:    1,
	}
}

// CreateOrder:创建订单并锁定库存(≈ 事务的预占)
func (s *Store) CreateOrder(item string, qty int) (Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.inventory[item] < qty {
		return Order{}, errors.New("insufficient stock")
	}
	o := Order{ID: s.nextID, Item: item, Quantity: qty, Paid: false}
	s.nextID++
	s.orders[o.ID] = &o
	return o, nil
}

func (s *Store) GetOrder(id int64) (Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	if !ok {
		return Order{}, errors.New("order not found")
	}
	return *o, nil
}

// PayOrder:幂等支付。首次调用扣减库存并置 Paid;
// 重复调用(并发场景)返回 nil,不重复扣减 —— 用 Mutex 保证原子性
func (s *Store) PayOrder(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	if !ok {
		return errors.New("order not found")
	}
	if o.Paid {
		return nil // 幂等:已支付直接成功
	}
	if s.inventory[o.Item] < o.Quantity {
		return errors.New("insufficient stock")
	}
	s.inventory[o.Item] -= o.Quantity
	o.Paid = true
	return nil
}

func (s *Store) Inventory(item string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.inventory[item]
}

func (s *Store) TotalOrders() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.nextID - 1
}
