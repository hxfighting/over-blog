<?php

namespace app\models;

use Yii;
use yii\behaviors\TimestampBehavior;

/**
 * This is the model class for table "rotation_image".
 *
 * @property int $id
 * @property string $words 轮播图文字
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class RotationImage extends \yii\db\ActiveRecord
{
    public $image_url;
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'rotation_image';
    }

    public function behaviors()
    {
        return [
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
            ['id','required','on' => ['rotationUpdate','delRotation']],
            ['id','integer','on' => ['rotationUpdate','delRotation']],
            [['image_url','words'],'required','on' => ['rotationAdd','rotationUpdate']],
            ['image_url','url','on' => ['rotationAdd','rotationUpdate']],
            ['words', 'string', 'max' => 100,'min' => 2,'on' => ['rotationAdd','rotationUpdate']]
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'words' => 'Words',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
        ];
    }
}
