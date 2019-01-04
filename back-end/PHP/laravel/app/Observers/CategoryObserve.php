<?php

namespace App\Observers;

use App\Http\Models\Category;
use Illuminate\Support\Facades\Cache;

class CategoryObserve
{
    //监听分类创建事件
    public function created(Category $category)
    {
        db('category')->where('id',$category->id)->update(['url'=>'/category/'.$category->id]);
        Cache::forget(config('blog.category_cache_key'));

    }

    //监听分类更新事件
    public function updated(Category $category)
    {
        Cache::forget(config('blog.category_cache_key'));
    }

    //监听删除事件
    public function deleted(Category $category)
    {
        Cache::forget(config('blog.category_cache_key'));
    }
}
