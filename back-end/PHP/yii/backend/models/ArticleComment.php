<?php

namespace app\models;

use Yii;
use yii\behaviors\AttributeTypecastBehavior;

/**
 * This is the model class for table "article_comment".
 *
 * @property int $id
 * @property int $pid 父ID,只是为了评论的层级,这里设置为两层
 * @property int $reply_id 被评论人的ID
 * @property int $user_id 评论人ID
 * @property int $article_id 文章ID
 * @property string $content 评论内容
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class ArticleComment extends \yii\db\ActiveRecord
{

    public $pageSize;
    public $pageNum;
    public $reply_content;

    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'article_comment';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            ['id','required','on' => ['commentDelete','commentReply']],
            ['id','integer','min' => 1,'on' => ['commentDelete','commentReply']],
            ['reply_content','required','on' => 'commentReply'],
            ['reply_content','string','min' => 2,'max' => 255,'on' => 'commentReply'],
            [['pageSize','pageNum'],'required','on' => 'commentList'],
            [['pageSize','pageNum'],'integer','min' => 1,'on' => 'commentList']
        ];
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

}
