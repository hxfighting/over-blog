<?php

namespace App\Http\Middleware;

use Closure;
use Tymon\JWTAuth\Exceptions\JWTException;
use Tymon\JWTAuth\Exceptions\TokenExpiredException;
use Tymon\JWTAuth\Http\Middleware\BaseMiddleware;
class AdminLoginCheckMiddleware extends BaseMiddleware
{
    /**
     * @param         $request
     * @param Closure $next
     * @return \Illuminate\Http\JsonResponse|\Illuminate\Http\Response|mixed
     * @throws JWTException
     */
    public function handle($request, Closure $next)
    {
        // 检查此次请求中是否带有 token，如果没有则抛出异常。
        $this->checkForToken($request);
        // 使用 try 包裹，以捕捉 token 过期所抛出的 TokenExpiredException  异常
        try
        {// 检测用户的登录状态，如果正常则通过
            if (auth('admin')->check())
            {
                return $next($request);
            } else
            {
                // 刷新用户的 token
                $token = auth('admin')->refresh(true, true);
                // 使用一次性登录以保证此次请求的成功
                $user_id = $this->auth->manager()->getPayloadFactory()->buildClaimsCollection()->toPlainArray()['sub'];
                auth('admin')->onceUsingId($user_id);
            }
        } catch (JWTException $e)
        {
            return renderError('登录状态已失效,请重新登录',null,401);
        }
        // 在响应头中返回新的 token
        return $this->setAuthenticationHeader($next($request), $token);
    }
}
