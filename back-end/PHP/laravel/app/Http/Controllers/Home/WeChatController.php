<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Cache;
use Illuminate\Support\Str;

class WeChatController extends BasicController
{
    /**
     * 获取小程序scene
     * Date: 2019/1/16 13:37
     * @return JsonResponse
     */
    public function getScene(): JsonResponse
    {
        $scene = md5(Str::uuid()->getHex());
        return renderSuccess('获取scene成功', $scene);
    }

    /**
     * 获取小程序码
     * Date: 2019/1/16 13:52
     * @param $scene
     * @return mixed
     */
    public function getQrCode($scene)
    {
        $image = app('miniProgram')->app_code->getUnlimit($scene, ['width' => 200]);
        session(['target_url' => $_SERVER['HTTP_REFERER']]);
        return $image;
    }

    /**
     * 获取登录状态
     * Date: 2019/1/16 14:24
     * @param Request $request
     * @return JsonResponse
     */
    public function getLoginResult(Request $request)
    {
        if ($request->ajax())
        {
            $scene = $request->input('scene', null);
            if (is_null($scene))
            {
                return renderError('参数缺失，请重试');
            }
            $user_info = Cache::get($scene);
            if ($user_info)
            {
                session($user_info);
                auth('web')->loginUsingId($user_info['user']['id']);
                return renderSuccess('登录成功');
            }
        } else
        {
            return renderError('非法请求方式', null,403);
        }
    }

    /**
     * 微信小程序登录
     * Date: 2019/1/16 14:24
     * @param Request $request
     * @param User    $user
     * @return JsonResponse
     * @throws \Illuminate\Validation\ValidationException
     */
    public function login(Request $request, User $user): JsonResponse
    {
        $data = $this->validate($request, [
            'code' => 'required',
            'avatar' => 'required',
            'name' => 'required|max:100',
            'sence' => 'required'
        ]);
        $user_info = app('miniProgram')->auth->session($data['code']);
        if (!empty($user_info))
        {
            $exist_user = $user->firstOrNew(['openid' => $user_info['openid']]);
            $exist_user->access_token = $user_info['session_key'];
            $exist_user->avatar = $data['avatar'];
            $exist_user->name = $data['name'];
            $exist_user->last_login_ip = getIP();
            $exist_user->login_times = isset($exist_user->login_times) ? $exist_user->login_times + 1 : 1;
            $exist_user->save();
            $sessionData = [
                'user' => [
                    'id' => $exist_user->id,
                    'name' => $data['name'],
                    'is_admin' => $exist_user->is_admin
                ]
            ];
            Cache::put($data['sence'], $sessionData, 60 * 60 * 24);
            return renderSuccess('登录成功');
        } else
        {
            return renderError('从微信获取用户信息失败', null, 500);
        }
    }
}
