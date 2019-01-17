<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\Link;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use App\Http\Controllers\Controller;

class LinkController extends Controller
{
    /**
     * 申请友联
     * Date: 2019/1/16 13:18
     * @param Request $request
     * @param Link    $link
     * @return JsonResponse
     * @throws \Illuminate\Validation\ValidationException
     */
    public function store(Request $request,Link $link): JsonResponse
    {
        $data = $this->validate($request, [
            'name' => 'required|between:2,50',
            'url' => 'required|url',
            'description' => 'required|between:2,50'], [
            'name.required' => '请输入友联名称',
            'name.between' => '友联名称在2到50个字符之间',
            'url.required' => '请输入友联地址',
            'url.url' => '友联地址不正确',
            'description.required' => '请输入友联描述',
            'description.between' => '友联描述在2到50个字符之间'
        ]);
        $res = $link->create($data);
        return $res?renderSuccess('申请友联成功!')
            :renderError('申请友联失败,请稍后再试!');
    }
}
