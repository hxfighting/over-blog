<?php

namespace App\Http\Models;


class Category extends BasicModel
{
    protected $table = 'category';
    protected $dateFormat = 'U';
    protected $fillable = ['pid','title'];

    public function children()
    {
        return $this->hasMany(self::class,'pid','id');
    }

    //获取文章分类
    public function getArticleCategory()
    {
        $data = $this->where('is_article_category',1)->select('id','title')->get();
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
