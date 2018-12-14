<?php

namespace App\Mail;

use Illuminate\Bus\Queueable;
use Illuminate\Mail\Mailable;
use Illuminate\Queue\SerializesModels;
use Illuminate\Contracts\Queue\ShouldQueue;

class ReplyUserMail extends Mailable implements ShouldQueue
{
    use Queueable, SerializesModels;

    public $subject = '拖油瓶博客留言回复';

    public $reply_content;

    public $name;
    /**
     * Create a new message instance.
     *
     * @return void
     */
    public function __construct($name,$reply_content = null)
    {
        $this->name = $name;
        $this->reply_content = $reply_content ?? '非常感谢你的留言,我会尽快回复你的.';
    }

    /**
     * Build the message.
     *
     * @return $this
     */
    public function build()
    {
        return $this->markdown('reply');
    }
}
