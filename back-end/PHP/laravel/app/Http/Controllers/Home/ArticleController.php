<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\Article;
use App\Http\Models\ArticleComment;
use Illuminate\Http\Request;

class ArticleController extends BasicController
{
    /**
     * 文章详情页
     * Date: 2019/1/16 10:53
     * @param Request $request
     * @param         $id
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request, $id)
    {
        //得到单一文章详情
        $single = Article::with('tags:tag.id,name', 'category:category.id,category.title')
            ->withCount('comments')
            ->findorfail($id);
        //文章查看次数,一个用户一天只增加一次
        $view = $this->countView($request);
        $prev = $this->getPreArticle($id);
        $next = $this->getNextArticle($id);
        $randArticle = $this->getRandomArticle();
        $comment = $this->handleCommentData($id);
        return response()->view('home.single', compact('single', 'randArticle', 'prev',
            'next','comment'))->cookie('view_article', $view, 1440);
    }

    /**
     * 获取上一篇文章
     * Date: 2019/1/11 15:49
     * @param int $id
     * @return mixed
     */
    private function getPreArticle(int $id)
    {
        return Article::where('id', '<', $id)
            ->show(1)
            ->latest('id')
            ->select('id', 'title')
            ->first();
    }

    /**
     * 获取下一篇文章
     * Date: 2019/1/11 15:50
     * @param int $id
     * @return mixed
     */
    private function getNextArticle(int $id)
    {
        return Article::where('id', '>', $id)
            ->show(1)
            ->oldest('id')
            ->select('id', 'title')
            ->first();
    }

    /**
     * 获取随机文章
     * Date: 2019/1/11 15:51
     * @return mixed
     */
    private function getRandomArticle()
    {
        return Article::inRandomOrder()
            ->show(1)
            ->take(10)
            ->select('title', 'thumb', 'id')
            ->get();
    }

    /**
     * 获取评论数据
     * Date: 2019/1/11 15:48
     * @param int $article_id
     * @return mixed
     */
    private function handleCommentData(int $article_id)
    {
        $where = [
            ['article_comment.article_id',$article_id],
            ['article_comment.pid',0]
        ];
        $query = ArticleComment::with(['children' => function ($q) {
            return $this->buildQuery($q);
        }]);
        $comment = $this->buildQuery($query)
            ->where($where)
            ->get();
        return $comment;
    }

    /**
     * 构造查询
     * Date: 2019/1/10 14:57
     * @param $q
     * @return mixed
     */
    private function buildQuery($q)
    {
        return $q->leftJoin('user as a', 'article_comment.reply_id', '=', 'a.id')
            ->leftJoin('user as b', 'article_comment.user_id', '=', 'b.id')
            ->select('article_comment.*', 'a.avatar as reply_avatar',
                'a.name as reply_name', 'b.avatar as user_avatar', 'b.name as username');
    }

    /**
     * 搜索文章
     * Date: 2019/1/16 10:57
     * @param Request $request
     * @param Article $article
     * @return \Illuminate\Contracts\View\Factory|\Illuminate\View\View
     */
    public function search(Request $request,Article $article)
    {
        $search = $request->input('search',null);
        $list = $article->with('tags:tag.id,name', 'category:category.id,category.title')
            ->when($search,function ($q) use ($search){
            return $q->where('title', 'like', '%' . $search . '%');
        })->latest()->paginate(7);
        return view('home.category', compact('list'));
    }
}
