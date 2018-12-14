<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\Article;
use App\Http\Models\Category;
use App\Http\Models\Image;
use App\Http\Models\Photo;
use App\Http\Models\RotationImage;
use Illuminate\Http\Request;
use App\Http\Controllers\Controller;

class IndexController extends Controller
{
    //博客首页分类
    public function category(Category $category)
    {
        $list = $category->getTree();
        return renderSuccess('获取分类数据成功',$list);
    }

    //博客首页轮播图
    public function rotation(Image $image)
    {
        $rotation = $image->rotationImage();
        $photo = $image->photoImage();
        return renderSuccess('获取轮播图数据成功',compact('rotation','photo'));
    }

    //获取文章列表
    public function getArticleList(Request $request,Article $article)
    {
        $data = $request->all();
        $data['order'] = 'is_top';
        $data['show'] = 1;
        $list = $article->getList($data);
        return $list->isNotEmpty()?renderSuccess('获取文章列表成功',$list)
            :renderError('暂无文章列表数据');
    }
}
