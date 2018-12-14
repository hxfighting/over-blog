<?php

namespace App\Http\Controllers;

use App\Http\Models\Tag;
use Illuminate\Http\Request;

class TagController extends Controller
{
    //获取标签列表数据
    public function getList(Request $request,Tag $tag)
    {
        $list = $tag->filter($request->all())->get();
        $total = $tag->count();
        if($list->isNotEmpty()){
            $data = compact('total','list');
            return renderSuccess('获取标签列表成功',$data);
        }
        return renderError('暂无标签列表数据');
    }

    //删除标签
    public function destroy(Request $request,Tag $tag)
    {
        $data = $this->validate($request,['id'=>'required|integer']);
        $res = $tag->destroy($data['id']);
        return $res?renderSuccess('删除标签成功')
            :renderError('删除标签失败,请稍后再试！');
    }

    //添加标签
    public function store(Request $request,Tag $tag)
    {
        $data = $this->validate($request,['name'=>'required|min:2|max:20|unique:tag,name']);
        $res = $tag->create($data);
        return $res?renderSuccess('添加标签成功')
            :renderError('添加标签失败,请稍后再试！');
    }

    //修改标签
    public function update(Request $request,Tag $tag)
    {
        $id = $request->id;
        $data = $this->validate($request,['id'=>'required|integer','name'=>'required|min:2|max:20|unique:tag,name,'.$id]);
        $res = $tag->where('id',$id)->update($data);
        return $res?renderSuccess('修改标签成功')
            :renderError('修改标签失败,请稍后再试！');
    }
}
