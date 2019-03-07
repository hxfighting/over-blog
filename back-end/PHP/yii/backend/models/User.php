<?php

namespace app\models;

use Yii;
use yii\behaviors\AttributeTypecastBehavior;
use yii\behaviors\TimestampBehavior;

/**
 * This is the model class for table "user".
 *
 * @property int $id
 * @property int $type 用户类型 1：QQ  2：微信 3：新浪微博
 * @property string $name 第三方用户名称
 * @property string $avatar 第三方用户头像
 * @property string $openid 第三方用户openid
 * @property string $access_token access_token token（当是微信时存的是session_key）
 * @property string $last_login_ip 最后登录IP
 * @property int $login_times 登录次数
 * @property string $email 邮箱
 * @property int $is_admin 是否是管理员,1是,0否
 * @property string $remember_token
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class User extends \yii\db\ActiveRecord
{
    public $pageSize;
    public $pageNum;
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'user';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            ['id','required','on' => ['updateUser','delUser']],
            ['id','integer','min' => 1,'on' => ['updateUser','delUser']],
            ['is_admin','required','on' => 'updateUser'],
            ['is_admin','integer','on' => 'updateUser'],
            [['pageSize', 'pageNum'], 'required', 'on' => 'userList'],
            [['pageSize', 'pageNum'], 'integer', 'on' => 'userList']
        ];
    }

    public function behaviors()
    {
        return [
            'typecast' => [
                'class' => AttributeTypecastBehavior::class,
                'attributeTypes' => [
                    'created_at' => function ($value) {
                        return date('Y-m-d H:i:s', $value);
                    },
                    'updated_at' => function ($value) {
                        return date('Y-m-d H:i:s', $value);
                    },
                    'last_login_ip' => function($value){
                        return $this->transformIPToAddress($value);
                    }
                ],
                'typecastAfterValidate' => true,
                'typecastBeforeSave' => false,
                'typecastAfterFind' => true,
            ],
            [
                'class' => TimestampBehavior::class
            ]
        ];
    }

    /**
     * 转换IP地址为具体地址
     * Date: 2019-03-07 13:35
     * @param $ip
     * @return string
     */
    private function transformIPToAddress($ip)
    {
        $address = \Yii::$app->geoip->ip($ip);
        if($address->hasResult()){
            return $this->handleAddress($address);
        }
        return '未知';
    }

    //处理转换后的地址数据
    private function handleAddress($address)
    {
        switch ($address->country->isoCode){
            case 'CN':
                return $address->country->names->{'zh-CN'}.'-'.$address->city->names->{'zh-CN'};
            case 'TW':
                return '中国-台湾';
            case 'MO':
                return '中国-澳门';
            default:
                return $address->country->names->{'zh-CN'}.'-'.$address->city->names->{'zh-CN'};
        }
    }
}
