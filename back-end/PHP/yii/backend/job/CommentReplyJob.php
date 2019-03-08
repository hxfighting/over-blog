<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-08
 * Time: 11:18
 */

namespace backend\job;

use console\models\Article;
use console\models\User;
use yii\base\BaseObject;
use yii\queue\JobInterface;

class CommentReplyJob extends BaseObject implements JobInterface
{
    public $user_id;

    public $reply_id;

    public $reply_content;

    public $article_id;

    public function execute($queue)
    {
        try
        {
            $da = [
                $this->user_id,
                $this->reply_content,
                $this->reply_id,
                $this->article_id
            ];
            \Yii::$app->log->getLogger()->log($da,1);
            $subject = \Yii::$app->params['blog_name'] . '评论回复';
            $reply_user = User::find()->where(['id' => $this->reply_id])->one();
            \Yii::$app->log->getLogger()->log($reply_user,1);
            $comment_user = User::find()->where(['id' => $this->user_id])->one();
            \Yii::$app->log->getLogger()->log($comment_user,1);
            $article = Article::find()->where(['id' => $this->article_id])->one();
            \Yii::$app->log->getLogger()->log($article,1);
            $url = \Yii::$app->params['blog_index_url'] . '/article/' . $this->article_id;
            $data = [
                'reply_name' => $reply_user['name'],
                'comment_name' => $comment_user['name'],
                'title' => $article['title'],
                'url' => $url,
                'reply_content' => $this->reply_content
            ];
            \Yii::$app->mailer->compose('CommentReplyMail', $data)
                ->setTo($reply_user['email'])
                ->setSubject($subject)
                ->send();
        } catch (\Exception $e)
        {
            \Yii::$app->log->getLogger()->log($e->getMessage(),1);
        }
    }
}
