<?php

namespace app\models;

use Yii;

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
            [['type', 'name', 'avatar', 'openid', 'access_token', 'last_login_ip'], 'required'],
            [['type', 'login_times', 'is_admin', 'created_at', 'updated_at'], 'integer'],
            [['name', 'openid', 'access_token', 'email', 'remember_token'], 'string', 'max' => 100],
            [['avatar'], 'string', 'max' => 255],
            [['last_login_ip'], 'string', 'max' => 45],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'type' => 'Type',
            'name' => 'Name',
            'avatar' => 'Avatar',
            'openid' => 'Openid',
            'access_token' => 'Access Token',
            'last_login_ip' => 'Last Login Ip',
            'login_times' => 'Login Times',
            'email' => 'Email',
            'is_admin' => 'Is Admin',
            'remember_token' => 'Remember Token',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
        ];
    }
}
