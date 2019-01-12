<?php

namespace App\Http\Models;

use EloquentFilter\Filterable;
use Illuminate\Http\JsonResponse;
use Watson\Rememberable\Rememberable;

class Article extends BasicModel
{
    use Filterable,Rememberable;
    protected $table = 'article';
    protected $dateFormat = 'U';
    protected $guarded = [];
    protected $rememberCacheTag = 'h_article_cache';

    //多态关联
    public function images()
    {
        return $this->morphMany(Image::class, 'image');
    }

    //多对多关联文章标签
    public function tags()
    {
        return $this->belongsToMany(Tag::class, 'article_tag', 'article_id', 'tag_id');
    }

    //定义文章和分类的关系
    public function category()
    {
        return $this->belongsTo(Category::class, 'category_id');
    }

    //定义文章和评论的关系
    public function comments()
    {
        return $this->hasMany(ArticleComment::class, 'article_id');
    }

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s', $value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s', $value);
    }

    public function scopeShow($query,int $is_show)
    {
        return $query->where('is_show',$is_show);
    }

    //如果描述为空,默认文章前200个字
    public function setDescriptionAttribute($value)
    {
        if (empty($value))
        {
            $content = strip_tags(request('content_html'));
            $description = mb_strimwidth($content, 0, 200, '...', 'utf-8');
            $this->attributes['description'] = $description;
        }
    }

    //列表查询统计
    public function getArticleCount(array $data)
    {
        $name = $data['search'] ?? null;
        $category_id = $data['category_id'] ?? null;
        return $this->when($name, function ($q) use ($name) {
            return $q->where('title', 'like', '%' . $name . '%');
        })->when($category_id, function ($q) use ($category_id) {
            return $q->where('category_id', $category_id);
        })->count();
    }

    //获取文章列表
    public function getList(array $data=[],bool $flag=false)
    {
        $query = $this->with('tags:tag.id,name','category:category.id,category.title')
            ->withCount('comments');
        return $flag?$query->filter($data)->get()
            :$query->latest('is_top')->latest('created_at')->show(1)->paginate(8);
    }

}
