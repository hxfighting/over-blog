<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-18
 * Time: 18:01
 */

namespace backend\controllers;


use app\models\Article;
use app\models\ArticleComment;
use app\models\Category;
use app\models\Chat;
use app\models\Contact;
use app\models\Link;
use app\models\Tag;
use app\models\User;

class DashboardController extends BasicController
{
    /**
     * 获取统计数据
     * Date: 2019-02-21 15:31
     * @return \yii\web\Response
     */
    public function actionDashboardCount()
    {
        $user = User::find()->count();
        $link = Link::find()->count();
        $article = Article::find()->count();
        $comment = ArticleComment::find()->count();
        $tag = Tag::find()->count();
        $category = Category::find()->count();
        $contact = Contact::find()->count();
        $chat = Chat::find()->count();
        $data = [
            [
                'title' => '用户统计',
                'icon' => 'md-people',
                'count' => (int)$user,
                'color' => '#2d8cf0'
            ],
            [
                'title' => '友联统计',
                'icon' => 'ios-link',
                'count' => (int)$link,
                'color' => '#19be6b'
            ],
            [
                'title' => '文章统计',
                'icon' => 'ios-book',
                'count' => (int)$article,
                'color' => '#ff9900'
            ],
            [
                'title' => '评论统计',
                'icon' => 'ios-chatboxes',
                'count' => (int)$comment,
                'color' => '#ed3f14'
            ],
            [
                'title' => '标签统计',
                'icon' => 'md-pricetags',
                'count' => (int)$tag,
                'color' => '#E46CBB'
            ],
            [
                'title' => '分类统计',
                'icon' => 'md-list',
                'count' => (int)$category,
                'color' => '#9A66E4'
            ],
            [
                'title' => '留言统计',
                'icon' => 'md-mail',
                'count' => (int)$contact,
                'color' => '#FF99CC'
            ],
            [
                'title' => '说说统计',
                'icon' => 'ios-chatbubbles',
                'count' => (int)$chat,
                'color' => '#FFFF00'
            ]
        ];
        return $this->success('获取列表统计成功!', $data);
    }
}
