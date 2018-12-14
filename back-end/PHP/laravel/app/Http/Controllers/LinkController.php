<?php

namespace App\Http\Controllers;

use App\Http\Models\Link;
use Illuminate\Http\Request;

class LinkController extends Controller
{
    //获取友联列表数据
    public function getList(Request $request,Link $link)
    {
        $request_data = $request->all();
        $list = $link->filter($request_data)->get();
        if($list->isNotEmpty()){
            $total = $link->getLinkCount($request_data);
            return renderSuccess('获取友联列表成功',compact('total','list'));
        }
        return renderError('暂无友联列表数据');
    }

    //删除友联
    public function destroy(Request $request,Link $link)
    {
        $data = $this->validate($request,['id'=>'required|integer']);
        $res = $link->destroy($data['id']);
        return $res?renderSuccess('删除友联成功')
            :renderError('删除友联失败,请稍后再试！');
    }

    //添加友联
    public function store(Request $request,Link $link)
    {
        $data = $this->validate($request,[
            'url'           =>'required|url|unique:link,url',
            'name'          =>'required|min:2|max:30',
            'description'   =>'required|min:2|max:50',
            'order'         =>'nullable|integer|min:1',
            'is_show'       =>'nullable|integer|in:0,1'
        ]);
        $res = $link->create($data);
        return $res?renderSuccess('添加友联成功')
            :renderError('添加友联失败,请稍后再试!');
    }

    //修改友联
    public function update(Request $request,Link $link)
    {
        $id = $request->id;
        $data = $this->validate($request,[
            'id'            =>'required|integer',
            'url'           =>'required|url|unique:link,url,'.$id,
            'name'          =>'required|min:2|max:30',
            'description'   =>'required|min:2|max:50',
            'order'         =>'nullable|integer|min:1',
            'is_show'       =>'nullable|integer|in:0,1'
        ]);
        $res = $link->where('id',$id)->update($data);
        return $res?renderSuccess('修改友联成功')
            :renderError('修改友联失败,请稍后再试！');
    }
}
