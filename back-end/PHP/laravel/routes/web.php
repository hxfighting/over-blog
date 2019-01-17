<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/
Route::group(['namespace'=>'Home'],function(){

    Route::group([],function (){
        //前台首页路由
        Route::get('/', 'IndexController@index')->name('index');
        //文章详情页
        Route::get('/article/{id}.html', 'ArticleController@index')->where('id','\d+');
        //说说页面
        Route::get('/chat','ChatController@index');
        //分类文章页面
        Route::get('/category/{id}.html','CategoryController@index')->where('id','\d+');
        //标签文章页面
        Route::get('/tag/{id}.html','TagController@index')->where('id','\d+');
        //搜索文章
        Route::get('/search','ArticleController@search');
        //添加友联
        Route::post('/link','LinkController@store');
    });

    //微信组
    Route::group(['prefix'=>'wechat'],function (){
        //前台首页路由
        Route::get('/', 'IndexController@index');
        //获取小程序scene
        Route::get('/scene', 'WeChatController@getScene');
        //获取小程序码
        Route::get('/qrcode/{scene}', 'WeChatController@getQrCode')->where('scene','/^([a-fA-F0-9]{32})$/');
        //获取微信登录状态
        Route::get('/status', 'WeChatController@getLoginResult');
    });

    //微信小程序登录
    Route::post('/home/wechat', 'WeChatController@login');

    //三方登录组
    Route::group(['prefix'=>'oauth'],function (){
        //三方用户授权
        Route::get('/redirectToProvider/{service}','OauthController@oauth');
        //处理授权回调
        Route::get('/handleOauth/{service}','OauthController@callback');
    });

});
