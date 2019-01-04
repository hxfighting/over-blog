<?php

namespace App\Observers;

use App\Http\Models\WebConfig;
use Illuminate\Support\Facades\Cache;

class WebConfigObserve
{
    //监听创建事件
    public function created(WebConfig $config)
    {
        Cache::forget(config('blog.category_cache_key'));
    }

    //监听更新事件
    public function updated(WebConfig $config)
    {
        Cache::forget(config('blog.category_cache_key'));
    }

    //监听删除事件
    public function deleted(WebConfig $config)
    {
        Cache::forget(config('blog.category_cache_key'));
    }
}
