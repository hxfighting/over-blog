<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-01-28
 * Time: 14:56
 */

namespace App\Http\Controllers\Home;


use App\Http\Models\Contact;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class ContactController extends BasicController
{
    /**
     * 联系我
     * Date: 2019-01-28 15:04
     * @return \Illuminate\Contracts\View\Factory|\Illuminate\View\View
     */
    public function index()
    {
        return view('home.contact');
    }

    /**
     * 留言
     * Date: 2019-01-29 10:35
     * @param Request $request
     * @param Contact $contact
     * @return JsonResponse
     * @throws \Illuminate\Validation\ValidationException
     */
    public function store(Request $request,Contact $contact): JsonResponse
    {
        $data = $this->validate($request, [
            'name' => 'required|between:2,20',
            'email' => 'required|email',
            'content' => 'required|between:2,255'
        ], [
            'name.required' => '请输入姓名！',
            'name.between' => '姓名在2到20个字符之间！',
            'email.required' => '请输入邮箱',
            'email.email' => '请输入正确的邮箱',
            'content.required' => '请输入留言内容！',
            'content.between' => '留言内容在2到255个字符之间！'
        ]);
        $contact->name = $data['name'];
        $contact->email = $data['email'];
        $contact->content = $data['content'];
        $res = $contact->save();
        return $res?renderSuccess('留言成功！'):renderError('留言失败，请稍后再试！');
    }
}
