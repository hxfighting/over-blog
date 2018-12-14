<?php

namespace App\Observers;

use App\Http\Models\Contact;
use App\Mail\ReplyUserMail;
use Illuminate\Support\Facades\Mail;

class ContactObserve
{
    //监听留言创建事件
    public function created(Contact $contact)
    {
        Mail::to($contact->email)->queue(new ReplyUserMail($contact->name));
    }

    //监听留言更新事件
    public function updated(Contact $contact)
    {
        Mail::to($contact->email)->queue(new ReplyUserMail($contact->name,$contact->reply_content));
    }
}
