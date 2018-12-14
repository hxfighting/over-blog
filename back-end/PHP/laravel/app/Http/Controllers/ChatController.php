<?php

namespace App\Http\Controllers;

use App\Http\Models\Chat;
use Illuminate\Http\Request;

class ChatController extends Controller
{
    //获取说说列表数据
    public function getList(Request $request,Chat $chat)
    {
        $list = $chat->filter($request->all())->get();
        return $list->isNotEmpty()?renderSuccess('获取说说列表成功',$list)
            :renderError('暂无说说列表数据');
    }

    //删除说说
    public function destroy(Request $request,Chat $chat)
    {
        $data = $this->validate($request,['id'=>'required|integer']);
        $res = $chat->destroy($data['id']);
        return $res?renderSuccess('删除说说成功')
            :renderError('删除说说失败,请稍后再试！');
    }

    //添加说说
    public function store(Request $request,Chat $chat)
    {
        $data = $this->validate($request,[
            'is_show'=>'required|integer|in:0,1',
            'content'=>'required|min:2|max:255'
        ]);
        $res = $chat->create($data);
        return $res?renderSuccess('添加说说成功')
            :renderError('添加说说失败,请稍后再试！');
    }

    //修改说说
    public function update(Request $request,Chat $chat)
    {
        $data = $this->validate($request,[
            'id'=>'required|integer',
            'content'=>'required|min:2|max:255',
            'is_show'=>'required|integer|in:0,1'
        ]);
        $res = $chat->where('id',$data['id'])->update($data);
        return $res?renderSuccess('修改说说成功')
            :renderError('修改说说失败,请稍后再试！');
    }
}
