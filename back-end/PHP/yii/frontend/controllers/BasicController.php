<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-15
 * Time: 09:47
 */

namespace frontend\controllers;


use frontend\models\Article;
use frontend\models\Category;
use frontend\models\Comment;
use frontend\models\Link;
use frontend\models\Tag;
use frontend\models\WebConfig;
use yii\web\Controller;

class BasicController extends Controller
{
    protected $cache;

    public function init()
    {
        parent::init();
        $this->cache = \Yii::$app->cache;
    }

    public function beforeAction($action)
    {
        $this->getCommonData();
        return parent::beforeAction($action);
    }

    /**
     * 获取博客公共数据
     * Date: 2019-03-19 09:43
     */
    private function getCommonData()
    {
        $this->getCategoryData();
        $this->getSocialData();
        $this->getTagData();
        $this->getHotArticleData();
        $this->getCommentData();
        $this->getLinkData();
        $this->getFooterData();
    }

    /**
     * 获取博客首页导航数据
     * Date: 2019-03-19 09:42
     */
    private function getCategoryData()
    {
        $data = $this->cache->getOrSet(\Yii::$app->params['category_cache_key'], function () {
            return Category::find()
                ->with('children')
                ->where(['pid' => 0])
                ->asArray()
                ->all();
        });
        \Yii::$app->view->params['dh'] = $data;
    }

    /**
     * 获取博客首页社交数据
     * Date: 2019-03-19 09:43
     */
    private function getSocialData()
    {
        $data = $this->cache->getOrSet(\Yii::$app->params['social_cache_key'], function () {
            return WebConfig::find()->where(['type' => 1])->asArray()->all();
        });
        \Yii::$app->view->params['socialData'] = $data;
    }

    /**
     * 获取博客标签云数据
     * Date: 2019-03-19 09:43
     */
    private function getTagData()
    {
        $data = $this->cache->getOrSet(\Yii::$app->params['tag_cache_key'], function () {
            return Tag::find()->asArray()->all();
        });
        \Yii::$app->view->params['tagCloud'] = $data;
    }

    /**
     * 获取最热文章
     * Date: 2019-03-19 11:24
     */
    private function getHotArticleData()
    {
        $seconds = strtotime(date('Y-m-d 23:59:59'))-time();
        $data = $this->cache->getOrSet(\Yii::$app->params['hot_article_cache_key'], function () {
            return Article::find()
                ->orderBy('click desc,created_at desc')
                ->select('id,title,created_at,click,
                (select count(*) from article_comment where article_comment.article_id=article.id) as comment_count')
                ->asArray()
                ->where(['is_show'=>1])
                ->limit(10)
                ->all();
        },$seconds);
        \Yii::$app->view->params['hotArticle'] = $data;
    }

    /**
     * 获取评论
     * Date: 2019-03-20 13:42
     */
    private function getCommentData()
    {
        $data = $this->cache->getOrSet(\Yii::$app->params['comment_cache_key'], function () {
            return Comment::find()
                ->join('LEFT JOIN','user','user.id=article_comment.user_id')
                ->select('article_comment.*,user.name,user.avatar')
                ->asArray()
                ->limit(10)
                ->orderBy('created_at desc')
                ->all();
        });
        \Yii::$app->view->params['comment_t'] = $data;
    }

    /**
     * 获取友联数据
     * Date: 2019-03-20 13:47
     */
    private function getLinkData()
    {
        $data = $this->cache->getOrSet(\Yii::$app->params['link_cache_key'], function () {
            return Link::find()
                ->orderBy('order desc,created_at asc')
                ->where(['is_show'=>1])
                ->asArray()
                ->all();
        });
        \Yii::$app->view->params['friendLink'] = $data;
    }

    /**
     * 获取footer内容
     * Date: 2019-03-25 15:32
     */
    private function getFooterData()
    {
        $data = $this->cache->getOrSet(\Yii::$app->params['footer_cache_key'], function () {
            $data = WebConfig::find()
                ->where(['type'=>2])
                ->asArray()
                ->all();
            return $this->handleFooterData($data);
        });

        \Yii::$app->view->params['footerData'] = $data;
    }

    /**
     * 处理footer内容
     * Date: 2019-03-25 15:32
     * @param $data
     * @return array
     */
    private function handleFooterData($data)
    {
        $da = array_fill_keys(array_unique(array_column($data,'name')),[]);
        foreach ($data as $datum)
        {
            if(isset($da[$datum['name']])){
                $da[$datum['name']][] = $datum;
            }
        }
        return $da;
    }



}
