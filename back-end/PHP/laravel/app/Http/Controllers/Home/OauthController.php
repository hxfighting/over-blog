<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\User;
use Illuminate\Http\Request;

class OauthController extends BasicController
{
    /**
     * 三方登录授权
     * Date: 2019/1/16 16:01
     * @param $service
     * @return mixed
     */
    public function oauth($service)
    {
        session(['target_url' => $_SERVER['HTTP_REFERER']]);
        return app('socialite')->driver($service)->redirect();
    }

    /**
     * 授权回调
     * Date: 2019/1/16 16:10
     * @param Request $request
     * @param         $service
     * @param User    $user
     * @return \Illuminate\Http\RedirectResponse|\Illuminate\Routing\Redirector
     */
    public function callback(Request $request, $service, User $user)
    {
        $type = [
            'weibo' => 3,
            'wechat' => 2,
            'qq' => 1
        ];
        $oauth_user = app('socialite')->driver($service)->user();
        $data = [
            'openid' => $oauth_user->getId(),
            'type' => $type[$service],
        ];
        $exist_user = $user->firstOrNew($data);
        $exist_user->access_token = $oauth_user->getAccessToken();
        $exist_user->name = $oauth_user->getName();
        $exist_user->last_login_ip = $request->getClientIp();
        $exist_user->login_times = isset($exist_user->login_times) ? $exist_user->login_times + 1 : 1;
        $exist_user->avatar = str_replace('http', 'https', $oauth_user->getAvatar());
        $exist_user->save();
        $sessionData = [
            'user' => [
                'id' => $exist_user->id,
                'name' => $oauth_user->getName(),
                'is_admin' => $exist_user->is_admin,
                'email' => $exist_user->email ?? null
            ]
        ];
        session($sessionData);
        auth('web')->loginUsingId($exist_user->id);
        return redirect(session('target_url', url('/')));
    }

    /**
     * 退出登录
     * Date: 2019/1/16 16:11
     * @param Request $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function logout(Request $request)
    {
        auth('web')->logout();
        $request->session()->forget('user');
        return renderSuccess('退出成功！');
    }
}
