<?php

namespace App\Http\Controllers;

use App\Http\Models\Image;
use App\Http\Models\RotationImage;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class RotationController extends Controller
{
    //获取轮播图列表数据
    public function getList(Request $request,Image $image)
    {
        $list = $image->rotationImage();
        return $list->isNotEmpty()?renderSuccess('获取轮播图列表成功',$list)
            :renderError('暂无轮播图列表数据');
    }

    //删除轮播图
    public function destroy(Request $request,RotationImage $image)
    {
        $data = $this->validate($request,['id'=>'required|integer']);
        try
        {
            DB::transaction(function () use ($image, $data) {
                $where = [
                    'image_type'=>$this->image_type,
                    'image_id'  =>$data['id']
                ];
                Image::where($where)->delete();
                $image->destroy($data['id']);
            });
            return renderSuccess('删除轮播图成功');
        } catch (\Exception $e)
        {
            return renderError('删除轮播图失败,请稍后再试！');
        }
    }

    //添加轮播图
    public function store(Request $request,RotationImage $image)
    {
        $data = $this->validate($request,[
            'words'=>'nullable|min:2|max:60',
            'image_url'=>'required|url'
        ]);
        try
        {
            DB::transaction(function () use ($image, $data) {
                $res = $image->create(['words' => $data['words']]);
                $res->images()->create(['image_url' => $data['image_url']]);
            });
            return renderSuccess('添加轮播图成功');
        } catch (\Exception $e)
        {
            return renderError('添加轮播图失败,请稍后再试！');
        }
    }

    //修改轮播图
    public function update(Request $request,RotationImage $image)
    {
        $data = $this->validate($request,[
            'id'=>'required|integer',
            'words'=>'nullable|min:2|max:60',
            'image_url'=>'required|url'
        ]);
        try
        {
            DB::transaction(function () use ($image, $data) {
                $image->where('id',$data['id'])->update(['words' => $data['words']]);
                $where = [
                    'image_type'=>$this->image_type,
                    'image_id'  =>$data['id']
                ];
                Image::where($where)->update(['image_url'=>$data['image_url']]);
            });
            return renderSuccess('修改轮播图成功');
        } catch (\Exception $e)
        {
            return renderError('修改轮播图失败,请稍后再试！');
        }
    }
}
