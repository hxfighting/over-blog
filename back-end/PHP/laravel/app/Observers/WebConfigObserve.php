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
        $this->updateConfigFile();
    }

    //监听更新事件
    public function updated(WebConfig $config)
    {
        Cache::forget(config('blog.category_cache_key'));
        $this->updateConfigFile();
    }

    //监听删除事件
    public function deleted(WebConfig $config)
    {
        Cache::forget(config('blog.category_cache_key'));
        $this->updateConfigFile();
    }

    //写入配置文件
    private function updateConfigFile()
    {
        $config_data = WebConfig::where('type', 3)
            ->select('name', 'val')
            ->get()
            ->toArray();
        $path = config_path('webConfig.php');
        $str = '<?php return ' . var_export($config_data, true) . ';';    //将得到数组转换成字符串
        file_put_contents($path, $str); //写入文件
    }
}
