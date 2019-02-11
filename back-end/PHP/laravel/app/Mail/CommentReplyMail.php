<?php

namespace App\Mail;

use Illuminate\Bus\Queueable;
use Illuminate\Mail\Mailable;
use Illuminate\Queue\SerializesModels;
use Illuminate\Contracts\Queue\ShouldQueue;

class CommentReplyMail extends Mailable implements ShouldQueue
{
    use Queueable, SerializesModels;

    public $subject;

    public $reply_content;

    public $reply_name;

    public $comment_name;

    public $title;

    public $url;

    /**
     * Create a new message instance.
     *
     * @return void
     */
    public function __construct($reply_name,$comment_name,$title,$reply_content,$url)
    {
        $this->subject = config('app.name').'评论回复';
        $this->reply_name = $reply_name;
        $this->comment_name = $comment_name;
        $this->reply_content = $reply_content;
        $this->title = $title;
        $this->url = $url;
    }

    /**
     * Build the message.
     *
     * @return $this
     */
    public function build()
    {
        return $this->markdown('commentReply');
    }
}
