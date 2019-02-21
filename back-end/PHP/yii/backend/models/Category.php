<?php

namespace app\models;

use Yii;

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

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['title', 'pid'], 'required'],
            [['pid', 'created_at', 'updated_at', 'type'], 'integer'],
            [['title'], 'string', 'max' => 50],
            [['url'], 'string', 'max' => 100],
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
            'pid' => 'Pid',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
            'url' => 'Url',
            'type' => 'Type',
        ];
    }
}
