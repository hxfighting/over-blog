<?php

namespace App\Http\Controllers;

use App\Http\Models\Contact;
use Illuminate\Http\Request;

class ContactController extends Controller
{
    //获取留言列表数据
    public function getList(Request $request,Contact $contact)
    {
        $request_data = $request->all();
        $list = $contact->filter($request_data)->get();
        if($list->isNotEmpty()){
            $total = $contact->getContactCount($request_data);
            return renderSuccess('获取留言列表成功',compact('total','list'));
        }
        return renderError('暂无留言列表数据');
    }

    //删除留言
    public function destroy(Request $request,Contact $contact)
    {
        $data = $this->validate($request,['id'=>'required|integer']);
        $res = $contact->destroy($data['id']);
        return $res?renderSuccess('删除留言成功')
            :renderError('删除留言失败,请稍后再试！');
    }

    //留言
    public function store(Request $request,Contact $contact)
    {
        $data = $this->validate($request,[
            'name'           =>'required|min:2|max:20',
            'email'          =>'required|email',
            'content'        =>'required|min:2|max:255',
        ]);
        $contact->name      = $data['name'];
        $contact->email     = $data['email'];
        $contact->content   = $data['content'];
        $res = $contact->save();
        return $res?renderSuccess('留言成功')
            :renderError('留言失败,请稍后再试!');
        //TODO 如果第三方登录了再绑定邮箱到第三方账号
    }

    //回复留言
    public function reply(Request $request,Contact $contact)
    {
        $data = $this->validate($request,[
            'id'             =>'required|integer',
            'reply_content'  =>'required|min:2|max:255',
        ]);
        $exist_contact = $contact->find($data['id']);
        $exist_contact->is_reply = 1;
        $exist_contact->reply_content = $data['reply_content'];
        $res = $exist_contact->save();
        return $res?renderSuccess('回复留言成功')
            :renderError('回复留言失败,请稍后再试');
    }
}
