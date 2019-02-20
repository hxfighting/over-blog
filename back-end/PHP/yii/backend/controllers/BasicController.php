<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-18
 * Time: 17:20
 */

namespace backend\controllers;


use backend\filter\TokenFilter;
use Lcobucci\JWT\Parser;
use Lcobucci\JWT\ValidationData;
use riskivy\captcha\CaptchaHelper;
use sizeg\jwt\JwtHttpBearerAuth;
use yii\filters\Cors;
use yii\helpers\ArrayHelper;
use yii\helpers\VarDumper;
use yii\web\Controller;
use yii\web\Response;

class BasicController extends Controller
{
    protected $user_id;

    public function behaviors()
    {
        return ArrayHelper::merge([
            [
                'class'=>TokenFilter::class,
                'except' => [
                    'login', 'captcha', 'log-error', 'index'
                ]
            ],
            [
                'class' => Cors::class,
                'cors' => [
                    'Origin' => ['*'],
                    'Access-Control-Request-Method' => ['*'],
                    'Access-Control-Request-Headers' => ['*']
                ],

            ],
        ], parent::behaviors());
    }

    /**
     * 获取get参数
     * Date: 2019-02-18 17:27
     * @param        $key
     * @param string $value
     * @return array|mixed
     */
    public function get($key, $value = '')
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
     * Date: 2019-02-19 10:26
     * @return Response
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
}
