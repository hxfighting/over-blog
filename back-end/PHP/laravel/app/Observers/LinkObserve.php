<?php

namespace App\Observers;

use App\Http\Models\Link;

class LinkObserve
{
    //监听创建事件
    public function created(Link $link)
    {
        $link->flushCache();
    }

    //监听更新事件
    public function updated(Link $link)
    {
        $link->flushCache();
    }

    //监听删除事件
    public function deleted(Link $link)
    {
        $link->flushCache();
    }
}
