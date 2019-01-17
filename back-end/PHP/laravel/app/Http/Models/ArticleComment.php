<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;
use Watson\Rememberable\Rememberable;

class ArticleComment extends BasicModel
{
    use Filterable,Rememberable;
    protected $table = 'article_comment';
    protected $dateFormat = 'U';
    protected $guarded = [];
    protected $rememberCacheTag = 'h_comment_cache';

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    //定义和文章的关系
    public function articles()
    {
        return $this->belongsTo(Article::class,'article_id');
    }

    //定义子评论关系
    public function children()
    {
        return $this->hasMany(self::class,'pid');
    }

    //定义评论人关系
    public function user()
    {
        return $this->belongsTo(User::class,'user_id');
    }

    //定义回复人关系
    public function replier()
    {
        return $this->belongsTo(User::class,'reply_id');
    }

    //列表查询统计
    public function getCommentCount(array $data)
    {
        $article_id = $data['article_id'] ?? null;
        return $this->when($article_id, function ($q) use ($article_id) {
            return $q->where('article_id', $article_id);
        })->count();
    }

    //后台回复评论
    public function handleReplyComment($user,$data)
    {
        $exist_comment = $this->find($data['id']);
        $user_id = !$exist_comment->reply_id
            ? $user->where('is_admin', 1)->latest()->value('id')
            : $exist_comment->reply_id;
        if($user_id){
            $pid = $exist_comment->pid ? $exist_comment->pid : $exist_comment->user_id;
            $this->content = $data['reply_content'];
            $this->pid = $pid;
            $this->user_id = $user_id;
            $this->reply_id = $exist_comment->user_id;
            $this->article_id = $exist_comment->article_id;
            $res = $this->save();
            return $res?renderSuccess('回复评论成功'):renderError('回复评论失败,请稍后再试');
        }
        return renderError('回复评论失败,缺少前台管理员用户,请先绑定前台管理员');
    }
}
