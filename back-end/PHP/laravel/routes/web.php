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
    });

    //微信组
    Route::group(['prefix'=>'wechat'],function (){
        //前台首页路由
        Route::get('/', 'IndexController@index');
    });

    Route::get('/close','IndexController@shutdownWebsite');
    Route::get('search', 'IndexController@searchArticle');       //搜索路由
    Route::get('tag/{id}.html', 'IndexController@tagCloud');         //标签云路由
    Route::post('contact', 'IndexController@contacts');       //留言路由
    Route::get('article/{id}.html', 'ArticleController@article');
    Route::get('/chat', 'IndexController@chat');        //说说路由
    Route::get('category/{id}', 'IndexController@cate');      //导航路由
    Route::get('user/quit', 'OauthController@logout');        //用户退出
    Route::post('comment/store', 'CommentController@commentStore');        //评论存储
    Route::post('link/store', 'IndexController@linkStore');        //友链存储
    Route::get('oauth/redirectToProvider/{service}','OauthController@redirectToProvider');  //三方用户授权
    Route::get('oauth/handleOauth/{service}','OauthController@handleOauth');  //处理授权回调
    Route::post('/home/wechat', 'WeChatController@login')->middleware('cors'); //微信用户登录
    Route::get('/home/wechat/qrcode/{scene}', 'WeChatController@getMiniQrcode'); //获取小程序码
    Route::get('/home/wechat/getScene', 'WeChatController@getScene'); //获取scene
    Route::post('/home/wechat/getStatus', 'WeChatController@getUserStatus'); //获取登录状态
});
