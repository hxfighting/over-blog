<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class ArticleRequest extends FormRequest
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
        $data = [
            'category_id'   =>'required|integer|min:1',
            'title'         =>'required|min:2|max:100',
            'author'        =>'required|min:2|max:20',
            'keywords'      =>'required|min:2|max:255',
            'tags'          =>'required|array',
            'tags.*'        =>'required|integer',
            'description'   =>'nullable|max:255',
            'thumb'         =>'required|url',
            'content_html'  =>'required|min:2',
            'content_md'    =>'required|min:2',
            'is_show'       =>'required|integer|in:0,1',
            'is_top'        =>'required|integer|in:0,1',
            'is_original'   =>'required|integer|in:0,1'
        ];
        if(request()->isMethod('PUT')){
            $data['id'] = 'required|integer';
        }
        return $data;
    }
}
