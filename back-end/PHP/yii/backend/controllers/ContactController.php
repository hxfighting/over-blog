<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-05
 * Time: 09:39
 */

namespace backend\controllers;


use app\models\Contact;
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
     */
    public function actionContactList()
    {
        $data = $this->get();
        $this->contact->scenario = 'contactList';
        $this->contact->attributes = $data;
        if($this->contact->validate()){
            $query = $this->contact->find();
            if(isset($data['search']) && !empty($data['search'])){
                $query = $query->where(['like','name',$data['search']])
                    ->orWhere(['like','email',$data['search']]);
            }
            $total = $query->count();
            $list = $query
                ->offset(($data['pageNum'] - 1) * $data['pageSize'])
                ->limit($data['pageSize'])
                ->orderBy('created_at DESC')
                ->all();
            $data = compact('list','total');
            return !empty($list)?$this->success('获取留言列表成功！',$data)
                :$this->error('暂无留言列表数据！');
        }
        return $this->error(current($this->contact->firstErrors));
    }

    /**
     * 回复留言
     * Date: 2019-03-05 10:17
     * @return \yii\web\Response
     */
    public function actionReply()
    {
        $data = $this->post();
        $this->contact->scenario = 'contactReply';
        $this->contact->attributes = $data;
        if($this->contact->validate()){
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                $exist_contact = $this->contact->findOne($data['id']);
                $exist_contact->is_reply = 1;
                $exist_contact->reply_content = $data['reply_content'];
                $exist_contact->replied_at = time();
                $exist_contact->save(false, ['is_reply', 'reply_content', 'updated_at','replied_at']);
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
        return $this->error(current($this->contact->firstErrors));
    }

    /**
     * 删除留言
     * Date: 2019-03-06 10:43
     * @return \yii\web\Response
     */
    public function actionDelContact()
    {
        $data = $this->post();
        $this->contact->scenario = 'delContact';
        $this->contact->attributes = $data;
        if($this->contact->validate()){
            $res = $this->contact->deleteAll(['id'=>$data['id']]);
            return $res?$this->success('删除留言成功！')
                :$this->error('删除留言失败，请稍后再试！');
        }
        return $this->error(current($this->contact->firstErrors));
    }

}
