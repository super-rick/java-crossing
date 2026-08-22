<?php
// 入口:nginx 将所有动态请求转发到此文件(前端控制器模式)
require __DIR__ . '/../vendor/autoload.php';

use App\Store;
use function App\createApp;

$app = createApp(new Store());
$app->run();
