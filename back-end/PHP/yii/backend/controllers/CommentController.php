<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-08
 * Time: 09:15
 */

namespace backend\controllers;


use app\models\Article;
use app\models\ArticleComment;
use app\models\User;
use backend\job\CommentReplyJob;
use yii\db\Exception;
use yii\db\Query;

class CommentController extends BasicController
{
    public $enableCsrfValidation = false;
    private $comment;

    public function init()
    {
        parent::init();
        $this->comment = new ArticleComment();
    }

    /**
     * 获取评论列表
     * Date: 2019-03-08 09:40
     * @return \yii\web\Response
     */
    public function actionCommentList()
    {
        $data = $this->get();
        $this->comment->scenario = 'commentList';
        $this->comment->attributes = $data;
        if ($this->comment->validate())
        {
            $article = Article::find()->select('id,title')->all();
            $query = ArticleComment::find()
                ->alias('ac')
                ->leftJoin('user as u', "ac.user_id = u.id")
                ->leftJoin('user as r', "ac.reply_id = r.id")
                ->select('ac.*,u.name as username,r.name as reply_name');
            if (isset($data['article_id']) && is_numeric($data['article_id']))
            {
                $query = $query->where(['article_id' => $data['article_id']]);
            }
            $total = (int)($query->count());
            $list = $query
                ->offset(($data['pageNum'] - 1) * $data['pageSize'])
                ->limit($data['pageSize'])
                ->orderBy('created_at DESC')
                ->asArray()
                ->all();
            if (!empty($list))
            {
                $list = $this->handleCommentListData($list);
                $data = compact('list', 'total', 'article');
                return $this->success('获取评论列表成功！', $data);
            }
            return $this->error('暂无评论列表数据！');
        }
        return $this->error(current($this->comment->firstErrors));
    }

    /**
     * 处理列表数据
     * Date: 2019-03-08 10:48
     * @param $data
     * @return mixed
     */
    private function handleCommentListData($data)
    {
        foreach ($data as &$datum)
        {
            $datum['id'] = (int)$datum['id'];
            $datum['pid'] = (int)$datum['pid'];
            $datum['replier'] = $datum['reply_name'] ? ['id' => (int)$datum['reply_id'], 'name' => $datum['reply_name']] : null;
            $datum['user'] = ['id' => (int)$datum['user_id'], 'name' => $datum['username']];
            $datum['article_id'] = (int)$datum['article_id'];
            $datum['created_at'] = date('Y/m/d H:i:s', (int)$datum['created_at']);
            $datum['updated_at'] = date('Y/m/d H:i:s', (int)$datum['updated_at']);
        }
        unset($datum);
        return $data;
    }

    /**
     * 删除评论
     * Date: 2019-03-08 10:59
     * @return \yii\web\Response
     */
    public function actionDelComment()
    {
        $data = $this->post();
        $this->comment->scenario = 'commentDelete';
        $this->comment->attributes = $data;
        if ($this->comment->validate())
        {
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                $this->comment->deleteAll(['id' => $data['id']]);
                $this->comment->deleteAll(['pid' => $data['id']]);
                $tr->commit();
                return $this->success('删除评论成功！');
            } catch (Exception $e)
            {
                $tr->rollBack();
                return $this->error('删除评论失败，请稍后再试！');
            }
        }
        return $this->error(current($this->comment->firstErrors));
    }

    /**
     * 回复评论
     * Date: 2019-03-08 11:39
     * @return \yii\web\Response
     */
    public function actionReply()
    {
        $data = $this->post();
        $this->comment->scenario = 'commentReply';
        $this->comment->attributes = $data;
        if ($this->comment->validate())
        {
            $exist_comment = $this->comment->findOne($data['id']);
            $user_id = $this->getUserId($exist_comment['reply_id']);
            if (!$user_id)
            {
                return $this->error('请先绑定前台管理员！');
            }
            $co_data = [
                'pid' => $exist_comment['pid'] ? $exist_comment['pid'] : $exist_comment['user_id'],
                'content' => $data['reply_content'],
                'user_id' => $user_id,
                'reply_id' => $exist_comment['user_id'],
                'article_id' => $exist_comment['article_id'],
                'created_at' => time(),
                'updated_at' => time()
            ];
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                \Yii::$app->db->createCommand()->insert('article_comment', $co_data)->execute();
                $comment_data = [
                    'article_id' => $exist_comment['article_id'],
                    'user_id' => $user_id,
                    'reply_id' => $exist_comment['user_id'],
                    'reply_content' => $data['reply_content']
                ];
                \Yii::$app->queue->push(new CommentReplyJob($comment_data));
                $tr->commit();
                return $this->success('回复评论成功！');
            } catch (Exception $e)
            {
                $tr->rollBack();
                return $this->error('回复评论失败！');
            }
        }
        return $this->error(current($this->comment->firstErrors));
    }

    /**
     * 获取用户ID
     * Date: 2019-03-08 11:15
     * @param $reply_id
     * @return bool|mixed
     */
    private function getUserId($reply_id)
    {
        if (!$reply_id)
        {
            $reply_user = User::find()
                ->where(['is_admin' => 1])
                ->orderBy('created_at DESC')
                ->one();
            if (empty($reply_user))
            {
                return false;
            }
            $reply_id = $reply_user['id'];
        }
        return $reply_id;
    }
}
