<?php

namespace backend\controllers;

use app\models\Admin;
use Lcobucci\JWT\Builder;
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
            $token = $this->getToken($exist_admin->id);
            $token = 'Bearer ' . $token;
            return $this->success('登录成功！', $token);
        }else{
            return $this->error(current($admin->firstErrors));
        }
    }

    /**
     * 生成token
     * Date: 2019-02-20 09:44
     * @param $user_id
     * @return mixed
     */
    private function getToken($user_id)
    {
        $token = (new Builder())
            ->setIssuer(\Yii::$app->params['issuer'])// Configures the issuer (iss claim)
            ->setAudience(\Yii::$app->params['audience'])// Configures the audience (aud claim)
            ->setId(\Yii::$app->params['jwt_id'], true)// Configures the id (jti claim), replicating as a header item
            ->setIssuedAt(time())// Configures the time that the token was issue (iat claim)
            ->setNotBefore(time())// Configures the time before which the token cannot be accepted (nbf claim)
            ->setExpiration(time() + \Yii::$app->params['jwt_expire'])// Configures the expiration time of the token (exp claim)
            ->set('uid', $user_id)// Configures a new claim, called "uid"
//            ->sign($sign,\Yii::$app->params['jwt_sign'])
            ->getToken(); // Retrieves the generated token
        return $token;
    }


    public function actionAdminInfo()
    {
        var_dump(123);
    }
}
