<?php

use Illuminate\Http\Request;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/


Route::get('/captcha','Controller@getCaptcha');

//后台接口
Route::group(['prefix'=>'admin'],function (){
    //需要登录的接口
    Route::group(['middleware'=>'admin.token'],function (){

        //图片上传
        Route::post('/upload','UploadController@upload');

        //错误日志组
        Route::group(['prefix'=>'error'],function (){
            //获取错误日志列表
            Route::get('/','WebErrorController@getErrorList');
            //删除错误日志
            Route::delete('/','WebErrorController@destroy');
        });

        //管理员组
        Route::group([],function (){
            //获取管理员信息
            Route::get('/','AdminController@getAdminInfo');
            //修改个人信息
            Route::put('/','AdminController@updateInfo');
            //退出登录
            Route::post('/logout','AdminController@logout');
            //修改密码
            Route::put('/password','AdminController@resetPassword');
        });

        //文章组
        Route::group(['prefix'=>'article'],function (){
            //文章列表
            Route::get('/','ArticleController@getList');
            //添加文章
            Route::post('/','ArticleController@store');
            //修改文章
            Route::put('/','ArticleController@update');
            //删除文章
            Route::delete('/','ArticleController@destroy');
        });

        //标签组
        Route::group(['prefix'=>'tag'],function (){
            //标签列表
            Route::get('/','TagController@getList');
            //添加标签
            Route::post('/','TagController@store');
            //修改标签
            Route::put('/','TagController@update');
            //删除标签
            Route::delete('/','TagController@destroy');
        });

        //分类组
        Route::group(['prefix'=>'category'],function (){
            //分类列表
            Route::get('/','CategoryController@list');
            //添加分类
            Route::post('/','CategoryController@store');
            //修改分类
            Route::put('/','CategoryController@update');
            //删除分类
            Route::delete('/','CategoryController@destroy');
        });

        //说说组
        Route::group(['prefix'=>'chat'],function (){
            //说说列表
            Route::get('/','ChatController@getList');
            //添加说说
            Route::post('/','ChatController@store');
            //修改说说
            Route::put('/','ChatController@update');
            //删除说说
            Route::delete('/','ChatController@destroy');
        });

        //轮播图组
        Route::group(['prefix'=>'rotation'],function (){
            //轮播图列表
            Route::get('/','RotationController@getList');
            //添加轮播图
            Route::post('/','RotationController@store');
            //修改轮播图
            Route::put('/','RotationController@update');
            //删除轮播图
            Route::delete('/','RotationController@destroy');
        });

        //照片组
        Route::group(['prefix'=>'photo'],function (){
            //照片列表
            Route::get('/','PhotoController@getList');
            //添加照片
            Route::post('/','PhotoController@store');
            //修改照片
            Route::put('/','PhotoController@update');
            //删除照片
            Route::delete('/','PhotoController@destroy');
        });

        //友联组
        Route::group(['prefix'=>'link'],function (){
            //友联列表
            Route::get('/','LinkController@getList');
            //添加友联
            Route::post('/','LinkController@store');
            //修改友联
            Route::put('/','LinkController@update');
            //删除友联
            Route::delete('/','LinkController@destroy');
        });

        //留言组
        Route::group(['prefix'=>'contact'],function (){
            //友联列表
            Route::get('/','ContactController@getList');
            //删除友联
            Route::delete('/','ContactController@destroy');
            //回复留言
            Route::post('/reply','ContactController@reply');
        });

        //用户组
        Route::group(['prefix'=>'user'],function (){
            //用户列表
            Route::get('/','UserController@getList');
            //删除用户
            Route::delete('/','UserController@destroy');
            //修改用户为管理员
            Route::put('/','UserController@update');
        });

        //评论组
        Route::group(['prefix'=>'comment'],function (){
            //评论列表
            Route::get('/','CommentController@getList');
            //回复评论
            Route::post('/','CommentController@reply');
            //删除评论
            Route::delete('/','CommentController@destroy');
        });

    });

    //不需要登录的接口
    Route::group([],function (){
        //管理员登录
        Route::post('/login','AdminController@login');
        //记录后台接口错误日志
        Route::post('/error','WebErrorController@logError');
    });
});

//前台接口
Route::group(['prefix'=>'home','namespace'=>'Home'],function(){

    //需要登录的接口
    Route::group([],function(){

    });

    //不需要登录的接口
    Route::group([],function(){
        //首页分类
        Route::get('/category','IndexController@category');
        //首页轮播图
        Route::get('/rotation','IndexController@rotation');
    });
});
