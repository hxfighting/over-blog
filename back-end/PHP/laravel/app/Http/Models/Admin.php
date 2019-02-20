<?php

namespace App\Http\Models;

use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Facades\Hash;
use Mews\Captcha\Facades\Captcha;
use Tymon\JWTAuth\Contracts\JWTSubject;

class Admin extends Authenticatable implements JWTSubject
{
    protected $table = 'admin';
    protected $dateFormat = 'U';
    protected $guarded = [];

    //多态关联
    public function images()
    {
        return $this->morphMany(Image::class, 'image');
    }

    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    /**
     * Return a key value array, containing any custom claims to be added to the JWT.
     *
     * @return array
     */
    public function getJWTCustomClaims()
    {
        return [];
    }

    //登录
    public function login(array $data): JsonResponse
    {
        if(Captcha::check_api($data['captcha'],$data['key'])){
            $user = $this->where('name', $data['name'])->first();
            if ($user && password_verify($data['password'], $user->password))
            {
                $token = auth('admin')->login($user);
                $token = 'Bearer ' . $token;
                return renderSuccess('登录成功', $token);
            }
            return renderError('登录失败,用户名或密码错误');
        }
        return renderError('验证码错误');
    }

    //获取管理员信息
    public function getInfo(): JsonResponse
    {
        $user = auth('admin')->user();
        $index = strripos($user->avatar,'/');
        $user_data = [
            'avatar'    => $user->avatar,
            'name'      => $user->name,
            'user_id'   => $user->id,
            'access'    => ['super_admin'],
            'email'     => $user->email,
            'avatarName'=> substr($user->avatar,$index+1),
            'mobile'    => $user->mobile
        ];
        return renderSuccess('获取用户信息成功', $user_data);
    }

    //修改个人信息
    public function updateInfo(array $data): JsonResponse
    {
        $res = auth('admin')->user()->update($data);
        return $res?renderSuccess('修改个人信息成功')
            :renderError('修改个人信息失败，请稍后再试');
    }

    //退出登录
    public function logout(): JsonResponse
    {
        auth('admin')->logout(true,true);
        return renderSuccess('退出登录成功');
    }

    //修改密码
    public function resetPassword(string $password): JsonResponse
    {
        $password = Hash::make($password);
        $res = auth('admin')->user()->update(['password'=>$password]);
        return $res?renderSuccess('修改密码成功')
            :renderError('修改密码失败，请稍后再试');
    }
}
