<?php

namespace App\Http\Models;

use EloquentFilter\Filterable;
use Illuminate\Http\JsonResponse;

class Article extends BasicModel
{
    use Filterable;
    protected $table = 'article';
    protected $dateFormat = 'U';
    protected $guarded = [];

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

}
