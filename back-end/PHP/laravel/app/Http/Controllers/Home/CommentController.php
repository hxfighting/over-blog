<?php

namespace App\Http\Controllers\Home;

use App\Events\NotifyUserEvent;
use App\Http\Controllers\Controller;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class CommentController extends Controller
{
    /**
     * 评论
     * Date: 2019-01-29 11:46
     * @param Request        $request
     * @return JsonResponse
     * @throws \Illuminate\Validation\ValidationException
     * @throws \Throwable
     */
    public function store(Request $request): JsonResponse
    {
        $data = $this->validate($request, [
            'email' => 'required|email',
            'content' => 'required|min:2',
            'user_id' => 'required|integer|min:1',
            'reply_id' => 'nullable|integer|min:0',
            'article_id' => 'required|integer|min:1',
            'pid' => 'nullable|integer|min:0'
        ]);
        try
        {
            DB::transaction(function () use ($data) {
                $email = $data['email'];
                $data['created_at'] = time();
                $data['updated_at'] = time();
                unset($data['email']);
                \db('article_comment')->insert($data);
                \db('user')->where('id', $data['user_id'])->update(['email' => $email]);
                session(['user.email'=>$email]);
                $this->sendCommentEmail($data,$email);
            });
            return renderSuccess('评论成功！');
        } catch (\Exception $e)
        {
            return renderError('评论失败，请稍后再试！');
        }
    }

    /**
     * 发送邮件通知
     * Date: 2019-01-29 13:33
     * @param array  $data
     * @param string $email
     */
    private function sendCommentEmail(array $data,string $email): void
    {
        event(new NotifyUserEvent($data,$email));
    }
}
