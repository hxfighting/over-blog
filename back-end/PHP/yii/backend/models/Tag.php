<?php

namespace app\models;

use Yii;
use yii\behaviors\AttributeTypecastBehavior;
use yii\behaviors\TimestampBehavior;

/**
 * This is the model class for table "tag".
 *
 * @property int    $id
 * @property string $name       标签名称
 * @property int    $created_at 创建时间
 * @property int    $updated_at 修改时间
 */
class Tag extends \yii\db\ActiveRecord
{
    public $pageSize;
    public $pageNum;

    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'tag';
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
            [['pageSize', 'pageNum'], 'required', 'on' => 'tagList'],
            [['pageSize', 'pageNum'], 'integer', 'on' => 'tagList'],
            [['pageSize', 'pageNum'], 'number', 'min' => 1, 'on' => 'tagList'],
            ['id', 'required', 'on' => ['tagDelete','tagUpdate']],
            ['id', 'integer', 'on' => ['tagDelete','tagUpdate']],
            ['id', 'exist', 'on' => ['tagDelete','tagUpdate']],
            ['name', 'required', 'on' => ['tagAdd', 'tagUpdate']],
            ['name', 'string', 'min' => 2, 'max' => 20, 'on' => ['tagAdd', 'tagUpdate']],
            ['name', 'unique', 'on' => ['tagAdd'],'message' => '标签名字已经存在，请换一个'],
            ['name', 'validateName', 'on' => ['tagUpdate']]
        ];
    }

    public function validateName($attribute, $params)
    {
        $exist_name = self::find()->where(['name'=>$this->$attribute])->one();
        if($exist_name && $exist_name->id!=$this->id){
            $this->addError($attribute,'标签名字已存在，请换一个！');
        }
    }

}
