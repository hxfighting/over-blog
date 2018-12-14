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

Route::get('/', function () {
    $data = \Illuminate\Support\Facades\DB::connection('typ')->table('blog_article_tag')->orderBy('id')->get();
    foreach ($data as $datum)
    {
        $da[] = [
            'article_id'=>$datum->a_id,
            'tag_id'=>$datum->t_id
        ];
    }
    \Illuminate\Support\Facades\DB::table('article_tag')->insert($da);
});
