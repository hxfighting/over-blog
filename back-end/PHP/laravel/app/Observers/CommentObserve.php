<?php

namespace App\Observers;

use App\Http\Models\Article;
use App\Http\Models\ArticleComment;
use App\Http\Models\User;
use App\Mail\CommentReplyMail;
use Illuminate\Support\Facades\Mail;

class CommentObserve
{
    //监听评论创建事件
    public function created(ArticleComment $comment)
    {
        $reply_user = User::find($comment->reply_id);
        $comment_user_name = User::where('id',$comment->user_id)->value('name');
        $article_title = Article::where('id',$comment->article_id)->value('title');
        Mail::to($reply_user->email)
            ->queue(new CommentReplyMail($reply_user->name,$comment_user_name,$article_title,$comment->content));
    }
}
