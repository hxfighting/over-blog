<?php
return [
    'timeout' => 5.0,

    // 默认发送配置
    'default' => [
        // 网关调用策略，默认：顺序调用
        'strategy' => \Overtrue\EasySms\Strategies\OrderStrategy::class,

        // 默认可用的发送网关
        'gateways' => [
            'aliyun', 'alidayu',
        ],
    ],
    // 可用的网关配置
    'gateways' => [
        'errorlog' => [
            'file' => public_path('easy-sms.log'),
        ],
        'aliyun' => [
            'access_key_id' => env('SMS_KEY'),
            'access_key_secret' => env('SMS_SECRET'),
            'sign_name' => env('SMS_SIGN_NAME'),
        ],
        'alidayu' => [
            'app_key' => env('SMS_KEY'),
            'app_secret' => env('SMS_SECRET'),
            'sign_name' => env('SMS_SIGN_NAME'),
        ],
    ],
];
