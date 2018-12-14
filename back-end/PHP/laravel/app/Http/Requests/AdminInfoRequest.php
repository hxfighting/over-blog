<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class AdminInfoRequest extends FormRequest
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
            'avatar'    =>'required|url',
            'phone'    =>'required|phone:CN,mobile',
            'email'     =>'required|email'
        ];
    }
}
