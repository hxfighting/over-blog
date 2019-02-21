<?php

namespace app\models;

use Yii;

/**
 * This is the model class for table "image".
 *
 * @property int $id
 * @property string $image_type
 * @property string $image_id
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 * @property string $image_url
 */
class Image extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'image';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['image_type', 'image_id', 'image_url'], 'required'],
            [['image_id', 'created_at', 'updated_at'], 'integer'],
            [['image_type'], 'string', 'max' => 191],
            [['image_url'], 'string', 'max' => 255],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'image_type' => 'Image Type',
            'image_id' => 'Image ID',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
            'image_url' => 'Image Url',
        ];
    }
}
