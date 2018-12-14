<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class AdminLoginRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'name'      =>'required|min:2|max:20',
            'password'  =>'required|min:6|max:20',
//            'captcha'   =>'required|size:4',
//            'key'       =>'required'
        ];
    }

    public function messages()
    {
        return [
            'name.required'     =>'请输入用户名',
            'name.min'          =>'用户名或密码错误',
            'name.max'          =>'用户名或密码错误',
            'password.required' =>'请输入密码',
            'password.min'      =>'用户名或密码错误',
            'password.max'      =>'用户名或密码错误'
        ];
    }
}
