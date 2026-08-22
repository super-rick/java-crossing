<?php
// 构建 Slim 应用(工厂函数,便于测试注入 Store)
// 契约(与其他语言实现一致):
//   POST /orders            创建订单 -> 201 {order} / 400 / 409
//   GET  /orders/{id}       查询订单 -> 200 {order} / 404
//   POST /orders/{id}/pay   幂等支付 -> 200 {"paid":true} / 404 / 409
//   GET  /healthz           健康检查 -> 200
//   GET  /metrics           指标文本 -> 200
//
// 注意:withJson() 是 slim/slim 框架扩展;slim/psr7 只提供标准 PSR-7 方法,
// 这里用标准写法(getBody()->write + withHeader)——PSR 标准生态的教学点。
namespace App;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;

function jsonResponse(Response $response, array $data, int $status = 200): Response {
    $response->getBody()->write(json_encode($data, JSON_UNESCAPED_UNICODE));
    return $response
        ->withHeader('Content-Type', 'application/json')
        ->withStatus($status);
}

function createApp(Store $store): \Slim\App {
    $app = AppFactory::create();
    $app->addErrorMiddleware(false, true, true); // 开发模式;生产关掉 detail

    $app->post('/orders', function (Request $request, Response $response) use ($store) {
        $body = json_decode((string)$request->getBody(), true) ?? [];
        $item = $body['item'] ?? null;
        $quantity = $body['quantity'] ?? null;
        if (!is_string($item) || $item === '' || !is_int($quantity) || $quantity <= 0) {
            return jsonResponse($response, ['error' => 'invalid request'], 400);
        }
        try {
            $order = $store->createOrder($item, $quantity);
            return jsonResponse($response, $order, 201);
        } catch (\InvalidArgumentException $e) {
            return jsonResponse($response, ['error' => $e->getMessage()], (int)$e->getCode() ?: 500);
        }
    });

    $app->get('/orders/{id}', function (Request $request, Response $response, array $args) use ($store) {
        try {
            return jsonResponse($response, $store->getOrder((int)$args['id']));
        } catch (\InvalidArgumentException $e) {
            return jsonResponse($response, ['error' => $e->getMessage()], (int)$e->getCode() ?: 500);
        }
    });

    $app->post('/orders/{id}/pay', function (Request $request, Response $response, array $args) use ($store) {
        try {
            $store->payOrder((int)$args['id']);
            return jsonResponse($response, ['paid' => true]);
        } catch (\InvalidArgumentException $e) {
            return jsonResponse($response, ['error' => $e->getMessage()], (int)$e->getCode() ?: 500);
        }
    });

    $app->get('/healthz', fn(Request $request, Response $response) => jsonResponse($response, ['status' => 'ok']));

    $app->get('/metrics', function (Request $request, Response $response) use ($store) {
        $response->getBody()->write("orders_total " . $store->totalOrders() . "\n");
        return $response->withHeader('Content-Type', 'text/plain');
    });

    return $app;
}
