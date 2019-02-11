<?php

namespace App\Mail;

use Illuminate\Bus\Queueable;
use Illuminate\Mail\Mailable;
use Illuminate\Queue\SerializesModels;
use Illuminate\Contracts\Queue\ShouldQueue;

class CommentMail extends Mailable
{
    use Queueable, SerializesModels;

    public $subject;

    public $title;

    public $username;

    public $url;

    public $content;

    /**
     * Create a new message instance.
     *
     * @return void
     */
    public function __construct(string $title,string $username,string $url,string $content)
    {
        $this->subject = config('app.name');
        $this->title = $title;
        $this->username = $username;
        $this->url = $url;
        $this->content = $content;
    }

    /**
     * Build the message.
     *
     * @return $this
     */
    public function build()
    {
        return $this->markdown('CommentMail');
    }
}
