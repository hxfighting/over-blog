<?php
$web_config = require __DIR__.'/webConfig.php';
return [
    /**
     * jwt iss claim
     */
    'issuer' => '',
    /**
     * jwt aud claim
     */
    'audience' => '',
    /**
     * jwt jti claim
     */
    'jwt_id' => '',
    /**
     * jwt 过期时间（秒）
     */
    'jwt_expire' => 0,
    /**
     * jwt_sign
     */
    'jwt_sign' => '',
    /**
     * blog name
     */
    'blog_name' => '拖油瓶博客',
    /**
     * blog index 博客首页URL
     */
    'blog_index_url' => 'https://www.ohdata.top',
    /**
     * 网站配置
     */
    'web_config'=> $web_config
];
