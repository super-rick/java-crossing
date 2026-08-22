<?php
// PHPUnit 测试:单请求内的幂等/库存逻辑
// 并发语义说明:传统 PHP 每请求独立进程,不存在进程内并发;
// 跨请求「并发扣减」在生产中是数据库/Redis 原子操作问题(见 README)
namespace App\Tests;

use App\Store;
use PHPUnit\Framework\TestCase;

class StoreTest extends TestCase {
    public function testCreateGetPay(): void {
        $s = new Store();
        $order = $s->createOrder('coffee', 2);
        $this->assertSame(1, $order['id']);
        $this->assertFalse($order['paid']);

        $s->payOrder($order['id']);
        $this->assertTrue($s->getOrder(1)['paid']);
        $this->assertSame(98, $s->inventoryOf('coffee'));
    }

    public function testPayIdempotent(): void {
        $s = new Store();
        $order = $s->createOrder('coffee', 5);
        $s->payOrder($order['id']);
        $s->payOrder($order['id']); // 重复支付幂等
        $this->assertSame(95, $s->inventoryOf('coffee')); // 只扣一次
    }

    public function testRepeatedPayIdempotent100(): void {
        // 模拟「多次重复支付」:单进程内 100 次调用,库存只扣一次
        $s = new Store();
        $order = $s->createOrder('coffee', 5);
        for ($i = 0; $i < 100; $i++) {
            $s->payOrder($order['id']);
        }
        $this->assertSame(95, $s->inventoryOf('coffee'));
    }

    public function testInsufficientStock(): void {
        $s = new Store();
        $this->expectException(\InvalidArgumentException::class);
        $s->createOrder('coffee', 9999);
    }

    public function testPayNotFound(): void {
        $s = new Store();
        $this->expectException(\InvalidArgumentException::class);
        $s->payOrder(999);
    }
}
