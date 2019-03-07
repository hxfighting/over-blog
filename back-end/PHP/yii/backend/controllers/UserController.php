<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-07
 * Time: 10:17
 */

namespace backend\controllers;


use app\models\User;

class UserController extends BasicController
{
    public $enableCsrfValidation = false;
    private $user;

    public function init()
    {
        parent::init();
        $this->user = new User();
    }

    /**
     * 获取用户列表
     * Date: 2019-03-07 13:38
     * @return \yii\web\Response
     */
    public function actionUserList()
    {
        $data = $this->get();
        $this->user->scenario = 'userList';
        $this->user->attributes = $data;
        if($this->user->validate()){
            $query = $this->user->find();
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
            return !empty($list)?$this->success('获取友联列表成功！',$data)
                :$this->error('暂无友联列表数据！');
        }
        return $this->error(current($this->user->firstErrors));
    }

    /**
     * 修改会员信息
     * Date: 2019-03-07 13:41
     * @return \yii\web\Response
     */
    public function actionUserUpdate()
    {
        $data = $this->post();
        $this->user->scenario = 'userUpdate';
        $this->user->attributes = $data;
        if($this->user->validate()){
            $exist_user = $this->user->findOne($data['id']);
            $exist_user->is_admin = $data['is_admin'];
            $res = $exist_user->save(false,['is_admin','updated_at']);
            return $res?$this->success('修改会员信息成功！')
                :$this->error('修改会员信息失败，请稍后再试！');
        }
        return $this->error(current($this->user->firstErrors));
    }

    /**
     * 删除会员
     * Date: 2019-03-07 13:41
     * @return \yii\web\Response
     */
    public function actionDelUser()
    {
        $data = $this->post();
        $this->user->scenario = 'delUser';
        $this->user->attributes = $data;
        if($this->user->validate()){
            $res = $this->user->deleteAll(['id'=>$data['id']]);
            return $res?$this->success('删除会员成功！')
                :$this->error('删除会员失败，请稍后再试！');
        }
        return $this->error(current($this->user->firstErrors));
    }
}
