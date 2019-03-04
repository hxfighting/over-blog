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
    public $pageSize;
    public $pageNum;
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
            ['id','required','on' => ['updateLink','delLink']],
            ['id','integer','min' => 1,'on' => ['updateLink','delLink']],
            [['pageSize','pageNum'],'required','on' => 'linkList'],
            [['pageSize','pageNum'],'integer','min' => 1,'on' => 'linkList'],
            [['name','url','description','order','is_show'], 'required','on' => ['addLink','updateLink']],
            [['name','description'], 'string','min' => 1,'max' => 50,'on' => ['addLink','updateLink']],
            [['order','is_show'],'integer','on' => ['addLink','updateLink']],
            [['url'], 'string', 'max' => 255,'on' => ['addLink','updateLink']],
            [['url'], 'url', 'max' => 255,'on' => ['addLink','updateLink']]
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
