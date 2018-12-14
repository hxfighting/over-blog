<?php

namespace App\Http\Controllers;

use App\Http\Models\WebError;
use Illuminate\Http\Request;

class WebErrorController extends Controller
{
    //记录后台的接口错误日志
    public function logError(Request $request,WebError $error)
    {
        $data = $this->validate($request,
            [
                'code'=>'required|numeric',
                'mes'   =>'required|max:255',
                'type'  =>'required|max:255',
                'url'   =>'required|url'
            ]);
        $res = $error->create($data);
        return $res?renderSuccess('记录错误日志成功')
            :renderError('记录错误日志失败');
    }

    //获取错误日志列表
    public function getErrorList(Request $request,WebError $error)
    {
        $list = $error->filter($request->all())->get();
        $total = $error->count();
        if($list->isNotEmpty()){
            $data = compact('total','list');
            return renderSuccess('获取错误日志列表数据成功',$data);
        }
        return renderError('暂无错误日志数据');
    }

    //删除错误日志
    public function destroy(Request $request,WebError $error)
    {
        $data = $this->validate($request,['ids'=>'required|array','ids.*'=>'required|integer']);
        $res = $error->destroy($data['ids']);
        return $res?renderSuccess('删除错误日志成功')
            :renderError('删除错误日志失败,请稍后再试');
    }
}
