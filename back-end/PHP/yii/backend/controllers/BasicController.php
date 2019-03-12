<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-18
 * Time: 17:20
 */

namespace backend\controllers;


use backend\exception\ValidateException;
use backend\filter\TokenFilter;
use common\helper\Token;
use riskivy\captcha\CaptchaHelper;
use yii\filters\Cors;
use yii\helpers\ArrayHelper;
use yii\web\Controller;
use yii\web\Response;

class BasicController extends Controller
{
    protected $user_id;

    //请求的数据
    protected $request_data;

    //验证失败的信息
    protected $validate_error;

    protected $except_action = [
        'login', 'captcha', 'log-error', 'index'
    ];


    public function behaviors()
    {
        return ArrayHelper::merge([
//            [
//                'class'=>TokenFilter::class,
//                'except' => $this->except_action
//            ],
            [
                'class' => Cors::class,
                'cors' => [
                    'Origin' => ['*'],
                    'Access-Control-Request-Method' => ['*'],
                    'Access-Control-Request-Headers' => ['*'],
                    'Access-Control-Max-Age' => 86400,
                    'Access-Control-Expose-Headers' => ['Authorization']
                ],

            ],
        ], parent::behaviors());
    }

    public function beforeAction($action)
    {
//        $actions = \Yii::$app->controller->action->id;
//        if(!in_array($actions,$this->except_action)){
//            $this->user_id = Token::getUserId();
//        }
        $request = \Yii::$app->request;
        if ($request->isPost || $request->isPut || $request->isDelete)
        {
            $this->request_data = $this->post();
        } else
        {
            $this->request_data = $this->get();
        }
        return parent::beforeAction($action);
    }

    /**
     * 获取get参数
     * Date: 2019-02-22 11:25
     * @param null $key
     * @param null $value
     * @return array|mixed
     */
    public function get($key = null, $value = null)
    {
        return \Yii::$app->request->get($key, $value);
    }

    /**
     * 获取post参数
     * Date: 2019-02-18 17:28
     * @param null $name
     * @param null $defaultValue
     * @return array|mixed
     */
    public function post($name = null, $defaultValue = null)
    {
        return \Yii::$app->request->post($name, $defaultValue);
    }

    /**
     * 成功响应
     * Date: 2019-02-19 10:24
     * @param      $msg
     * @param null $data
     * @param int  $code
     * @return Response
     */
    public function success($msg, $data = null, $code = 200)
    {
        $data = compact('msg', 'data', 'code');
        return $this->asJson($data);
    }

    /**
     * 失败响应
     * Date: 2019-02-19 10:25
     * @param      $msg
     * @param null $data
     * @param int  $code
     * @return Response
     */
    public function error($msg, $data = null, $code = -1)
    {
        $data = compact('msg', 'data', 'code');
        return $this->asJson($data);
    }

    /**
     * 获取验证码
     * Date: 2019-02-21 09:44
     * @return Response
     * @throws \yii\base\Exception
     * @throws \yii\base\InvalidConfigException
     */
    public function actionCaptcha()
    {
        $captcha = (new CaptchaHelper())->generateImage();
        $data = [
            'sensitive' => $this->generateKey(),
            'key' => $this->generateKey(),
            'img' => $captcha
        ];
        return $this->success('获取验证码成功', $data);
    }

    /**
     * 生成随机字符串
     * Date: 2019-02-19 11:04
     * @param int $length
     * @return false|string
     * @throws \yii\base\Exception
     */
    protected function generateKey($length = 32)
    {
        return \Yii::$app->security->generateRandomString($length);
    }

    /**
     * 统一验证
     * Date: 2019-03-12 10:13
     * @param $model
     * @param $scenario
     * @return mixed
     * @throws ValidateException
     */
    protected function basicValidate($model, $scenario)
    {
        $model->scenario = $scenario;
        $model->attributes = $this->request_data;
        if (!$model->validate())
        {
            throw new ValidateException(current($model->firstErrors));
        }
        return $model;
    }
}
