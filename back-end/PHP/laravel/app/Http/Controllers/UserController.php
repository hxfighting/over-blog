<?php

namespace App\Http\Controllers;

use App\Http\Models\User;
use Illuminate\Http\Request;

class UserController extends Controller
{
    //获取会员列表数据
    public function getList(Request $request,User $user)
    {
        $request_data = $request->all();
        $list = $user->filter($request_data)->get();
        if($list->isNotEmpty()){
            $total = $user->getUserCount($request_data);
            return renderSuccess('获取会员列表成功',compact('total','list'));
        }
        return renderError('暂无会员列表数据');
    }

    //删除会员
    public function destroy(Request $request,User $user)
    {
        $data = $this->validate($request,['id'=>'required|integer']);
        $res = $user->destroy($data['id']);
        return $res?renderSuccess('删除会员成功')
            :renderError('删除会员失败,请稍后再试！');
    }

    //修改会员
    public function update(Request $request,User $user)
    {
        $id = $request->id;
        $data = $this->validate($request,[
            'id'            =>'required|integer',
            'is_admin'      =>'required|integer|in:0,1'
        ]);
        $res = $user->where('id',$id)->update($data);
        return $res?renderSuccess('修改会员成功')
            :renderError('修改会员失败,请稍后再试！');
    }
}
