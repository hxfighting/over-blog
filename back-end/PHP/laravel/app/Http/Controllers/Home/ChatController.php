<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\Chat;

class ChatController extends BasicController
{
    /**
     * 说说页面
     * Date: 2019/1/15 14:13
     * @param Chat $chat
     * @return \Illuminate\Contracts\View\Factory|\Illuminate\View\View
     */
    public function index(Chat $chat)
    {
        $data = $chat->latest()->show(1)->get();
        return view('home.chat', compact('data'));
    }
}
