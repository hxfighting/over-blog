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
            $view_cookie = $request->cookie('view_index', '');
            if (!empty($view_cookie))
            {
                list($rtime, $rip) = explode('#', $view_cookie);
                if ($this->calculateTimeDiff($rtime, $time) >= 1)
                {
                    Cache::increment(config('blog.blog_index_count_key'));
                }
            } else
            {
                Cache::increment(config('blog.blog_index_count_key'));
            }
            $view_cookie = $time . '#' . $ip;
        } else
        {
            $view_cookie = $request->cookie('view_article', '');
            $id = $request->route()->parameter('id');
            if (!empty($view_cookie))
            {
                list($rtime, $aid) = explode('#', $view_cookie);
                $aid = explode(',', $aid);
                if ($this->calculateTimeDiff($rtime, $time) >= 1 || !in_array($id, $aid))
                {
                    Cache::increment(config('blog_article_count_key'));
                }
                $aid[] = $id;
                $aid = array_unique($aid);
                $view_cookie = $time . '#' . implode(',', $aid);
            } else
            {
                $view_cookie = $time . '#' . $id;
                Cache::increment(config('blog_article_count_key'));
            }
        }
        return $view_cookie;
    }

    //计算时间
    private function calculateTimeDiff($start, $end)
    {
        return floor(($end - $start) / 86400);
    }

}
