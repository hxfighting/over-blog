<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-20
 * Time: 16:51
 */

namespace backend\filter;


use common\helper\Token;
use Lcobucci\JWT\Parser;
use yii\base\ActionFilter;
use yii\web\Response;

class TokenFilter extends ActionFilter
{
    private $code = [
        'error' => 0,
        'expire' => 1,
        'success' => 2
    ];
    private $expire = false;
    private $user_id;

    public function beforeAction($action)
    {
        if ($code = $this->validateToken())
        {
            return parent::beforeAction($action);
        } else
        {
            $response = \Yii::$app->getResponse();
            $response->format = Response::FORMAT_JSON;
            $response->data = [
                'code' => -1,
                'msg' => 'token非法！',
                'data' => null
            ];
            return false;
        }
    }

    /**
     * 验证token
     * Date: 2019-02-21 09:37
     * @return mixed
     */
    private function validateToken()
    {
        $token = \Yii::$app->request->headers->get('Authorization', null);
        if (!$token)
        {
            return $this->code['error'];
        }
        $key = md5($token);
        preg_match('/^Bearer\s+(.*?)$/', $token, $matches);
        if (!isset($matches[1]) || empty($matches[1]))
        {
            return $this->code['error'];
        }
        try
        {
            $token = $this->parseToken($matches[1]);
            if(!$this->checkExpire($token,$key)){
                return $this->code['error'];
            }
            return $this->code['success'];
        } catch (\Exception $e)
        {
            return $this->code['error'];
        }
    }

    /**
     * 解析token
     * Date: 2019-02-21 09:34
     * @param string $token
     * @return \Lcobucci\JWT\Token
     */
    private function parseToken(string $token)
    {
        return (new Parser())->parse($token);
    }

    /**
     * 检测token过期
     * Date: 2019-02-21 10:21
     * @param $token
     * @param $key
     * @return bool
     */
    private function checkExpire($token,$key)
    {
        $redis = \Yii::$app->redis;
        $exist_token = $redis->get($key);
        if(!$exist_token){
            return false;
        }
        if(time()>(int)$exist_token){
            $this->expire = true;
            $this->user_id = $token->getClaim('uid',0);
            $redis->del($key);
            return true;
        }
        return true;
    }

    public function afterAction($action, $result)
    {
        if($this->expire){
            $headers = \Yii::$app->response->headers;
            $headers->set('Authorization',Token::getToken($this->user_id));
        }
        return parent::afterAction($action, $result);
    }
}
