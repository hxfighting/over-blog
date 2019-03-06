<?php

namespace app\models;

use Yii;
use yii\behaviors\AttributeTypecastBehavior;

/**
 * This is the model class for table "contact".
 *
 * @property int $id
 * @property string $name 留言人姓名
 * @property string $content 留言内容
 * @property string $email 留言人email
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 * @property int $is_reply 是否回复,1是0否
 * @property string $reply_content 回复内容
 * @property int $replied_at 回复时间
 */
class Contact extends \yii\db\ActiveRecord
{
    public $pageSize;
    public $pageNum;
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'contact';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            ['id','required','on' => ['delContact','contactReply']],
            ['id','integer','min' => 1,'on' => ['delContact','contactReply']],
            [['pageSize','pageNum'],'required','on' => 'contactList'],
            [['pageSize','pageNum'],'integer','min' => 1,'on' => 'contactList'],
            ['reply_content', 'required','on' => 'contactReply'],
            ['reply_content', 'string', 'max' => 255,'on' => 'contactReply']
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
                    'replied_at' => function ($value) {
                        return $value?date('Y-m-d H:i:s', $value):null;
                    },
                ],
                'typecastAfterValidate' => true,
                'typecastBeforeSave' => false,
                'typecastAfterFind' => true,
            ]
        ];
    }
}
