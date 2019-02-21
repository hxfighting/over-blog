<?php

namespace app\models;

use Yii;

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
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'rotation_image';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['created_at', 'updated_at'], 'integer'],
            [['words'], 'string', 'max' => 100],
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
