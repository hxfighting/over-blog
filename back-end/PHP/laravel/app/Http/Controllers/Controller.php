<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Routing\Controller as BaseController;
use Mews\Captcha\Facades\Captcha;

class Controller extends BaseController
{
    use AuthorizesRequests, DispatchesJobs, ValidatesRequests;


    //获取验证码
    public function getCaptcha()
    {
        $captcha = Captcha::create('flat',true);
        return renderSuccess('获取验证码成功',$captcha);
    }
}
