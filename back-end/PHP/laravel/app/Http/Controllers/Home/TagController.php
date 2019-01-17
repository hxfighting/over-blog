<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\Article;
use App\Http\Models\Tag;

class TagController extends BasicController
{
    /**
     * 标签相关文章
     * Date: 2019/1/16 10:51
     * @param         $id
     * @param Tag     $tag
     * @param Article $article
     * @return \Illuminate\Contracts\View\Factory|\Illuminate\View\View
     */
    public function index($id,Tag $tag,Article $article)
    {
        //获取标签名字
        $tag_name = $tag->where('id', $id)->value('name');
        //获取拥有该标签的文章id
        $list = $article->with('tags:tag.id,name', 'category:category.id,category.title')
            ->whereHas('tags',function ($q) use ($id){
            return $q->where('tag.id',$id);
        })->latest()->paginate(7);
        return view('home.category', compact('list', 'tag_name'));
    }
}
