<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019/1/16
 * Time: 15:39
 */

return [
    'weibo' => [
        'client_id' => env('WEIBO_KEY'),
        'client_secret' => env('WEIBO_SECRET'),
        'redirect' => env('APP_URL') . '/oauth/handleOauth/weibo'
    ],
    'qq' => [
        'client_id' => env('QQ_KEY'),
        'client_secret' => env('QQ_SECRET'),
        'redirect' => env('APP_URL') . '/oauth/handleOauth/qq'
    ],
];
