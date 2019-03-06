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

        'GET api/admin/count' => 'dashboard/dashboard-count',   //后台首页获取统计数据

        /**
         * 标签组
         */
        'GET api/admin/tag' => 'tag/tag-list',                  //获取标签列表
        'POST api/admin/tag' => 'tag/add-tag',                  //添加标签
        'PUT api/admin/tag' => 'tag/update-tag',                //修改标签
        'DELETE api/admin/tag' => 'tag/del-tag',                //删除标签

        /**
         * 分类组
         */
        'GET api/admin/category' => 'category/category-list',   //获取分类列表
        'POST api/admin/category' => 'category/category-add',   //添加分类
        'PUT api/admin/category' => 'category/category-update', //修改分类
        'DELETE api/admin/category' => 'category/del-category', //删除分类

        /**
         * 说说组
         */
        'GET api/admin/chat' => 'chat/chat-list',               //获取说说列表
        'POST api/admin/chat' => 'chat/chat-add',               //添加说说
        'PUT api/admin/chat' => 'chat/chat-update',             //修改说说
        'DELETE api/admin/chat' => 'chat/del-chat',             //删除说说

        /**
         * 轮播图组
         */
        'GET api/admin/rotation' => 'rotation/rotation-list',   //获取轮播图列表
        'POST api/admin/rotation' => 'rotation/rotation-add',   //添加轮播图
        'PUT api/admin/rotation' => 'rotation/rotation-update', //修改轮播图
        'DELETE api/admin/rotation' => 'rotation/del-rotation', //删除轮播图

        /**
         * 照片组
         */
        'GET api/admin/photo' => 'photo/photo-list',            //获取照片列表
        'POST api/admin/photo' => 'photo/photo-add',            //添加照片
        'PUT api/admin/photo' => 'photo/photo-update',          //修改照片
        'DELETE api/admin/photo' => 'photo/del-photo',          //删除照片

        /**
         * 留言组
         */
        'GET api/admin/contact' => 'contact/contact-list',      //获取留言列表
        'DELETE api/admin/contact' => 'contact/del-contact',    //删除留言
        'POST api/admin/contact/reply' => 'contact/reply',      //回复留言
    ],
];
