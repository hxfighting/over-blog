<?php

namespace App\Listeners;

use App\Events\NotifyUserEvent;
use App\Http\Models\Article;
use App\Http\Models\User;
use App\Mail\CommentMail;
use App\Mail\CommentReplyMail;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Mail;

class NotifyUserListener implements ShouldQueue
{
    /**
     * Create the event listener.
     *
     * @return void
     */
    public function __construct()
    {
        //
    }

    /**
     * Created by huxin.
     * Date: 2019-01-30 17:19
     * @param NotifyUserEvent $event
     */
    public function handle(NotifyUserEvent $event)
    {
        try
        {
            $data = $event->data;
            $comment_user_name = User::where('id', $data['user_id'])->value('name');
            $article_title = Article::where('id', $data['article_id'])->value('title');
            $url = config('app.url') . '/article/' . $data['article_id'];
            if (isset($data['reply_id']) && $data['reply_id'])
            {
                $reply_user = User::find($data['reply_id']);
                Mail::to($reply_user->email)
                    ->queue(new CommentReplyMail($reply_user->name, $comment_user_name, $article_title, $data['content'],$url));
            } else
            {
                Mail::to(config('mail.from.address'))
                    ->queue(new CommentMail($article_title, $comment_user_name, $url, $data['content']));
            }
        } catch (\Exception $e)
        {
            Log::error($e->getMessage());
        }
    }
}
