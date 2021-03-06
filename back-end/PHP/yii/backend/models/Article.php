<?php

namespace app\models;

use Yii;
use yii\behaviors\AttributeTypecastBehavior;

/**
 * This is the model class for table "article".
 *
 * @property int $id
 * @property string $title 文章标题
 * @property string $author 文章作者
 * @property string $content_html 文章html内容
 * @property string $content_md 文章markdown内容
 * @property string $keywords 文章关键词
 * @property string $description 文章描述
 * @property string $thumb 文章缩略图
 * @property int $is_show 是否显示,1是0否
 * @property int $is_original 是否原创,1是0否
 * @property int $is_top 是否置顶,1是0否
 * @property int $click 文章点击次数
 * @property int $category_id 文章分类ID
 * @property int $deleted_at 删除时间
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class Article extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'article';
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
            ]
        ];
    }
    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['title', 'author', 'content_html', 'keywords', 'description', 'thumb', 'category_id'], 'required'],
            [['content_html', 'content_md'], 'string'],
            [['is_show', 'is_original', 'is_top', 'click', 'category_id', 'deleted_at', 'created_at', 'updated_at'], 'integer'],
            [['title'], 'string', 'max' => 100],
            [['author'], 'string', 'max' => 20],
            [['keywords', 'description', 'thumb'], 'string', 'max' => 255],
        ];
    }
}
