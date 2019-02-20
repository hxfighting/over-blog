<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-20
 * Time: 16:51
 */

namespace backend\filter;


use Lcobucci\JWT\Builder;
use Lcobucci\JWT\Parser;
use Lcobucci\JWT\ValidationData;
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
            if($code==$this->code['expire']){
                $this->expire = true;
            }
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

    private function validateToken()
    {
        $token = \Yii::$app->request->headers->get('Authorization', null);
        if (!$token)
        {
            return $this->code['error'];
        }
        preg_match('/^Bearer\s+(.*?)$/', $token, $matches);
        if (!isset($matches[1]) || empty($matches[1]))
        {
            return $this->code['error'];
        }
        try
        {
            $token = (new Parser())->parse($matches[1]);
            if(!$this->checkExpire($token)){
                return $this->code['expire'];
            }
            $data = new ValidationData();
            $data->setIssuer(\Yii::$app->params['issuer']);
            $data->setAudience(\Yii::$app->params['audience']);
            $data->setId(\Yii::$app->params['jwt_id']);
            return $token->validate($data)?$this->code['success']
                :$this->code['error'];
        } catch (\Exception $e)
        {
            return $this->code['error'];
        }
    }

    /**
     * 检测是否是当天的token
     * Date: 2019-02-20 17:21
     * @param $token
     * @return bool
     */
    private function checkExpire($token)
    {
        $expire = $token->getClaim('exp');
        if (!$expire)
        {
            return false;
        }
        $diff_day = (time() - $expire) / 86400;
        if($diff_day>=1){
            $this->user_id = $token->getClaim('uid',0);
            return false;
        }
        return true;
    }

    public function afterAction($action, $result)
    {
        if($this->expire){
            $headers = \Yii::$app->response->headers;
            $headers->set('Authorization',$this->getToken());
        }
        return parent::afterAction($action, $result);
    }

    /**
     * 获得token
     * Date: 2019-02-20 17:48
     * @return string
     */
    private function getToken()
    {
        $token = (new Builder())
            ->setIssuer(\Yii::$app->params['issuer'])// Configures the issuer (iss claim)
            ->setAudience(\Yii::$app->params['audience'])// Configures the audience (aud claim)
            ->setId(\Yii::$app->params['jwt_id'], true)// Configures the id (jti claim), replicating as a header item
            ->setIssuedAt(time())// Configures the time that the token was issue (iat claim)
            ->setNotBefore(time())// Configures the time before which the token cannot be accepted (nbf claim)
            ->setExpiration(time() + \Yii::$app->params['jwt_expire'])// Configures the expiration time of the token (exp claim)
            ->set('uid', $this->user_id)// Configures a new claim, called "uid"
            ->getToken(); // Retrieves the generated token
        return 'Bearer '.$token;
    }
}
