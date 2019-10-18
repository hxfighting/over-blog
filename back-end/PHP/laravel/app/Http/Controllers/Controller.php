<?php

namespace App\Http\Controllers;

use Gregwar\Captcha\CaptchaBuilder;
use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Http\Request;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Support\Str;

class Controller extends BaseController
{
    use AuthorizesRequests, DispatchesJobs, ValidatesRequests;


    //获取验证码
    public function getCaptcha(CaptchaBuilder $builder)
    {
        $builder->build();
        $code = $builder->getPhrase();
        $key = md5(Str::random(32) . uniqid() . session_id() . microtime(true));
        $this->cacheCaptchaCode($code, $key);
        $img = $builder->inline();
        return renderSuccess('获取验证码成功！', compact('img', 'key'));
    }

    /**
     * 缓存验证码code
     * Date: 2019-04-08 14:18
     * @param string $code
     */
    private function cacheCaptchaCode(string $code, string $key)
    {
        app('cache')->put($key, $code, 60);
    }

    /**
     * 刷新token
     * Date: 2019/10/18 10:44 上午
     * @param Request $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function refreshToken(Request $request)
    {
        try
        {
            $newToken = auth('admin')->refresh(false, true);
            $token = 'Bearer ' . $newToken;
            $expire = time() + auth()->factory()->getTTL() * 60;
            return renderSuccess('获取token成功', compact('token', 'expire'));
        } catch (\Exception $e)
        {
            return renderError('非法token');
        }
    }
}
