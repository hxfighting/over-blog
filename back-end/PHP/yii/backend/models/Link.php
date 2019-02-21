<?php

namespace app\models;

use Yii;

/**
 * This is the model class for table "link".
 *
 * @property int $id
 * @property string $url 友联url
 * @property string $name 友联名称
 * @property string $description 友联描述
 * @property int $order 友联排序
 * @property int $is_show 友联是否显示,1是,0否
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class Link extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'link';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['url', 'name', 'description'], 'required'],
            [['order', 'is_show', 'created_at', 'updated_at'], 'integer'],
            [['url'], 'string', 'max' => 255],
            [['name', 'description'], 'string', 'max' => 50],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'url' => 'Url',
            'name' => 'Name',
            'description' => 'Description',
            'order' => 'Order',
            'is_show' => 'Is Show',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
        ];
    }
}
