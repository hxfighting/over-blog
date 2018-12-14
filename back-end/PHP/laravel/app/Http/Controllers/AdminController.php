<?php

namespace App\Http\Controllers;

use App\Http\Models\Admin;
use App\Http\Requests\AdminInfoRequest;
use App\Http\Requests\AdminLoginRequest;
use Illuminate\Http\Request;

class AdminController extends Controller
{
    private $admin;

    public function __construct(Admin $admin)
    {
        $this->admin = $admin;
    }

    //管理员登录
    public function login(AdminLoginRequest $request)
    {
        $data = $request->only(['name','password','captcha','key']);
        return $this->admin->login($data);
    }

    //获取管理员信息
    public function getAdminInfo()
    {
        return $this->admin->getInfo();
    }
    
    //修改个人信息
    public function updateInfo(AdminInfoRequest $request)
    {
        $data = $request->only(['avatar','phone','email','name']);
        $data['mobile'] = ltrim(phone($data['phone'], 'CN', 'E164'), '+86');
        unset($data['phone']);
        return $this->admin->updateInfo($data);
    }

    //退出登录
    public function logout()
    {
        return $this->admin->logout();
    }

    //修改密码
    public function resetPassword(Request $request)
    {
        $data = $this->validate($request,['password'=>'required|min:6|max:20|confirmed']);
        return $this->admin->resetPassword($data['password']);
    }
}
