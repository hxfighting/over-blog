<?php

namespace app\models;

use Yii;
use yii\behaviors\AttributeTypecastBehavior;
use yii\behaviors\TimestampBehavior;

/**
 * This is the model class for table "chat".
 *
 * @property int $id
 * @property string $content 说说内容
 * @property int $is_show 是否显示1是,0否
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class Chat extends \yii\db\ActiveRecord
{
    public $pageSize;
    public $pageNum;
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'chat';
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
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            ['id','required','on' => ['updateChat','chatDelete']],
            ['id','integer','on' => ['updateChat','chatDelete']],
            [['pageSize','pageNum'],'required','on' => 'chatList'],
            [['pageSize','pageNum'],'integer','on' => 'chatList'],
            [['pageSize','pageNum'],'number','min' => 1,'on' => 'chatList'],
            [['content','is_show'], 'required','on' => ['addChat','updateChat']],
            [['is_show'], 'integer','on' => ['addChat','updateChat']],
            [['content'], 'string', 'max' => 255,'on' => ['addChat','updateChat']],
        ];
    }
}
