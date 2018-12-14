<?php

namespace App\Http\Controllers;

use App\Http\Models\Article;
use App\Http\Models\Category;
use App\Http\Models\Tag;
use App\Http\Requests\ArticleRequest;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class ArticleController extends Controller
{
    //获取文章列表
    public function getList(Request $request,Article $article,Category $category,Tag $tag)
    {
        $request_data = $request->all();
        $list = $article->getList($request_data);
        $category = $category->getArticleCategory();
        $tag = $tag->get(['id','name']);
        if($list->isNotEmpty()){
            $total = $article->getArticleCount($request_data);
            return renderSuccess('获取文章列表成功',compact('total','list','category','tag'));
        }
        return renderError('暂无文章列表数据',compact('category','tag'));
    }

    //删除文章
    public function destroy(Request $request,Article $article)
    {
        $data = $this->validate($request,['id'=>'required|integer']);
        try
        {
            DB::transaction(function () use ($article, $data) {
                $article = $article->find($data['id']);
                $article->tags()->detach();
                $article->delete();
            });
            return renderSuccess('删除文章成功');
        } catch (\Exception $e)
        {
            return renderError('删除文章失败,请稍后再试！');
        }
    }

    //添加文章
    public function store(ArticleRequest $request,Article $article)
    {
        $data = $request->all();
        $tags = $data['tags'];
        unset($data['tags']);
        try
        {
            DB::transaction(function () use ($article, $data, $tags) {
                $res = $article->create($data);
                $res->tags()->attach($tags);
            });
            return renderSuccess('添加文章成功');
        } catch (\Exception $e)
        {
            return renderError('添加文章失败,请稍后再试');
        }
    }

    //修改文章
    public function update(ArticleRequest $request,Article $article)
    {
        $data = $request->all();
        $tags = $data['tags'];
        unset($data['tags']);
        $article = $article->find($data['id']);
        try
        {
            DB::transaction(function () use ($article, $data, $tags) {
                $article->where('id',$data['id'])->update($data);
                $article->tags()->sync($tags);
            });
            return renderSuccess('修改文章成功');
        } catch (\Exception $e)
        {
            return renderError('修改文章失败,请稍后再试！');
        }
    }
}
