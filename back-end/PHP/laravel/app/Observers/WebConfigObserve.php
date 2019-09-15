<?php

namespace App\Observers;

use App\Http\Models\WebConfig;
use Illuminate\Support\Facades\Cache;

class WebConfigObserve
{
    //监听创建事件
    public function created(WebConfig $config)
    {
        $this->updateConfigFile();
        $config->flushCache(config('blog.blog_footer_cache_key'));
        $config->flushCache('h_web_config_cache');
    }

    //监听更新事件
    public function updated(WebConfig $config)
    {
        $this->updateConfigFile();
        $config->flushCache(config('blog.blog_footer_cache_key'));
        $config->flushCache('h_web_config_cache');
    }


    //监听删除事件
    public function deleted(WebConfig $config)
    {
        $this->updateConfigFile();
        $config->flushCache(config('blog.blog_footer_cache_key'));
        $config->flushCache('h_web_config_cache');
    }

    //写入配置文件
    private function updateConfigFile()
    {
        $config_data = WebConfig::type(3)
            ->select('name', 'val')
            ->get()
            ->toArray();
        $config_data = array_combine(array_column($config_data,'name'),array_column($config_data,'val'));
        $path = config_path('webConfig.php');
        $str = '<?php return ' . var_export($config_data, true) . ';';    //将得到数组转换成字符串
        chmod($path,777);
        file_put_contents($path, $str); //写入文件
    }
}
