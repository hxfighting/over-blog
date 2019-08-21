<?php

namespace App\Http\Controllers;

use App\Http\Models\Article;
use App\Http\Models\ArticleComment;
use App\Http\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class CommentController extends Controller
{
    //获取评论列表数据
    public function getList(Request $request, ArticleComment $comment, Article $article)
    {
        $request_data = $request->all();
        $list = $comment
            ->with('user:id,name', 'replier:id,name')
            ->filter($request_data)
            ->get();
        $article = $article->get(['id', 'title']);
        if ($list->isNotEmpty())
        {
            $total = $comment->getCommentCount($request_data);
            return renderSuccess('获取评论列表成功', compact('total', 'list', 'article'));
        }
        return renderError('暂无评论数据', compact('article'));
    }

    //删除评论
    public function destroy(Request $request, ArticleComment $comment)
    {
        $data = $this->validate($request, ['id' => 'required|integer']);
        try
        {
            DB::transaction(function () use ($comment, $data) {
                $exist_comment = $comment->find($data['id']);
                if ($exist_comment->pid == 0)
                {
                    $comment->where('pid', $exist_comment->user_id)->delete();
                }
                $exist_comment->delete();
            });
            return renderSuccess('删除评论成功');
        } catch (\Exception $e)
        {
            return renderError('删除评论失败,请稍后再试！');
        }
    }

    //后台回复评论
    public function reply(Request $request, User $user, ArticleComment $comment)
    {
        $data = $this->validate($request,
            [
                'id'            => 'required|integer|exists:article_comment,id',
                'reply_content' => 'required|min:2|max:255'
            ]);
        return $comment->handleReplyComment($user, $data);
    }

}
