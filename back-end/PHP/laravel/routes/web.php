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
        Route::get('article/{id}.html', 'ArticleController@index');
    });

    //微信组
    Route::group(['prefix'=>'wechat'],function (){
        //前台首页路由
        Route::get('/', 'IndexController@index');
    });
});
