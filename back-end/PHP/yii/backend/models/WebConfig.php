<?php

namespace app\models;

use Yii;

/**
 * This is the model class for table "web_config".
 *
 * @property int $id
 * @property string $title 配置名称
 * @property string $name 配置变量名
 * @property string $val 配置值
 * @property int $type 字段类型,1:社交地址,2footer内容,3其他配置
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class WebConfig extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'web_config';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['title', 'name', 'val', 'type'], 'required'],
            [['val'], 'string'],
            [['type', 'created_at', 'updated_at'], 'integer'],
            [['title', 'name'], 'string', 'max' => 100],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'title' => 'Title',
            'name' => 'Name',
            'val' => 'Val',
            'type' => 'Type',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
        ];
    }
}
