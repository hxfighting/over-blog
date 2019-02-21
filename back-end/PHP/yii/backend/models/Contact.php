<?php

namespace app\models;

use Yii;

/**
 * This is the model class for table "contact".
 *
 * @property int $id
 * @property string $name 留言人姓名
 * @property string $content 留言内容
 * @property string $email 留言人email
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 * @property int $is_reply 是否回复,1是0否
 * @property string $reply_content 回复内容
 * @property int $replied_at 回复时间
 */
class Contact extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'contact';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['name', 'content', 'email'], 'required'],
            [['created_at', 'updated_at', 'is_reply', 'replied_at'], 'integer'],
            [['name'], 'string', 'max' => 30],
            [['content', 'reply_content'], 'string', 'max' => 255],
            [['email'], 'string', 'max' => 60],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'name' => 'Name',
            'content' => 'Content',
            'email' => 'Email',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
            'is_reply' => 'Is Reply',
            'reply_content' => 'Reply Content',
            'replied_at' => 'Replied At',
        ];
    }
}
