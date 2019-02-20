<?php

return [
    'enablePrettyUrl' => true,
    'showScriptName' => false,
    'rules' => [
        '/' => 'dashboard/index',
        '/api/captcha' => 'basic/captcha',
        '/api/admin/error' => 'error/log-error',
        '/api/admin/login' => 'admin/login',
        '/api/admin' => 'admin/admin-info'
    ],
];
