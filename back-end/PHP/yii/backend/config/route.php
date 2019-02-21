<?php

return [
    'enablePrettyUrl' => true,
    'showScriptName' => false,
    'rules' => [
        'GET /' => 'dashboard/index',
        'GET api/captcha' => 'basic/captcha',                   //获取验证码
        'GET api/admin' => 'admin/admin-info',                  //登录获取管理员信息
        'PUT api/admin' => 'admin/change-info',                 //管理员修改个人信息
        'PUT api/admin/password' => 'admin/password',           //管理员修改密码
        'POST api/admin/error' => 'error/log-error',            //记录网站错误
        'DELETE api/admin/error' => 'error/del-error',          //删除网站错误
        'POST api/admin/login' => 'admin/login',                //管理员登录
        'POST api/admin/logout' => 'admin/logout',              //管理员退出登录

        'GET api/admin/count' => 'dashboard/dashboard-count'    //后台首页获取统计数据
    ],
];
