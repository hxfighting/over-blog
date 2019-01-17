<?php

namespace App\Http\Models;


use Watson\Rememberable\Rememberable;

class Category extends BasicModel
{
    use Rememberable;
    protected $table = 'category';
    protected $dateFormat = 'U';
    protected $fillable = ['pid','title'];
    protected $rememberCacheTag = 'h_category_cache';

    public function children()
    {
        return $this->hasMany(self::class,'pid','id');
    }

    //获取文章分类
    public function getArticleCategory()
    {
        $data = $this->where('type',1)->select('id','title')->get();
        return $data->isNotEmpty()?$data:null;
    }

    //获取分类
    public function getTree()
    {
        return $this->with('children')
            ->where('pid',0)
            ->get();
    }
}
