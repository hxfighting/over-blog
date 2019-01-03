<?php

namespace App\Http\Controllers;

use App\Http\Models\WebConfig;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class ConfigController extends Controller
{
    //社交地址type
    private const SOCIAL_TYPE = 1;
    //footer内容配置
    private const FOOTER_TYPE = 2;


    //获取配置项列表数据
    public function getList(Request $request, WebConfig $config): JsonResponse
    {
        $list = $config->filter($request->all())->get();
        if ($list->isNotEmpty())
        {
            $total = $config->getListCount($request->all());
            return renderSuccess('获取配置项成功!', compact('list', 'total'));
        }
        return renderError('暂无配置项数据!');
    }

    //添加配置项
    public function store(Request $request, WebConfig $config): JsonResponse
    {
        $data = $this->validate($request, [
            'type' => 'required|integer|in:1,2,3',
            'name' => 'required|between:2,100',
            'title' => 'required|between:2,100',
            'val' => 'required|min:1',
        ]);
        if ($data['type'] == self::FOOTER_TYPE)
        {
            $count = $config->where(['name' => $data['name'], 'type' => $data['type']])->count();
            if ($data['name'] == 'copyright' && $count >= 1)
            {
                return renderError('版权信息只能有一条!');
            }
            if ($count >= 4)
            {
                return renderError('footer内容每一项最多添加4个!');
            }
        }
        $exist = $config->firstOrNew(['name' => $data['name'], 'type' => $data['type']]);
        if (!$exist->id)
        {
            $exist->title = $data['title'];
            $exist->val = $data['val'];
            $res = $exist->save();
            return $res ? renderSuccess('添加配置项成功!')
                : renderError('添加配置项失败,请稍后再试!');
        } else
        {
            return renderError('已添加该配置,请勿重复添加!');
        }
    }

    /**
     * 修改配置
     * Date: 2019/1/3 17:16
     * @param Request   $request
     * @param WebConfig $config
     * @return JsonResponse
     * @throws \Illuminate\Validation\ValidationException
     */
    public function update(Request $request, WebConfig $config): JsonResponse
    {
        $data = $this->validate($request, [
            'id' => 'required|integer',
            'type' => 'required|integer|in:1,2,3',
            'name' => 'required|between:2,100',
            'title' => 'required|between:2,100',
            'val' => 'required|min:1',
        ]);
        if ($data['type'] == self::FOOTER_TYPE)
        {
            $count = $config->where(['name' => $data['name'], 'type' => $data['type']])->count();
            if ($data['name'] == 'copyright' && $count >= 1)
            {
                return renderError('版权信息只能有一条!');
            }
            if ($count >= 4)
            {
                return renderError('footer内容每一项最多添加4个!');
            }
        }
        $exist = $config->firstOrNew(['name' => $data['name'], 'type' => $data['type']]);
        if ($exist->id && $exist->id != $data['id'])
        {
            return renderError('已添加该配置,请勿重复添加!');
        }
        $exist->id = $data['id'];
        $exist->title = $data['title'];
        $exist->val = $data['val'];
        $res = $exist->save();
        return $res ? renderSuccess('修改配置项成功!')
            : renderError('修改配置项失败,请稍后再试!');
    }

    /**
     * 删除配置
     * Date: 2019/1/3 17:25
     * @param Request   $request
     * @param WebConfig $config
     * @return JsonResponse
     * @throws \Illuminate\Validation\ValidationException
     */
    public function delete(Request $request, WebConfig $config): JsonResponse
    {
        $data = $this->validate($request, [
            'id' => 'required|integer'
        ]);
        $res = $config->where('id', $data['id'])->delete();
        return $res ? renderSuccess('删除配置成功!')
            : renderError('删除配置失败,请稍后再试!');
    }
}
