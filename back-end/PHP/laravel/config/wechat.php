<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019/1/16
 * Time: 13:42
 */
return [
    'app_id'  => env('WECHAT_APP_ID',''),         // AppID
    'secret'  => env('WECHAT_APP_SECRET',''),     // AppSecret
    // 下面为可选项
    // 指定 API 调用返回结果的类型：array(default)/collection/object/raw/自定义类名
    'response_type' => 'array',

    'log' => [
        'default' => 'dev', // 默认使用的 channel，生产环境可以改为下面的 prod
        'channels' => [
            // 测试环境
            'dev' => [
                'driver' => 'single',
                'path' => storage_path('logs/easywechat.log'),
                'level' => 'debug',
            ],
            // 生产环境
            'prod' => [
                'driver' => 'daily',
                'path' => storage_path('logs/easywechat.log'),
                'level' => 'info',
            ]
        ]
    ],
];
