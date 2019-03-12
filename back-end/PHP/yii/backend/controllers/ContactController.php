<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-05
 * Time: 09:39
 */

namespace backend\controllers;


use app\models\Contact;
use backend\exception\ValidateException;
use backend\job\ReplyMailJob;
use yii\db\Exception;

class ContactController extends BasicController
{
    public $enableCsrfValidation = false;
    private $contact;

    public function init()
    {
        parent::init();
        $this->contact = new Contact();
    }

    /**
     * 获取留言列表
     * Date: 2019-03-05 09:44
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionContactList()
    {
        $this->basicValidate($this->config, 'contactList');
        $data = $this->request_data;
        $query = $this->contact->find();
        if (isset($data['search']) && !empty($data['search']))
        {
            $query = $query->where(['like', 'name', $data['search']])
                ->orWhere(['like', 'email', $data['search']]);
        }
        $total = $query->count();
        $list = $query
            ->offset(($data['pageNum'] - 1) * $data['pageSize'])
            ->limit($data['pageSize'])
            ->orderBy('created_at DESC')
            ->all();
        $data = compact('list', 'total');
        return !empty($list) ? $this->success('获取留言列表成功！', $data)
            : $this->error('暂无留言列表数据！');
    }

    /**
     * 回复留言
     * Date: 2019-03-05 10:17
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionReply()
    {
        $this->basicValidate($this->config, 'contactReply');
        $data = $this->request_data;
        $tr = \Yii::$app->db->beginTransaction();
        try
        {
            $exist_contact = $this->contact->findOne($data['id']);
            $exist_contact->is_reply = 1;
            $exist_contact->reply_content = $data['reply_content'];
            $exist_contact->replied_at = time();
            $exist_contact->save(false, ['is_reply', 'reply_content', 'updated_at', 'replied_at']);
            $reply_data = [
                'email' => $exist_contact->email,
                'reply_content' => $data['reply_content'],
                'name' => $exist_contact->name
            ];
            \Yii::$app->queue->push(new ReplyMailJob($reply_data));
            $tr->commit();
            return $this->success('回复留言成功！');
        } catch (Exception $e)
        {
            $tr->rollBack();
            return $this->error('回复留言失败，请稍后再试！');
        }
    }

    /**
     * 删除留言
     * Date: 2019-03-06 10:43
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionDelContact()
    {
        $this->basicValidate($this->config, 'delContact');
        $data = $this->request_data;
        $res = $this->contact->deleteAll(['id' => $data['id']]);
        return $res ? $this->success('删除留言成功！')
            : $this->error('删除留言失败，请稍后再试！');
    }

}
