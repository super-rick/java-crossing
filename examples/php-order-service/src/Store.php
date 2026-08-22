<?php
// 订单存储(单请求内有效)
// 注意:传统 PHP 每个请求独立进程,此 Store 每次请求重建;
// 跨请求共享状态必须外部化(数据库/Redis)——见 README「生产说明」
namespace App;

class Store {
    private array $orders = [];
    private array $inventory = ['coffee' => 100];
    private int $nextId = 1;

    public function createOrder(string $item, int $quantity): array {
        if (($this->inventory[$item] ?? 0) < $quantity) {
            throw new \InvalidArgumentException('insufficient stock', 409);
        }
        $order = ['id' => $this->nextId++, 'item' => $item, 'quantity' => $quantity, 'paid' => false];
        $this->orders[$order['id']] = $order;
        return $order;
    }

    public function getOrder(int $id): array {
        if (!isset($this->orders[$id])) {
            throw new \InvalidArgumentException('order not found', 404);
        }
        return $this->orders[$id];
    }

    public function payOrder(int $id): void {
        if (!isset($this->orders[$id])) {
            throw new \InvalidArgumentException('order not found', 404);
        }
        if ($this->orders[$id]['paid']) {
            return; // 幂等:已支付直接成功
        }
        $item = $this->orders[$id]['item'];
        if (($this->inventory[$item] ?? 0) < $this->orders[$id]['quantity']) {
            throw new \InvalidArgumentException('insufficient stock', 409);
        }
        $this->inventory[$item] -= $this->orders[$id]['quantity'];
        $this->orders[$id]['paid'] = true;
    }

    public function inventoryOf(string $item): int {
        return $this->inventory[$item] ?? 0;
    }

    public function totalOrders(): int {
        return $this->nextId - 1;
    }
}
