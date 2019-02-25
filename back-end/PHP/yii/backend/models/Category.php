<?php

namespace app\models;

use Yii;
use yii\behaviors\AttributeTypecastBehavior;
use yii\behaviors\TimestampBehavior;

/**
 * This is the model class for table "category".
 *
 * @property int $id
 * @property string $title 分类名称
 * @property int $pid 父ID
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 * @property string $url 分类URL
 * @property int $type 分类类型:1文章分类,2联系我,3说说
 */
class Category extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'category';
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
            ['id','required','on' => ['categoryUpdate','delCategory']],
            ['id','integer','on' => ['categoryUpdate','delCategory']],
            [['title', 'pid'], 'required','on' => ['categoryAdd','categoryUpdate']],
            ['pid', 'integer','on' => ['categoryAdd','categoryUpdate']],
            [['title'], 'string', 'max' => 50, 'on' => ['categoryAdd','categoryUpdate']]
        ];
    }

    public function getChildren()
    {
        return $this->hasMany(self::class,['pid'=>'id']);
    }
}
