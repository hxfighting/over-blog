<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-21
 * Time: 09:39
 */

namespace common\helper;


use Lcobucci\JWT\Builder;
use Lcobucci\JWT\Parser;

class Token
{
    /**
     * 生成token
     * Date: 2019-02-21 09:42
     * @param $user_id
     * @return string
     */
    public static function getToken($user_id)
    {
        $expire = time() + \Yii::$app->params['jwt_expire'];
        $token = (new Builder())
            ->setIssuer(\Yii::$app->params['issuer'])// Configures the issuer (iss claim)
            ->setAudience(\Yii::$app->params['audience'])// Configures the audience (aud claim)
            ->setId(\Yii::$app->params['jwt_id'], true)// Configures the id (jti claim), replicating as a header item
            ->setIssuedAt(time())// Configures the time that the token was issue (iat claim)
            ->setNotBefore(time())// Configures the time before which the token cannot be accepted (nbf claim)
            ->setExpiration($expire)// Configures the expiration time of the token (exp claim)
            ->set('uid', $user_id)// Configures a new claim, called "uid"
            ->getToken(); // Retrieves the generated token
        $token = 'Bearer '.$token;
        self::setCache($token,$expire);
        return $token;
    }

    /**
     * 缓存token
     * Date: 2019-02-21 10:35
     * @param string $token
     * @param int    $expire
     */
    private static function setCache(string $token,int $expire)
    {
        $redis = \Yii::$app->redis;
        $key = md5($token);
        $redis->set($key,$expire);
        $time_diff = strtotime(date('Ymd'.' 23:59:59'))-time();
        $redis->expire($key,$time_diff);
    }

    /**
     * 解析token获取用户ID
     * Date: 2019-02-21 11:34
     * @return mixed
     */
    public static function getUserId()
    {
        $token = \Yii::$app->request->headers->get('Authorization', null);
        preg_match('/^Bearer\s+(.*?)$/', $token, $matches);
        $token = (new Parser())->parse($matches[1]);
        $user_id = $token->getClaim('uid');
        return $user_id;
    }
}
