<?php

namespace App\Http\Controllers\Home;

use App\Http\Models\{Article,ArticleComment,Category,Image,Link,Tag,WebConfig};
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Cache;

class IndexController extends BasicController
{
    //社交地址type
    private const SOCIAL_TYPE = 1;
    //footer内容type
    private const FOOTER_TYPE = 2;

    /**
     * 博客首页
     * Date: 2019/1/8 18:49
     * @param Request $request
     * @param Article $article
     * @param Image   $image
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request,Article $article,Image $image)
    {
        $view = $this->countView($request);
        $newArticle = $article->getList();
        $img = $this->rotation($image);
        $rotation = $img['rotation'];
        $photo = $img['photo'];
        return response()->view('home.index', compact('newArticle', 'rotation', 'photo'))
            ->cookie('view_index', $view, $this->getLastMinute());
    }

    //获取今日剩余分钟
    private function getLastMinute()
    {
        return (strtotime(date('Y-m-d').' 23:59:59')-time())/60;
    }


    //博客首页分类
    public function category(Category $category,WebConfig $config): JsonResponse
    {
        $data = Cache::rememberForever(config('blog.category_cache_key'),function () use ($category,$config){
            $list = $category->getTree();
            $social = $config->where('type',self::SOCIAL_TYPE)->get()->toArray();
            $social = $this->handleSocialData($social);
            return compact('list','social');
        });
        return renderSuccess('获取分类数据成功',$data);
    }

    /**
     * 处理数据
     * Date: 2019/1/4 16:06
     * @param array $data
     * @return array
     */
    private function handleSocialData(array $data): array
    {
        $new_data = [];
        foreach ($data as $datum)
        {
            $new_data[$datum['name']] = $datum['val'];
        }
        return $new_data;
    }

    //博客首页轮播图
    public function rotation(Image $image): array
    {
        $data = Cache::rememberForever(config('blog.rotation_cache_key'),function () use ($image){
            $rotation = $image->rotationImage();
            $photo = $image->photoImage();
            return compact('rotation','photo');
        });
        return $data;
    }

    //获取文章列表
    public function getArticleList(Request $request,Article $article): JsonResponse
    {
        $data = $request->all();
        $data['order'] = 'is_top';
        $data['show'] = 1;
        $list = $article->getList($data);
        return $list->isNotEmpty()?renderSuccess('获取文章列表成功',$list)
            :renderError('暂无文章列表数据');
    }

    /**
     * 获取footer内容
     * Date: 2019/1/4 16:33
     * @param WebConfig $config
     * @return JsonResponse
     */
    public function getFooter(WebConfig $config): JsonResponse
    {
        $list = $config->where('type',self::FOOTER_TYPE)->get()->groupBy('name');
        return $list->isNotEmpty()?renderSuccess('获取footer列表成功',$list)
            :renderError('暂无footer列表数据');
    }

    /**
     * 获取侧边栏数据
     * Date: 2019/1/4 17:52
     * @return JsonResponse
     */
    public function getSidebarData(): JsonResponse
    {
        $tag = $this->getTagData();
        $hot_article = $this->getHotArticleData();
        $comment = $this->getCommentData();
        $link = $this->getLinkData();
        return renderSuccess('获取侧边栏数据成功!',compact('tag','hot_article','comment','link'));
    }

    /**
     * 获取标签数据
     * Date: 2019/1/4 17:51
     * @return mixed
     */
    private function getTagData()
    {
        $data = Cache::rememberForever(config('blog.tag_cache_key'),function (){
            $list = Tag::get();
            return $list->isNotEmpty()?$list:null;
        });
        return $data;
    }

    /**
     * 获取热文数据
     * Date: 2019/1/4 17:51
     * @return mixed
     */
    private function getHotArticleData()
    {
        $data = Cache::rememberForever(config('blog.hot_article_cache_key'),function (){
            $list = Article::where('is_show',1)
                ->selectRaw('id,title,created_at,click,
                (select count(*) from article_comment where article_comment.article_id=article.id) as comment_count')
                ->latest('click')
                ->take(10)
                ->get();
            return $list->isNotEmpty()?$list:null;
        });
        return $data;
    }

    /**
     * 获取评论数据
     * Date: 2019/1/4 17:51
     * @return mixed
     */
    private function getCommentData()
    {
        $data = Cache::rememberForever(config('blog.comment_cache_key'),function (){
            $list = ArticleComment::join('user','article_comment.user_id','=','user.id')
                ->select('article_comment.id','article_comment.article_id',
                    'article_comment.content','article_comment.created_at','user.name','user.avatar')
                ->take(10)
                ->latest()
                ->get();
            return $list->isNotEmpty()?$list:null;
        });
        return $data;
    }

    /**
     * 获取友联数据
     * Date: 2019/1/4 17:51
     * @return mixed
     */
    private function getLinkData()
    {
        $data = Cache::rememberForever(config('blog.link_cache_key'),function (){
            $list = Link::where('is_show',1)->latest('order')->latest('created_at')->get();
            return $list->isNotEmpty()?$list:null;
        });
        return $data;
    }

}
