<?php

namespace App\Observers;

use App\Http\Models\Category;

class CategoryObserve
{
    //监听分类创建事件
    public function created(Category $category)
    {
        db('category')->where('id',$category->id)->update(['url'=>'/category/'.$category->id]);
    }
}
