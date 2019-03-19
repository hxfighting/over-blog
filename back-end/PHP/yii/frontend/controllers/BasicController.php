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


}
