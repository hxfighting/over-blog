<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019/1/8
 * Time: 15:54
 */

namespace App\Http\Controllers\Home;


use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Cache;

class BasicController extends Controller
{

    //统计页面访问次数
    public function countView($request)
    {
        $ip = getIP();
        $time = time();
        $page_name = $request->route()->getName();
        if ($page_name == 'index')
        {
            $key = config('blog.blog_index_count_key');
            $view_cookie = $request->cookie('view_index', '');
            if (!empty($view_cookie))
            {
                list($rtime, $rip) = explode('#', $view_cookie);
                if ($this->calculateTimeDiff($rtime, $time) >= 1)
                {
                    $this->handleCacheData($key);
                }
            } else
            {
                $this->handleCacheData($key);
            }
            $view_cookie = $time . '#' . $ip;
        } else
        {
            $view_cookie = $request->cookie('view_article', '');
            $id = $request->route()->parameter('id');
            $key = config('blog.blog_article_count_key');
            if (!empty($view_cookie))
            {
                list($rtime, $aid) = explode('#', $view_cookie);
                $aid = explode(',', $aid);
                if ($this->calculateTimeDiff($rtime, $time) >= 1 || !in_array($id, $aid))
                {
                    $this->handleCacheData($key);
                }
                $aid[] = $id;
                $aid = array_unique($aid);
                $view_cookie = $time . '#' . implode(',', $aid);
            } else
            {
                $view_cookie = $time . '#' . $id;
                $this->handleCacheData($key);
            }
        }
        return $view_cookie;
    }

    //计算时间
    private function calculateTimeDiff($start, $end)
    {
        return floor(($end - $start) / 86400);
    }

    //处理缓存
    private function handleCacheData($key)
    {
        if(Cache::has($key)){
            Cache::increment($key);
        }else{
            $val = db('web_config')
                ->where('name','blog_view_count')
                ->value('val');
            Cache::put($key,$val,86400);
        }
    }

}
