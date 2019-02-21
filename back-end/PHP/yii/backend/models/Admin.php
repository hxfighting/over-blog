<?php

namespace app\models;

use mikk150\phonevalidator\PhoneNumberValidator;
use riskivy\captcha\CaptchaHelper;
use Yii;
use yii\web\IdentityInterface;

/**
 * This is the model class for table "admin".
 *
 * @property int    $id
 * @property string $name       用户名
 * @property string $password   密码
 * @property string $remember_token
 * @property string $email      邮箱
 * @property string $mobile     电话
 * @property string $avatar     头像
 * @property int    $created_at 创建时间
 * @property int    $updated_at 修改时间
 */
class Admin extends \yii\db\ActiveRecord implements IdentityInterface
{
    public $captcha;
    public $password_confirmation;

    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'admin';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            ['name', 'required', 'on' => ['changeInfo', 'login']],
            ['captcha', 'required','on' => 'login'],
            ['captcha', 'validateCaptcha','on' => 'login'],
            ['name', 'string', 'max' => 30, 'on' => ['changeInfo', 'login']],
            [['email', 'mobile', 'avatar'], 'required', 'on' => 'changeInfo'],
            ['avatar', 'url', 'on' => 'changeInfo'],
            ['password', 'required', 'on' => ['changePassword', 'login']],
            ['password_confirmation','required','on' => 'changePassword'],
            ['password', 'compare', 'compareAttribute' => 'password_confirmation','on' => 'changePassword'],
            ['password', 'string', 'min' => 6, 'max' => 16, 'on' => ['changePassword', 'login']],
            [['email'], 'string', 'max' => 60, 'on' => 'changeInfo'],
            [['email'], 'email', 'on' => 'changeInfo'],
            [['mobile'], 'string', 'max' => 11, 'on' => 'changeInfo'],
            ['mobile', PhoneNumberValidator::class, 'country' => 'CN',
                'on' => 'changeInfo','message' => '请输入正确的电话号码',
                'invalidFormatMessage' => '请输入正确的电话号码']
        ];
    }

    /**
     * 验证验证码
     * Date: 2019-02-20 10:13
     * @param $attribute
     * @param $params
     * @throws \yii\base\Exception
     * @throws \yii\base\InvalidConfigException
     */
    public function validateCaptcha($attribute, $params)
    {
        $res = (new CaptchaHelper())->verify($this->$attribute);
        if(!$res){
            $this->addError($attribute,$this->$attribute);
        }
    }


    public function getId()
    {
        return $this->getPrimaryKey();
    }

    public function getAuthKey()
    {
        // TODO: Implement getAuthKey() method.
    }

    public static function findIdentity($id)
    {
        // TODO: Implement findIdentity() method.
    }

    public static function findIdentityByAccessToken($token, $type = null)
    {
        // TODO: Implement findIdentityByAccessToken() method.
    }

    public function validateAuthKey($authKey)
    {
        // TODO: Implement validateAuthKey() method.
    }
}
