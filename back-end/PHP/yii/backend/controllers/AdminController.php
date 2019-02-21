<?php

namespace backend\controllers;

use app\models\Admin;
use common\helper\Token;
use yii\filters\VerbFilter;
use yii\helpers\ArrayHelper;

class AdminController extends BasicController
{
    public $enableCsrfValidation = false;

    public function behaviors()
    {
        return ArrayHelper::merge([
            'verbs' => [
                'class' => VerbFilter::class,
                'actions' => [
                    'login' => ['post'],
                    'admin-info' => ['get']
                ],
            ],
        ], parent::behaviors());
    }

    /**
     * 登录
     * Date: 2019-02-20 09:43
     * @return \yii\web\Response
     */
    public function actionLogin()
    {
        $admin = new Admin();
        $admin->scenario = 'login';
        $data = $this->post();
        $admin->attributes = $data;
        if ($admin->validate())
        {
            $exist_admin = $admin->find()->where(['name' => $data['name']])->one();
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
            return $this->error(current($admin->firstErrors));
        }
    }


    /**
     * 获取管理员信息
     * Date: 2019-02-21 11:38
     * @return \yii\web\Response
     */
    public function actionAdminInfo()
    {
        $user = Admin::find()->where(['id'=>$this->user_id])->one();
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
}
