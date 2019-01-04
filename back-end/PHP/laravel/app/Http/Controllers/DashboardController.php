<?php

namespace App\Http\Controllers;

use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class DashboardController extends Controller
{
    /**
     * 获取首页统计
     * Date: 2019/1/4 14:13
     * @return JsonResponse
     */
    public function index(): JsonResponse
    {
        $user = db('user')->count();
        $link = db('link')->count();
        $article = db('article')->count();
        $comment = db('article_comment')->count();
        $tag = db('tag')->count();
        $category = db('category')->count();
        $contact = db('contact')->count();
        $chat = db('chat')->count();
        $data = [
            [
                'title' => '用户统计',
                'icon' => 'md-people',
                'count' => $user,
                'color' => '#2d8cf0'
            ],
            [
                'title' => '友联统计',
                'icon' => 'ios-link',
                'count' => $link,
                'color' => '#19be6b'
            ],
            [
                'title' => '文章统计',
                'icon' => 'ios-book',
                'count' => $article,
                'color' => '#ff9900'
            ],
            [
                'title' => '评论统计',
                'icon' => 'ios-chatboxes',
                'count' => $comment,
                'color' => '#ed3f14'
            ],
            [
                'title' => '标签统计',
                'icon' => 'md-pricetags',
                'count' => $tag,
                'color' => '#E46CBB'
            ],
            [
                'title' => '分类统计',
                'icon' => 'md-list',
                'count' => $category,
                'color' => '#9A66E4'
            ],
            [
                'title' => '留言统计',
                'icon' => 'md-mail',
                'count' => $contact,
                'color' => '#FF99CC'
            ],
            [
                'title' => '说说统计',
                'icon' => 'ios-chatbubbles',
                'count' => $chat,
                'color' => '#FFFF00'
            ]
        ];
        return renderSuccess('获取列表统计成功!', $data);
    }
}
