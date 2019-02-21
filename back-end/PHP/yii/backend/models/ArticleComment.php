<?php

namespace app\models;

use Yii;

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
            [['pid', 'reply_id', 'user_id', 'article_id', 'created_at', 'updated_at'], 'integer'],
            [['user_id', 'article_id', 'content'], 'required'],
            [['content'], 'string', 'max' => 255],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'pid' => 'Pid',
            'reply_id' => 'Reply ID',
            'user_id' => 'User ID',
            'article_id' => 'Article ID',
            'content' => 'Content',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
        ];
    }
}
