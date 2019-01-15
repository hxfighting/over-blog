<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\Article;

class CategoryController extends BasicController
{
    /**
     * 分类文章页面
     * Date: 2019/1/15 14:51
     * @param         $id
     * @param Article $article
     * @return \Illuminate\Contracts\View\Factory|\Illuminate\View\View
     */
    public function index($id,Article $article)
    {
        $list = $article->where('category_id', $id)->latest()->paginate(6);
        return $list->isNotEmpty()?view('home.category', compact('list'))
            :view('home.kong');
    }
}
