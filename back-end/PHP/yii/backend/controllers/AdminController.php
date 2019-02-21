<?php

namespace backend\controllers;

use app\models\Admin;
use common\helper\Token;

class AdminController extends BasicController
{
    public $enableCsrfValidation = false;
    private $admin;

    public function init()
    {
        parent::init();
        $this->admin = new Admin();
    }

    /**
     * 登录
     * Date: 2019-02-20 09:43
     * @return \yii\web\Response
     */
    public function actionLogin()
    {
        $this->admin->scenario = 'login';
        $data = $this->post();
        $this->admin->attributes = $data;
        if ($this->admin->validate())
        {
            $exist_admin = $this->admin->find()->where(['name' => $data['name']])->one();
            if (!$exist_admin)
            {
                return $this->error('用户名或密码错误！');
            }
            if (!password_verify($data['password'], $exist_admin->password))
            {
                return $this->error('用户名或密码错误！');
            }
            $token = Token::getToken($exist_admin->id);
            return $this->success('登录成功！', $token);
        }else{
            return $this->error(current($this->admin->firstErrors));
        }
    }


    /**
     * 获取管理员信息
     * Date: 2019-02-21 11:38
     * @return \yii\web\Response
     */
    public function actionAdminInfo()
    {
        $user = $this->admin->find()->where(['id'=>$this->user_id])->one();
        if(!$user){
            return $this->error('暂无该用户信息！');
        }
        $index = strripos($user->avatar,'/');
        $user_data = [
            'avatar'    => $user->avatar,
            'name'      => $user->name,
            'user_id'   => $user->id,
            'access'    => ['super_admin'],
            'email'     => $user->email,
            'avatarName'=> substr($user->avatar,$index+1),
            'mobile'    => $user->mobile
        ];
        return $this->success('获取用户信息成功', $user_data);
    }

    /**
     * 退出登录
     * Date: 2019-02-21 11:52
     * @return \yii\web\Response
     */
    public function actionLogout()
    {
        $res = Token::removeToken();
        return $res?$this->success('退出登录成功！')
            :$this->error('退出登录失败，请稍后再试！');
    }

    /**
     * 修改密码
     * Date: 2019-02-21 13:13
     * @return \yii\web\Response
     */
    public function actionPassword()
    {
        $data = $this->post();
        $exist_admin = $this->admin->findOne($this->user_id);
        $exist_admin->scenario = 'changePassword';
        $exist_admin->attributes = $data;
        if($exist_admin->validate()){
            $exist_admin->password = password_hash($data['password'],PASSWORD_BCRYPT);
            $res = $exist_admin->save(false,['password']);
            return $res?$this->success('修改密码成功！')
                :$this->error('修改密码失败，请稍后再试！');
        }
        return $this->error(current($exist_admin->firstErrors));

    }

    /**
     * 修改个人信息
     * Date: 2019-02-21 13:19
     * @return \yii\web\Response
     */
    public function actionChangeInfo()
    {
        $data = $this->post();
        $exist_admin = $this->admin->findOne($this->user_id);
        $exist_admin->scenario = 'changeInfo';
        $exist_admin->attributes = $data;
        if($exist_admin->validate()){
            $exist_admin->mobile = ltrim($data['mobile'], '+86');
            $res = $exist_admin->save(false);
            return $res?$this->success('修改个人信息成功！')
                :$this->error('修改个人信息失败，请稍后再试！');
        }
        return $this->error(current($exist_admin->firstErrors));

    }
}
